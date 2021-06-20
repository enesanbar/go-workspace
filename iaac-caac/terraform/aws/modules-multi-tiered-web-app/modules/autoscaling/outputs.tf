output "lb_dns_name" {
  value = module.alb.this_lb_dns_name
}

output "cloudinit_data" {
  value = data.cloudinit_config.config.rendered
}