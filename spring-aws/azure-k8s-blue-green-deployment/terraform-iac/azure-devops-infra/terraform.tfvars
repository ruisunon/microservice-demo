location         = "westus2"
rg_name          = "demo-azure-devops-west-coast"
vnet_name        = "demo-azure-devops-west-coast-vnet"
vnet_cidr_blocks = ["10.20.0.0/16"]
subnet_cidr_blocks = [
  "10.20.0.0/23",
  "10.20.2.0/23",
  "10.20.4.0/23",
  "10.20.10.0/23",
  "10.20.12.0/23",
  "10.20.14.0/23"
]
subnet_names = [
  "demo-azure-devops-private-subnet-0",
  "demo-azure-devops-private-subnet-1",
  "demo-azure-devops-private-subnet-2",
  "demo-azure-devops-public-subnet-0",
  "demo-azure-devops-public-subnet-1",
  "demo-azure-devops-public-subnet-2"
]

vmss_sku       = "Standard_DC2s_v2"
vmss_name      = "demo-azure-devops-agent"
vmss_instances = 2

vmss_extension_enabled = true
linux_based_vmss       = true
source_image_id        = "/subscriptions/da149aec-98ab-4deb-b618-603896c291f0/resourceGroups/demo-terraform-state/providers/Microsoft.Compute/images/devops-agent-image-20211109231022"

common_tags = {
  "env" : "prod"
  "costCenter" : "company"
  "accountOwner" : "demouser"
}

file_uris      = ["https://vstsagenttools.blob.core.windows.net/tools/ElasticPools/Linux/6/enableagent.sh"]
azp_agent_name = "Docker"
azp_pat        = "3fpnos6madm6avvej3ouijumfdxb3t6otldpeex4642kbwez36za"
azp_url        = "https://dev.azure.com/barathtrainer2022"

admin_username   = "barathtrainer"
admin_password   = "Barath123456"
subnet_name      = "demo-azure-devops-public-subnet-0"
vmss_subnet_name = "demo-azure-devops-public-subnet-0"