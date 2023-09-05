variable "resource_group_name" {

}

variable "kv_name" {

}

variable "location" {

}

variable "enabled_for_disk_encryption" {
  default = true
}

variable "sku_name" {
  default = "standard"
}


variable "soft_delete_retention_days" {
  default = 7
}

variable "purge_protection_enabled" {
  default = false
}

variable "tags" {

}