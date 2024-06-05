variable "test_id" {
  type = string
  validation {
    condition     = can(regex("^\\d{11}$", var.test_id))
    error_message = "test id must be a 11 character numeric string"
  }
}

variable "azs_this" {
  description = "Availability Zones for requester VPC"
  type        = list(string)
  default     = ["ap-southeast-2a", "ap-southeast-2b", "ap-southeast-2c"]
}

variable "azs_peer" {
  description = "Availability Zones for accepter VPC"
  type        = list(string)
  default     = ["ap-southeast-2a", "ap-southeast-2b", "ap-southeast-2c"]
}
