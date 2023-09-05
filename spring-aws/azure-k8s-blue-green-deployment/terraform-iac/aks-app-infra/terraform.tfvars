location         = "westus2"
rg_name          = "demo-west-coast"
vnet_name        = "demo-west-coast-vnet"
vnet_cidr_blocks = ["10.10.0.0/16"]
subnet_cidr_blocks = [
  "10.10.0.0/23",
  "10.10.2.0/23",
  "10.10.4.0/23",
  "10.10.10.0/23",
  "10.10.12.0/23",
  "10.10.14.0/23",
  "10.10.22.0/26",
  "10.10.96.0/20"
]
subnet_names = [
  "demo-private-subnet-0",
  "demo-private-subnet-1",
  "demo-private-subnet-2",
  "demo-public-subnet-0",
  "demo-public-subnet-1",
  "demo-public-subnet-2",
  "demo-ingress-gateway-subnet",
  "demo-app-k8s-private-cluster-subnet"
]

k8s_vnet_subnet_name = "demo-app-k8s-private-cluster-subnet"

k8s_cluster_version = "1.21.2"
k8s_cluster_name    = "demo-west-coast-k8s"

k8s_prefix                       = "demo-k8s"
k8s_sku_tier                     = "Free"
private_cluster_enabled          = false
enable_log_analytics_workspace   = false
enable_azure_policy              = true
enable_kube_dashboard            = false
enable_role_based_access_control = true




default_node_pool = {
  default_pool = {
    name                   = "syspool1",
    vm_size                = "Standard_DS2_v2",
    node_count             = 1,
    max_count              = 1,
    min_count              = 1,
    max_pods               = 50,
    availability_zones     = ["1"],
    enable_auto_scaling    = true,
    enable_host_encryption = false,
    mode                   = "System",
    node_taints            = null,
    node_labels            = { "env" : "prod", "platform" : "apps" },
    kubernetes_version     = "1.21.2",
    os_type                = "Linux",
    type                   = "VirtualMachineScaleSets",
    os_disk_type           = "Managed",
    os_disk_size_gb        = 100,
    priority               = "Spot"
  }
}


k8s_node_pools = {
  pool1 = {
    name                   = "pool1",
    vm_size                = "Standard_D32as_v5",
    node_count             = 1,
    max_count              = 1,
    min_count              = 1,
    max_pods               = 50,
    availability_zones     = ["2"],
    enable_auto_scaling    = true,
    enable_host_encryption = false,
    mode                   = "User",
    node_taints            = ["env=demo:NoSchedule"],
    node_labels            = { "env" : "demo", "platform" : "apps" },
    kubernetes_version     = "1.21.2",
    os_type                = "Linux",
    os_disk_type           = "Managed",
    os_disk_size_gb        = 100,
    priority               = "Spot"
  }
}



## Application Gateway Values
create_app_gw_subnet     = false
create_app_gw            = true
app_gw_name              = "demo-ingress-gateway"
app_gw_public_ip_name    = "demo-west-coast-ip"
app_gw_sku_tier          = "Standard_v2"
app_gw_sku_name          = "Standard_v2"
app_gw_capacity          = 1
app_gw_subnet_cidr_block = "10.10.22.0/26"

## ACR

acr_name = "demoaksbarath2022"
acr_sku  = "Standard"

## Common Tags
common_tags = {
  customer    = "demo"
  env         = "demo"
  environment = "demo"
  dc          = "west-coast"
  costcenter  = "EASTUS2"
}