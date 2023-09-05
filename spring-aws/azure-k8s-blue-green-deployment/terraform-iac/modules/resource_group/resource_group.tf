resource "azurerm_resource_group" "resource_group" {
  name         = var.rg_name
  location     = var.location
  tags         = merge({
                    Name = var.rg_name
                 }, var.tags)
}