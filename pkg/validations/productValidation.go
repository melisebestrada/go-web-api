package validations

import (
	"fmt"
	"net/http"
	"time"

	"github.com/melisebestrada/go-web-api/pkg/web"
)

func ValidateDate(date string) error {
	layout := "02/01/2006"

	_, err := time.Parse(layout, date)
	if err != nil {
		return fmt.Errorf("wrong date, enter a valid format: DD/MM/AAAA")
	}
	return nil
}

func ValidatedEmptyFields(w http.ResponseWriter, reqBody web.RequestBodyProduct) bool {
	switch {
	case reqBody.Name == "":
		web.SendResponse(w, "Name field is required", nil, true, http.StatusBadRequest)
		return true
	case reqBody.Quantity == 0:
		web.SendResponse(w, "Quantity field is required", nil, true, http.StatusBadRequest)
		return true
	case reqBody.CodeValue == "":
		web.SendResponse(w, "CodeValue field is required", nil, true, http.StatusBadRequest)
		return true
	case reqBody.Expiration == "":
		web.SendResponse(w, "Expiration field is required", nil, true, http.StatusBadRequest)
		return true
	case reqBody.Price == 0:
		web.SendResponse(w, "Price field is required", nil, true, http.StatusBadRequest)
		return true
	}
	return false
}
