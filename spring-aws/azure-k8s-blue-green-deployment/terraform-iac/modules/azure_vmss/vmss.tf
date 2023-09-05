
data "azurerm_resource_group" "main" {
  name = var.resource_group_name
}

data "azurerm_subnet" "subnet" {
  name = var.subnet_name
  resource_group_name = data.azurerm_resource_group.main.name
  virtual_network_name = var.vnet_name
}

resource "azurerm_linux_virtual_machine_scale_set" "linux_vmss" {

  count = var.linux_based_vmss ? 1 : 0

  name                = var.vmss_name
  resource_group_name = data.azurerm_resource_group.main.name
  location            = data.azurerm_resource_group.main.location
  sku                 = var.sku
  instances           =  var.vmss_instances
  admin_username      = var.admin_username
  admin_password      =  var.admin_password
  disable_password_authentication = false
  single_placement_group = var.single_placement_group
  overprovision      =   var.overprovision
  source_image_id = var.source_image_id



  os_disk {
    storage_account_type = var.storage_account_type
    caching              = var.storage_account_caching
  }

  network_interface {
    name    =  join("-",[var.vmss_name,"nic"])
    primary = true


    ip_configuration {
      name      = "internal"
      primary   = true
      subnet_id = data.azurerm_subnet.subnet.id

      public_ip_address {
        name = join("-",[var.vmss_name,"public-ip"])
        public_ip_prefix_id = azurerm_public_ip_prefix.public_ip_prefix.id
      }
    }


  }

  tags = var.tags
}

resource "azurerm_public_ip_prefix" "public_ip_prefix" {
  name                = join("-",[var.vmss_name,"public"])
  resource_group_name = data.azurerm_resource_group.main.name
  location            = data.azurerm_resource_group.main.location

  prefix_length = 31

  tags = var.tags
}

resource "azurerm_virtual_machine_scale_set_extension" "vmss_extension" {

  count = var.vmss_extension_enabled ? 1 : 0
  name                         = join("-",[var.vmss_name,"extension"])
  virtual_machine_scale_set_id = azurerm_linux_virtual_machine_scale_set.linux_vmss[count.index].id
  publisher                    = "Microsoft.Azure.Extensions"
  type                         = "CustomScript"
  type_handler_version         = "2.0"
  settings = jsonencode({
    "fileUris" = var.file_uris,
    "commandToExecute" = "bash ./enableagent.sh ${var.azp_url} ${var.azp_agent_name} ${var.azp_pat}   "
  })
}