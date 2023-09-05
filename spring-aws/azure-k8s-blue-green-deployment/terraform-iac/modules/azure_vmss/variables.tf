variable "resource_group_name" {

}

variable "vmss_name" {
    
}

variable "linux_based_vmss" {
  default = true
  type = bool
}

variable "vmss_instances" {

}

variable "sku" {

}

variable "file_uris" {
  default = []
}

variable "azp_url" {
  default = ""
}

variable "azp_pat" {
  default = ""
}
variable "azp_agent_name" {
  default = ""
}

variable "single_placement_group" {
  default = false
}

variable "overprovision" {
  default = false
}

variable "tags" {}

variable "admin_username" {}

variable "admin_password" {}

variable "vnet_name" {}

variable "subnet_name" {

}

variable "vmss_extension_enabled" {

}

variable "source_image_reference" {
  default = {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "16.04-LTS"
    version   = "latest"
  }
}

variable "source_image_id" {

}

variable "storage_account_type" {
  default = "Standard_LRS"
}

variable "storage_account_caching" {
  default = "ReadWrite"
}