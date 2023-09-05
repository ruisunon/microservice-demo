terraform {
  backend "azurerm" {
    resource_group_name  = "demo-terraform-state"
    storage_account_name = "terraformstatebarath2022"
    container_name       = "tfstate"
    key                  = "appinfra.eastus2.terraform.tfstate"
    use_azuread_auth     = true
  }
}