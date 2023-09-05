***demo IaC Terraform***
===

Terraform IaC(Infrastructure as Code) to setup demo Infrastructure on Azure Cloud.

##### Compatability Matrix

choose the branch based on below maintained versions.

<table>
 <tr>
    <th style="text-align:left">Provider</th>
    <th style="text-align:left">Version</th>
  </tr>
  <tr>
    <td>Terraform </td>
    <td> v1.0.0 </td>
  </tr>  
  <tr>
    <td>AzureRM</td>  
    <td>v2.78.0</td>
  </tr>  
</table>


##### Tools

* [Terraform](https://www.terraform.io/downloads.html)
* [Azure CLI](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli)

##### Terraform Instructions

- Setup Terraform backend based of Azure Storage Account
```
export RESOURCE_GROUP_NAME="demo-terraform-state"
export STORAGE_ACCOUNT_NAME="terraformstatebarath2022"
ACCOUNT_KEY=$(az storage account keys list --resource-group $RESOURCE_GROUP_NAME --account-name $STORAGE_ACCOUNT_NAME --query '[0].value' -o tsv)
export ARM_ACCESS_KEY=$ACCOUNT_KEY
```

- Initialize the modules

```
$ terraform init
```

- Prepare the terraform values file [westus2.tfvars](westus2.tfvars)
    
  This contains the values to the variables to be considered that includes resource group, location etc.

- Plan the infrastructure

```
$ terraform plan --var-file=westus22.tfvars --out westus22
```

- Apply the infrastructure

```
$ terraform apply westus22
```

- Delete the infrastructure (Handle with care)

```
$ terraform destroy --var-file=westus22.tfvars
```

Notes:

az feature register --namespace "Microsoft.Compute" --name "EncryptionAtHost"


az feature list -o table --query "[?contains(name, 'Microsoft.Compute/EncryptionAtHost')].{Name:name,State:properties.state}"