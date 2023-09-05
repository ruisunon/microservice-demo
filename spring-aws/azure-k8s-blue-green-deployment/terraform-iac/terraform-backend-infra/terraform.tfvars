rg_name  = "demo-terraform-state"
location = "eastus2"
common_tags = {
  "env" : "prod"
  "costCenter" : "company"
  "accountOwner" : "demouser"
}
storage_account_name   = "terraformstatebarath2022"
storage_container_name = "tfstate"
container_access_type = "private"