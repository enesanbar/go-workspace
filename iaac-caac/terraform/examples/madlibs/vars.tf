variable "words" {
  description = "A word pool to use for Mad Libs"
  type = object({
    nouns      = list(string),
    adjectives = list(string),
    verbs      = list(string),
    adverbs    = list(string),
    numbers    = list(number),
  })

  validation {
    condition = length(var.words["nouns"]) >= 12
    error_message = "At least 20 nouns must be supplied"
  }
}

variable "num_files" {
  default = 100
  type = number
}