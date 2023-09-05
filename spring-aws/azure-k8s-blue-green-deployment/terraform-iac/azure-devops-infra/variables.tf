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

variable "vmss_subnet_name" {

}
variable "vmss_name" {

}

variable "vmss_extension_enabled" {

}

variable "linux_based_vmss" {
  default = true
  type    = bool
}

variable "vmss_instances" {

}

variable "vmss_sku" {

}

variable "admin_username" {}

variable "admin_password" {}

variable "subnet_name" {

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


variable "common_tags" {

}