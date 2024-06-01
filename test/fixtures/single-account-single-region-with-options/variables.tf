variable "test_id" {
  type = string
  validation {
    condition     = can(regex("^\\d{11}$", var.test_id))
    error_message = "test id must be a 11 character numeric string"
  }
}

variable "this_subnets" {
  description = "Subnet list for _this_ VPC"
  type        = list(string)
  default     = ["172.20.0.0/24", "172.20.1.0/24", "172.20.2.0/24"]
}

variable "peer_subnets" {
  description = "Subnet list for _peer_ VPC"
  type        = list(string)
  default     = ["172.21.0.0/24", "172.21.1.0/24", "172.21.2.0/24"]
}

variable "azs" {
  description = "Availability Zones"
  type        = list(string)
  default     = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]
}
