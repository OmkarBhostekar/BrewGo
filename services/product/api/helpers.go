package api

import (
	"github.com/omkarbhostekar/brewgo/proto/gen"
	db "github.com/omkarbhostekar/brewgo/services/product/db/sqlc"
)

func convertProductModelToProto(product db.Product) *gen.Product {
	return &gen.Product{
		ProductId:   product.ID,
		Name:        product.Name,
		Description: product.Description,
		Category:    product.Category,
		Price:       float32(product.Price.InexactFloat64()),
		ItemType:    product.ItemType,
		IsAvailable: product.IsAvailable,
	}
}