data "scaffolding_example" "test" {
  id="example-id"
}

output "test" {
  value = data.scaffolding_example.test
}