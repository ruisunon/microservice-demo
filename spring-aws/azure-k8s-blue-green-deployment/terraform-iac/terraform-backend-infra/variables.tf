## Resource Group module variables

variable "location" {
  description = "Azure Data Center Location"
}

variable "rg_name" {
  description = "Name of the resource group"
}

variable "storage_account_name" {}
variable "storage_container_name" {}

variable "container_access_type" {}

variable "common_tags" {}