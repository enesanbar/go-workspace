#resource "aws_acm_certificate" "ecs_domain_certificate" {
#  domain_name       = "*.${var.ecs_domain_name}"
#  validation_method = "DNS"
#
#  tags = {
#    Name = "${var.ecs_cluster_name} Certificate"
#  }
#}
#
#data "aws_route53_zone" "ecs_domain" {
#  name = var.ecs_domain_name
#}
#
#resource "aws_route53_record" "ecs_cert_validation_record" {
#  name            = aws_acm_certificate.ecs_domain_certificate.domain_validation_options.0.resources_record_name
#  type            = aws_acm_certificate.ecs_domain_certificate.domain_validation_options.0.resource_record_type
#  zone_id         = data.aws_route53_zone.ecs_domain.zone_id
#  records         = [aws_acm_certificate.ecs_domain_certificate.domain_validation_options.0.resource_record_value]
#  ttl             = 60
#  allow_overwrite = true
#}
#
#resource "aws_acm_certificate_validation" "ecs_domain_certificate_validation" {
#  certificate_arn         = aws_acm_certificate.ecs_domain_certificate.arn
#  validation_record_fqdns = [aws_route53_record.ecs_cert_validation_record.fqdn]
#}
#
#resource "aws_route53_record" "ecs_load_balancer_record" {
#  name    = "*.${var.ecs_domain_name}"
#  type    = "A"
#  zone_id = data.aws_route53_zone.ecs_domain.zone_id
#
#  alias {
#    evaluate_target_health = false
#    name                   = aws_alb.ecs_cluster_alb.dns_name
#    zone_id                = aws_alb.ecs_cluster_alb.zone_id
#  }
#}