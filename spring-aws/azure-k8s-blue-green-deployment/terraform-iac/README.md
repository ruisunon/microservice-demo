


az vmss extension set \
--vmss-name <scaleset name> \
--resource-group <resource group> \
--name CustomScript \
--version 2.0 \
--publisher Microsoft.Azure.Extensions \
--settings '{ \"fileUris\":[\"https://<myGitHubRepoUrl>/myScript.sh\"], \"commandToExecute\": \"bash ./myScript.sh /myArgs \" }'