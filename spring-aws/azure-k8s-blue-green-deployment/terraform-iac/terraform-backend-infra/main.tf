locals {

  rg_name     = var.rg_name
  common_tags = var.common_tags

}


module "rg" {
  source   = "../modules/resource_group"
  location = var.location
  rg_name  = local.rg_name
  tags     = local.common_tags
}

module "storage_account" {
  source                 = "../modules/azure_storage_account"
  rg_name                = module.rg.rg_name
  tags                   = local.common_tags
  storage_account_name   = var.storage_account_name
  storage_container_name = var.storage_container_name
  container_access_type = var.container_access_type

  depends_on = [module.rg]
}