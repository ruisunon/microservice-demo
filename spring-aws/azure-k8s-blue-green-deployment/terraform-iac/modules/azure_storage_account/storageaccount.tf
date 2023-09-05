data "azurerm_resource_group" "main" {
  name = var.rg_name
}

resource "azurerm_storage_account" "storage_account" {
  name                     =  var.storage_account_name
  resource_group_name      = data.azurerm_resource_group.main.name
  location                 = data.azurerm_resource_group.main.location
  account_tier             = var.storage_account_tier
  account_replication_type = var.storage_account_replication_type

  tags = var.tags
}

resource "azurerm_storage_container" "storage_account_container" {
  name                  =  var.storage_container_name
  storage_account_name  = azurerm_storage_account.storage_account.name
  container_access_type =  var.container_access_type
}