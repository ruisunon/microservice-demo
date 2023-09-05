output "storage_account_name" {
  value = module.storage_account.azurerm_storage_account_name
}

output "resource_group_name" {
  value = var.rg_name
}