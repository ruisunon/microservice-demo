locals {

  rg_name     = var.rg_name
  common_tags = var.common_tags

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

module "vmss" {
  source                 = "../modules/azure_vmss"
  resource_group_name    = module.resource_group.rg_name
  vmss_instances         = var.vmss_instances
  vmss_name              = var.vmss_name
  sku                    = var.vmss_sku
  admin_password         = var.admin_password
  admin_username         = var.admin_username
  source_image_id        = var.source_image_id
  subnet_name            = var.vmss_subnet_name
  vnet_name              = var.vnet_name
  tags                   = local.common_tags
  azp_agent_name         = var.azp_agent_name
  azp_pat                = var.azp_pat
  azp_url                = var.azp_url
  file_uris              = var.file_uris
  vmss_extension_enabled = var.vmss_extension_enabled

  depends_on = [module.resource_group, module.virtual_network]
}