output "resource_group" {
  value = module.resource_group.rg_name
}

output "resource_group_location" {
  value = module.resource_group.rg_location
}

output "kubernetes_cluster_name" {
  value = module.aks.k8s_cluster_name
}

output "app_gw_id" {
  value = module.application_gateway.app_gw_id
}

output "app_gw_name" {
  value = module.application_gateway.app_gw_name
}