
data "azurerm_resource_group" "main" {
  name = var.resource_group_name
}

data "azurerm_subnet" "k8s_subnet" {
  name                 =  var.k8s_vnet_subnet_name
  virtual_network_name =  var.vnet_name
  resource_group_name  = var.resource_group_name
}

resource "azurerm_kubernetes_cluster" "aks_cluster" {
  name                    = var.cluster_name == null ? "${var.prefix}-aks" : var.cluster_name
  kubernetes_version      = var.kubernetes_version
  location                = data.azurerm_resource_group.main.location
  resource_group_name     = data.azurerm_resource_group.main.name
  dns_prefix              = var.prefix
  sku_tier                = var.sku_tier
  private_cluster_enabled = var.private_cluster_enabled


  dynamic "default_node_pool" {
    for_each = var.default_node_pool
    content {
      orchestrator_version   = default_node_pool.value.kubernetes_version
      name                   = default_node_pool.value.name
      vm_size                = default_node_pool.value.vm_size
      vnet_subnet_id         = data.azurerm_subnet.k8s_subnet.id
      enable_auto_scaling    = default_node_pool.value.enable_auto_scaling
      node_count             = default_node_pool.value.node_count
      max_count              = default_node_pool.value.max_count
      min_count              = default_node_pool.value.min_count
      enable_node_public_ip  = var.enable_node_public_ip
      availability_zones     = default_node_pool.value.availability_zones
      node_labels            = default_node_pool.value.node_labels
      os_disk_type           = default_node_pool.value.os_disk_type
      os_disk_size_gb        = default_node_pool.value.os_disk_size_gb
      tags                   = var.tags
      type                   = default_node_pool.value.type
      max_pods               = default_node_pool.value.max_pods
      enable_host_encryption = var.enable_host_encryption
    }
  }

  dynamic "service_principal" {
    for_each = var.client_id != "" && var.client_secret != "" ? ["service_principal"] : []
    content {
      client_id     = var.client_id
      client_secret = var.client_secret
    }
  }

  dynamic "identity" {
    for_each = var.client_id == "" || var.client_secret == "" ? ["identity"] : []
    content {
      type                      = var.identity_type
      user_assigned_identity_id = var.user_assigned_identity_id
    }
  }

  addon_profile {
    http_application_routing {
      enabled = var.enable_http_application_routing
    }

    kube_dashboard {
      enabled = var.enable_kube_dashboard
    }

    azure_policy {
      enabled = var.enable_azure_policy
    }

    oms_agent {
      enabled                    = var.enable_log_analytics_workspace
      log_analytics_workspace_id = var.enable_log_analytics_workspace ? azurerm_log_analytics_workspace.main[0].id : null
    }
  }

  role_based_access_control {
    enabled = var.enable_role_based_access_control

    dynamic "azure_active_directory" {
      for_each = var.enable_role_based_access_control && var.rbac_aad_managed ? ["rbac"] : []
      content {
        managed                = true
        admin_group_object_ids = var.rbac_aad_admin_group_object_ids
      }
    }

    dynamic "azure_active_directory" {
      for_each = var.enable_role_based_access_control && !var.rbac_aad_managed ? ["rbac"] : []
      content {
        managed           = false
        client_app_id     = var.rbac_aad_client_app_id
        server_app_id     = var.rbac_aad_server_app_id
        server_app_secret = var.rbac_aad_server_app_secret
      }
    }
  }

  network_profile {
    network_plugin     = var.network_plugin
    network_policy     = var.network_policy
    dns_service_ip     = var.net_profile_dns_service_ip
    docker_bridge_cidr = var.net_profile_docker_bridge_cidr
    outbound_type      = var.net_profile_outbound_type
    pod_cidr           = var.net_profile_pod_cidr
    service_cidr       = var.net_profile_service_cidr
  }

  tags = var.tags
}


resource "azurerm_kubernetes_cluster_node_pool" "aks_node_pool" {
  for_each = var.node_pools
  name                   = each.value.name
  kubernetes_cluster_id  = azurerm_kubernetes_cluster.aks_cluster.id
  vm_size                = each.value.vm_size
  node_count             = each.value.node_count
  max_count              = each.value.max_count
  min_count              = each.value.min_count
  enable_auto_scaling    = each.value.enable_auto_scaling
  enable_host_encryption = each.value.enable_host_encryption
  availability_zones     = each.value.availability_zones
  mode                   = each.value.mode
  node_labels            = each.value.node_labels
  orchestrator_version   = var.kubernetes_version
  os_disk_size_gb        = each.value.os_disk_size_gb
  os_disk_type           = each.value.os_disk_type
  os_type                = each.value.os_type
  max_pods               = each.value.max_pods
  priority               = each.value.priority
  tags                   = var.tags
  vnet_subnet_id         = data.azurerm_subnet.k8s_subnet.id
}


resource "azurerm_log_analytics_workspace" "main" {
  count               = var.enable_log_analytics_workspace ? 1 : 0
  name                = var.cluster_log_analytics_workspace_name == null ? "${var.prefix}-workspace" : var.cluster_log_analytics_workspace_name
  location            = data.azurerm_resource_group.main.location
  resource_group_name = var.resource_group_name
  sku                 = var.log_analytics_workspace_sku
  retention_in_days   = var.log_retention_in_days

  tags = var.tags
}

resource "azurerm_log_analytics_solution" "main" {
  count                 = var.enable_log_analytics_workspace ? 1 : 0
  solution_name         = "ContainerInsights"
  location              = data.azurerm_resource_group.main.location
  resource_group_name   = data.azurerm_resource_group.main.name
  workspace_resource_id = azurerm_log_analytics_workspace.main[0].id
  workspace_name        = azurerm_log_analytics_workspace.main[0].name

  plan {
    publisher = "Microsoft"
    product   = "OMSGallery/ContainerInsights"
  }

  tags = var.tags
}
