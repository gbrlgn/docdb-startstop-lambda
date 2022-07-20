data "archive_file" "create_dist_pkg" {
  type        = "zip"
  source_file = "${path.module}/lambda/bin/lambda"
  output_path = "${path.module}/lambda/lambda.zip"
}

resource "aws_lambda_function" "docdb_lambda" {
  function_name = "docdb_lambda"
  role          = aws_iam_role.docdb_lambda_role.arn
  handler       = "lambda"
  filename      = data.archive_file.create_dist_pkg.output_path
  runtime       = "go1.x"
  memory_size   = 128
  timeout       = 15
  environment {
    START_TIME  = split(" ", var.cron_schedule_start)[1]
    STOP_TIME   = split(" ", var.cron_schedule_stop)[1]
  }
}

resource "aws_lambda_permission" "docdb_lambda_start_perm" {
  statement_id  = "AllowExecutionFromCloudWatchStart"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.docdb_lambda.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.docdb_lambda_start.arn
}