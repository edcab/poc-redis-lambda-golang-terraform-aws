# Cloud
provider "aws" {
    version = "~> 3.0"
    region  = "us-east-1"
}

provider "archive" {}

# local
# intento fallido, local stack no soporta api gw tipo
// terraform {
//   backend "local" {}
// }
//
// provider "aws" {
//   access_key                  = "mock_access_key"
//   region                      = "us-east-1"
//   s3_force_path_style         = true
//   secret_key                  = "mock_secret_key"
//   skip_credentials_validation = true
//   skip_metadata_api_check     = true
//   skip_requesting_account_id  = true
//
//   endpoints {
//     s3          = "http://0.0.0.0:4566"
//     lambda      = "http://0.0.0.0:4566"
//     iam         = "http://0.0.0.0:4566"
//     apigateway  = "http://0.0.0.0:4566"
//     dynamodb    = "http://0.0.0.0:4566"
//     redis       = "http://0.0.0.0:4566"
//   }
// }