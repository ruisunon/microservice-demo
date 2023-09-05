output "app_gw_id" {
  value = azurerm_application_gateway.application_gateway.*.id
}

output "app_gw_name" {
  value = azurerm_application_gateway.application_gateway.*.name
}