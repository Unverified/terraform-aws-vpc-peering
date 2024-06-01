variable "test_id" {
  type = string
  validation {
    condition     = can(regex("^\\d{11}$", var.test_id))
    error_message = "test id must be a 11 character numeric string"
  }
}

variable "this_vpc_id" {
  type = string
}

variable "peer_vpc_id" {
  type = string
}
