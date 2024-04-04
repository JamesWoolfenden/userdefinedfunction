data "scaffolding_example" "test" {
}

output "test" {
  value = data.scaffolding_example.test
}


#output "function" {
#  value = provider::scaffolding::example("testvalue")
#}

output "test2" {
  value = provider::scaffolding::echo("hello world")
}