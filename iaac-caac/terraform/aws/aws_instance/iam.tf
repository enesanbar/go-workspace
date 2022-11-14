resource "aws_iam_role" "iam-read-only-role" {
  name               = "iam-read-only-role"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF

  tags = merge(
    {
      "Name" = "${local.prefix}-iam-read-only-role"
    },
    local.common_tags,
  )
}

resource "aws_iam_instance_profile" "iam-read-only-role-instanceprofile" {
  name = "iam-read-only-instance-profile"
  role = aws_iam_role.iam-read-only-role.name

  tags = merge(
    {
      "Name" = "${local.prefix}-iam-read-only-instance-profile"
    },
    local.common_tags,
  )
}

resource "aws_iam_role_policy" "iam-read-only-role-policy" {
  name   = "iam-read-only-role-policy"
  role   = aws_iam_role.iam-read-only-role.id
  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
              "iam:Get*",
              "iam:List*"
            ],
            "Resource": "*"
        }
    ]
}
EOF

}

