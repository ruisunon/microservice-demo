

locals {

  rg_name          = var.rg_name
  k8s_cluster_name = var.k8s_cluster_name
  common_tags      = var.common_tags

}

module "resource_group" {
  source   = "../modules/resource_group"
  location = var.location
  rg_name  = var.rg_name
  tags     = var.common_tags
}

module "virtual_network" {
  source              = "Azure/network/azurerm"
  resource_group_name = module.resource_group.rg_name
  vnet_name           = var.vnet_name
  address_spaces      = var.vnet_cidr_blocks
  subnet_prefixes     = var.subnet_cidr_blocks
  subnet_names        = var.subnet_names
  tags                = local.common_tags

  depends_on = [module.resource_group]
}


module "acr" {
  source              = "../modules/azure_acr"
  acr_name            = var.acr_name
  resource_group_name = module.resource_group.rg_name
  sku                 = var.acr_sku
  tags                = local.common_tags

  depends_on = [module.resource_group, module.virtual_network]

}
module "aks" {
  source                           = "../modules/azure_aks"
  resource_group_name              = module.resource_group.rg_name
  kubernetes_version               = var.k8s_cluster_version
  orchestrator_version             = var.k8s_cluster_version
  prefix                           = var.k8s_prefix
  cluster_name                     = var.k8s_cluster_name
  network_plugin                   = var.network_plugin
  vnet_name                        = var.vnet_name
  sku_tier                         = var.k8s_sku_tier
  enable_role_based_access_control = true
  rbac_aad_admin_group_object_ids  = var.rbac_aad_admin_group_object_ids
  rbac_aad_managed                 = var.rbac_aad_managed
  private_cluster_enabled          = var.private_cluster_enabled
  enable_http_application_routing  = var.enable_http_application_routing
  enable_azure_policy              = var.enable_azure_policy
  enable_auto_scaling              = var.enable_auto_scaling
  enable_host_encryption           = var.enable_host_encryption
  enable_kube_dashboard            = var.enable_kube_dashboard
  default_node_pool                = var.default_node_pool
  enable_log_analytics_workspace   = var.enable_log_analytics_workspace
  node_pools                       = var.k8s_node_pools
  k8s_vnet_subnet_name             = var.k8s_vnet_subnet_name


  network_policy                 = var.network_policy
  net_profile_dns_service_ip     = var.net_profile_dns_service_ip
  net_profile_docker_bridge_cidr = var.net_profile_docker_bridge_cidr
  net_profile_outbound_type      = var.net_profile_outbound_type
  net_profile_pod_cidr           = var.net_profile_pod_cidr
  net_profile_service_cidr       = var.net_profile_service_cidr

  depends_on = [module.resource_group, module.virtual_network]
}


module "application_gateway" {
  source                   = "../modules/azurerm_application_gateway"
  location                 = var.location
  rg_name                  = var.rg_name
  vnet_name                = var.vnet_name
  create_app_gw            = var.create_app_gw
  create_app_gw_subnet     = var.create_app_gw_subnet
  app_gw_name              = var.app_gw_name
  app_gw_public_ip_name    = var.app_gw_public_ip_name
  app_gw_subnet_cidr_block = var.app_gw_subnet_cidr_block
  app_gw_sku_tier          = var.app_gw_sku_tier
  app_gw_capacity          = var.app_gw_capacity
  app_gw_sku_name          = var.app_gw_sku_name

  depends_on = [module.resource_group, module.virtual_network]
}
