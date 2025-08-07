package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	rp := mocks.NewMockRepositoryProducts(ctrl)

	t.Run("Should search without id", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		p := map[int]internal.Product{1: {Id: 1, ProductAttributes: internal.ProductAttributes{SellerId: 1, Price: 0.1}}}
		rp.EXPECT().SearchProducts(gomock.Any()).Return(p, nil)
		expectedStatus := http.StatusOK

		h := handler.NewProductsDefault(rp).Get()
		h.ServeHTTP(w, r)

		assert.Equal(t, expectedStatus, w.Code)
	})

	t.Run("Should search with id", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/?id=1", nil)
		w := httptest.NewRecorder()
		rp.EXPECT().SearchProducts(gomock.Any())
		expectedStatus := http.StatusOK

		h := handler.NewProductsDefault(rp).Get()
		h.ServeHTTP(w, r)

		assert.Equal(t, expectedStatus, w.Code)
	})

	t.Run("Should throw error at repo", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/?id=1", nil)
		w := httptest.NewRecorder()
		rp.EXPECT().SearchProducts(gomock.Any()).Return(nil, errors.New("errouuu"))
		expectedStatus := http.StatusInternalServerError

		h := handler.NewProductsDefault(rp).Get()
		h.ServeHTTP(w, r)

		assert.Equal(t, expectedStatus, w.Code)
	})

	t.Run("Should search with id and id is not a number", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/?id=abc", nil)
		w := httptest.NewRecorder()
		expectedStatus := http.StatusBadRequest

		h := handler.NewProductsDefault(rp).Get()
		h.ServeHTTP(w, r)

		assert.Equal(t, expectedStatus, w.Code)
	})

	// t.Run("Should not find vehicles when trying to get average max speed", func(t *testing.T) {
	// 	r := httptest.NewRequest("GET", "/", nil)
	// 	w := httptest.NewRecorder()
	// 	ctx := chi.NewRouteContext()
	// 	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
	// 	ctx.URLParams.Add("brand", "not found")
	// 	expectedStatus := http.StatusNotFound

	// 	h := handler.NewHandlerVehicle(sv).AverageMaxSpeedByBrand()
	// 	h.ServeHTTP(w, r)

	// 	assert.Equal(t, expectedStatus, w.Code)
	// })

	// t.Run("Should throw internal server when trying to get average max speed", func(t *testing.T) {
	// 	r := httptest.NewRequest("GET", "/", nil)
	// 	w := httptest.NewRecorder()
	// 	ctx := chi.NewRouteContext()
	// 	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
	// 	ctx.URLParams.Add("brand", "err")
	// 	expectedStatus := http.StatusInternalServerError

	// 	h := handler.NewHandlerVehicle(sv).AverageMaxSpeedByBrand()
	// 	h.ServeHTTP(w, r)

	// 	assert.Equal(t, expectedStatus, w.Code)
	// })
}
