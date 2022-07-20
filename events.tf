resource "aws_cloudwatch_event_rule" "docdb_lambda_start" {
  name                = "docdb_lambda_start"
  description         = "Start DocumentDB cluster"
  schedule_expression = "cron(${var.cron_schedule_start})"
  is_enabled          = var.scheduler_enable
}

resource "aws_cloudwatch_event_target" "docdb_lambda_start_target" {
  rule      = aws_cloudwatch_event_rule.docdb_lambda_start.name
  target_id = "docdb_lambda_start"
  arn       = aws_lambda_function.docdb_lambda.arn

  input = <<EOI
    {
      "event": "start",
      "db": "${var.cluster_identifier}"
    }
    EOI
}

resource "aws_cloudwatch_event_rule" "docdb_lambda_stop" {
  name                = "docdb_lambda_stop"
  description         = "Stop DocumentDB cluster"
  schedule_expression = "cron(${var.cron_schedule_stop})"
  is_enabled          = var.scheduler_enable
}

resource "aws_cloudwatch_event_target" "docdb_lambda_stop_target" {
  rule      = aws_cloudwatch_event_rule.docdb_lambda_stop.name
  target_id = "docdb_lambda_stop"
  arn       = aws_lambda_function.docdb_lambda.arn

  input = <<EOI
    {
      "event": "stop",
      "db": "${var.cluster_identifier}"
    }
    EOI
}