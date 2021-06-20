output "db_password" {
  value     = module.database.db_config.password
  sensitive = true
}

output "lb_dns_name" {
  value = module.autoscaling.lb_dns_name
}

output "cloudinit_data" {
  value = module.autoscaling.cloudinit_data
}