variable "allocation_method" {

  default = "Static"
}

variable "vnet_name" {}

variable "app_gw_name" {}

variable "rg_name" {}

variable "app_gw_subnet_cidr_block" {

}

variable "create_app_gw" {
  default = true
  type = bool
}

variable "create_app_gw_subnet" {
  type = bool
  default = true
}

variable "app_gw_public_ip_name" {}

variable "app_gw_capacity" {}

variable "app_gw_sku_tier" {}

variable "app_gw_sku_name" {}

variable "location" {}

variable "public_ip_sku" {
  default = "Standard"
}