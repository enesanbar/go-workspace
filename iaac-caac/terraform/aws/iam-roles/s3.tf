resource "random_pet" "name" {}

resource "aws_s3_bucket" "b" {
  bucket = "mybucket-${random_pet.name.id}"

  tags = {
    Name = "mybucket-${random_pet.name.id}"
  }
}

resource "aws_s3_bucket_acl" "example" {
  bucket = aws_s3_bucket.b.id
  acl    = "private"
}
