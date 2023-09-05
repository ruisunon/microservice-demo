***Spring Boot ***
===

##### How to build



Terraform IaC(Infrastructure as Code) to setup Infrastructure on Azure Cloud.

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

- Navigate to terraform directory

```
$ cd infra/terraform
```

- Setup Terraform backend based of Azure Storage Account
```
RESOURCE_GROUP_NAME="demo-production"
STORAGE_ACCOUNT_NAME="demoterraformstate"
ACCOUNT_KEY=$(az storage account keys list --resource-group $RESOURCE_GROUP_NAME --account-name $STORAGE_ACCOUNT_NAME --query '[0].value' -o tsv)
export ARM_ACCESS_KEY=$ACCOUNT_KEY
```

- Initialize the modules

```
$ terraform init
```

- Prepare the terraform values file [terraform.tfvars](terraform.tfvars)
    
  This contains the values to the variables to be considered that includes resource group, location etc.

- Plan the infrastructure

```
$ terraform plan
```

- Apply the infrastructure

```
$ terraform apply
```

- Delete the infrastructure (Handle with care)

```
$ terraform destroy
```

