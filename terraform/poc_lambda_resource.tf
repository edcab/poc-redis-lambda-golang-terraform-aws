data "archive_file" "OnCreate" {
  type        = "zip"
  source_dir = "dist/poc_create"
  output_path = "dist/poc_create.zip"
}

data "archive_file" "OnRetrieve" {
  type        = "zip"
  source_dir = "dist/poc_retrieve"
  output_path = "dist/poc_retrieve.zip"
}