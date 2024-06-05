variable "test_id" {
  type = string
  validation {
    condition     = can(regex("^\\d{11}$", var.test_id))
    error_message = "test id must be a 11 character numeric string"
  }
}

variable "azs" {
  description = "Availability Zones"
  type        = list(string)
  default     = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]
}
