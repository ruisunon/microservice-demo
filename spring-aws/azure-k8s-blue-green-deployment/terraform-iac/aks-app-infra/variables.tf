## Resource Group module variables

variable "location" {
  description = "Azure Data Center Location"
}

variable "rg_name" {
  description = "Name of the resource group"
}

## VirtualNetwork module variables

variable "vnet_name" {
  description = "Name of the virtual network"
}

variable "vnet_cidr_blocks" {
  description = "CIDR block of the virtual network"
}

variable "subnet_names" {
  description = "Name of the subnet groups"
}

variable "subnet_cidr_blocks" {
  description = "CIDR blocks of the subnet groups to be created"
}

## AKS module variables

variable "k8s_cluster_name" {
  description = "AKS kubernetes cluster name"
}

variable "k8s_cluster_version" {
  description = "AKS kubernetes cluster version"
}

variable "network_plugin" {
  description = "AKS kubernetes network plugin to be used"
  default     = "azure"
}

variable "network_policy" {
  description = "AKS kubernetes network policy to be used"
  default     = "azure"
}

variable "rbac_aad_managed" {
  description = "AKS kubernetes RBAC Managed"
  default     = true
}

variable "k8s_sku_tier" {
  description = "AKS kubernetes SKU"
  default     = "Free"
}


variable "k8s_vnet_subnet_name" {
  description = "Name of the k8s subnet to lookup"
}

variable "private_cluster_enabled" {
  description = "AKS kubernetes private cluster flag"
  default     = false
}

variable "enable_http_application_routing" {
  default = false
}

variable "enable_azure_policy" {
  default = true
}

variable "enable_auto_scaling" {
  description = "Enable Log analytics workspace"
  default     = true
}

variable "enable_host_encryption" {
  description = "Enable Log analytics workspace"
  default     = true
}



variable "k8s_prefix" {

}

variable "rbac_aad_admin_group_object_ids" {
  description = "Active Directory Group Admin Object IDs"
  default     = []
}


variable "enable_log_analytics_workspace" {
  description = "Enable Log analytics workspace"
  default     = false
}

variable "enable_kube_dashboard" {
  description = "Enable Kubernetes Dashboard"
  default     = true
}


variable "default_node_pool" {

  description = "Default AKS nodepool to be created"
  type = map(object({
    name                   = string,
    vm_size                = string,
    node_count             = number,
    min_count              = number,
    max_count              = number,
    max_pods               = number,
    availability_zones     = list(string),
    enable_auto_scaling    = bool,
    enable_host_encryption = bool,
    mode                   = string,
    node_taints            = list(string),
    node_labels            = map(string),
    kubernetes_version     = string,
    os_type                = string,
    type                   = string,
    os_disk_type           = string,
    os_disk_size_gb        = number,
    priority               = string
  }))
}
variable "k8s_node_pools" {
  description = "Additional AKS nodepools to be created"
  type = map(object({
    name                   = string,
    vm_size                = string,
    node_count             = number,
    min_count              = number,
    max_count              = number,
    max_pods               = number,
    availability_zones     = list(string),
    enable_auto_scaling    = bool,
    enable_host_encryption = bool,
    mode                   = string,
    node_taints            = list(string),
    node_labels            = map(string),
    kubernetes_version     = string,
    os_type                = string,
    os_disk_type           = string,
    os_disk_size_gb        = number,
    priority               = string
  }))
}

### Application Gateway Variables

variable "create_app_gw_subnet" {
  type    = bool
  default = true
}

variable "create_app_gw" {
  default = true
  type    = bool
}
variable "app_gw_name" {

}

variable "app_gw_public_ip_name" {

}

variable "app_gw_subnet_cidr_block" {

}

variable "app_gw_sku_tier" {

}

variable "app_gw_capacity" {}

variable "app_gw_sku_name" {}

variable "net_profile_dns_service_ip" {
  default = null
}

variable "net_profile_outbound_type" {
  default = null
}

variable "net_profile_pod_cidr" {
  default = null
}
variable "enable_role_based_access_control" {
  default = true
}
variable "net_profile_service_cidr" {
  default = null
}

variable "net_profile_docker_bridge_cidr" {
  default = null
}

## ACR Variables

variable "acr_name" {

}

variable "acr_sku" {

}

## Common Tags
variable "common_tags" {
  description = "Common tags to be added to all the resources"
}