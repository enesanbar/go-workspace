resource "tls_private_key" "key" {
  algorithm = "RSA"
}

resource "local_sensitive_file" "foo" {
  content  = tls_private_key.key.private_key_pem
  filename = "${path.module}/ansible-key.pem"
  file_permission = "0400"
}

resource "aws_key_pair" "key_pair" {
  key_name = "ansible-key"
  public_key = tls_private_key.key.public_key_openssh
}
