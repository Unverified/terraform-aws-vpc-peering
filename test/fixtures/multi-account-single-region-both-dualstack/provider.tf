provider "aws" {
  alias = "this"

  endpoints {
    ec2 = "http://localhost:4566"
    s3  = "http://localhost:4566"
    sts = "http://localhost:4566"
  }
  region                      = "ap-southeast-2"
  access_key                  = "0${var.test_id}"
  secret_key                  = "null"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
}

provider "aws" {
  alias = "peer"

  endpoints {
    ec2 = "http://localhost:4566"
    s3  = "http://localhost:4566"
    sts = "http://localhost:4566"
  }

  region                      = "ap-southeast-2"
  access_key                  = "1${var.test_id}"
  secret_key                  = "null"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
}
