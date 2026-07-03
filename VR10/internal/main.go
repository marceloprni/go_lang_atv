package main

import (
	"emailn/internal/domain/campaign"

	"github.com/go-playground/validator/v10"
)

func main() {

	contacts := []campaign.Contact{
		{Email: ""},
	}
	campaign := campaign.Campaign{Contacts: contacts}
	validate := validator.New()
	err := validate.Struct(campaign)

	if err == nil {
		println("Nehum erro")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {

			switch validationError.Tag() {
			case "email":
				println("O campo " + validationError.StructField() + " is email wrong")
			case "min":
				println("O campo " + validationError.StructField() + " e required with min")
			case "max":
				println("O campo " + validationError.StructField() + " e required with max")
			case "required":
				println("O campo " + validationError.StructField() + " e required")
			default:
				println("O campo " + validationError.StructField() + " é inválido")
			}
		}
	}
}
