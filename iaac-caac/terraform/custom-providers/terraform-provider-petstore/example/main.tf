provider "petstore" {
  address = "https://tcln1rvts1.execute-api.us-west-2.amazonaws.com/v1"
}

resource "petstore_pet" "pet" {
  name    = "snowball"
  species = "cat"
  age     = 8
}

data "petstore_pet_ids" "pets" {
  depends_on = [petstore_pet.pet]
  names      = ["snowball"]
}

output "pet_ids" {
  value = data.petstore_pet_ids.pets.ids
}
