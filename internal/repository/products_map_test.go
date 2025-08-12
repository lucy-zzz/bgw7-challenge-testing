package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Get(t *testing.T) {
	t.Run("Should search without id", func(t *testing.T) {
		db := make(map[int]internal.Product)
		rp, _ := repository.NewProductsMap(db).SearchProducts(internal.ProductQuery{})

		assert.Equal(t, rp, db)
	})

	t.Run("Should search with id", func(t *testing.T) {
		db := map[int]internal.Product{1: {Id: 1, ProductAttributes: internal.ProductAttributes{SellerId: 1, Price: 0.1}}, 2: {Id: 2, ProductAttributes: internal.ProductAttributes{SellerId: 1, Price: 0.1}}}
		rp, _ := repository.NewProductsMap(db).SearchProducts(internal.ProductQuery{Id: 1})

		assert.Equal(t, rp[0], db[0])
	})

	t.Run("Should search with id and not find", func(t *testing.T) {
		db := map[int]internal.Product{1: {Id: 1, ProductAttributes: internal.ProductAttributes{SellerId: 1, Price: 0.1}}, 2: {Id: 2, ProductAttributes: internal.ProductAttributes{SellerId: 1, Price: 0.1}}}
		rp, _ := repository.NewProductsMap(db).SearchProducts(internal.ProductQuery{Id: 4})

		assert.Equal(t, rp[0], db[0])
	})
}
