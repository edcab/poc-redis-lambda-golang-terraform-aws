data "archive_file" "OnCreate" {
  type        = "zip"
  source_dir = "dist/poc_create"
  output_path = "dist/poc_create.zip"
}