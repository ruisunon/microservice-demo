output "kv_name" {
  value = azurerm_key_vault.key_vault.name
}

output "kv_sku_name" {
  value = azurerm_key_vault.key_vault.sku_name
}