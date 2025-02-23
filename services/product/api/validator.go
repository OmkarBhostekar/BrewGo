package api

import (
	"fmt"

	"github.com/omkarbhostekar/brewgo/proto/gen"
)

func validateAddProductRequest(req *gen.AddProductRequest) error {
	if req.Name == "" {
		return fmt.Errorf("name is required")
	}
	if req.Description == "" {
		return fmt.Errorf("description is required")
	}
	if req.Category == "" {
		return fmt.Errorf("category is required")
	}
	if req.Price == 0 {
		return fmt.Errorf("price is required")
	}
	if req.ItemType == "" {
		return fmt.Errorf("item type is required")
	}
	return nil
}