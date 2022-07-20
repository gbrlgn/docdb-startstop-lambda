resource "aws_iam_role" "docdb_lambda_role" {
  name = "docdb_lambda"
  assume_role_policy = <<-EOF
  {
    "Version": "2012-10-17",
    "Statement": {
      "Action": "sts:AssumeRole",
      "Principal": { 
        "Service": [
          "lambda.amazonaws.com",
          "rds.amazonaws.com"
        ]
      },
      "Effect": "Allow",
      "Sid": ""
    }
  }
  EOF
}

resource "aws_iam_role_policy" "docdb_role_policy" {
  name   = "docdb_role_policy"
  role   = aws_iam_role.docdb_lambda_role.id
  policy = <<-EOF
  {
    "Version": "2012-10-17",
    "Statement": {
      "Action": [
        "rds:StartDBCluster",
        "rds:StopDBCluster",
        "rds:StartDBInstance",
        "rds:StopDBInstance"
      ],
      "Resource": "*",
      "Effect": "Allow"
    }
  }
  EOF
}