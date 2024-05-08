package httpmodels

import (
	"github.com/go-playground/validator/v10"
)

type CreateItemRequest struct {
	Name        string   `json:"Name" validate:"required"`
	Description string   `json:"Description" validate:"required"`
	MinBuyPrice int      `json:"MinBuyPrice" validate:"required"`
	SoldPrice   int      `json:"SoldPrice" validate:"required,gtefield=MinBuyPrice"`
	Pictures    []string `json:"Pictures" validate:"required"`
}

func (cus *CreateItemRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(cus)
	if err != nil {
		return err
	}

	return nil
}

type GetItemRequest struct {
	Id          int      `json:"Id"`
	Name        string   `json:"Name"`
	Description string   `json:"Description"`
	MinBuyPrice int      `json:"MinBuyPrice"`
	SoldPrice   int      `json:"SoldPrice"`
	Pictures    []string `json:"Pictures"`
}
