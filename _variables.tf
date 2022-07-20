variable "cluster_identifier" {
  type        = string
  description = "DocumentDB Cluster identifier"
}

variable "scheduler_enable" {
  type        = bool
  default     = false
  description = "Whether to schedule the cluster's execution"
}

variable "cron_schedule_start" {
  type        = string
  description = "Cron expression to define when to start the cluster"
}

variable "cron_schedule_stop" {
  type        = string
  description = "Cron expression to define when to stop the cluster"
}