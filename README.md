# docdb-startstop-lambda

This Terraform module deploys an AWS Lambda function for starting and stopping a DocumentDB cluster at a set time.

The following resources will be created:
- AWS Lambda function
- CloudWatch:
  - Event rule
  - Event target
- IAM role

## Usage
Usage example with DocumentDB cluster.
```hcl
module "docdb_startstop_lambda" {
  for_each            = { for doc_db in local.workspace.documentdb.doc_dbs : doc_db.name => doc_db }
  source              = "git::https://github.com/gbrlgn/docdb-startstop-lambda.git"
	
  cluster_identifier  = each.value.name
  schedule_enable     = true
  cron_schedule_start = "0 9 * * MON-FRI *"
  cron_schedule_stop  = "0 23 * * MON-FRI *"
}
```
