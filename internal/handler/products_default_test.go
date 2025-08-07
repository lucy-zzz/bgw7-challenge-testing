package handler_test

import (
	"app/internal/handler"
	"app/internal/repository/mocks"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	rp := mocks.NewMockRepositoryProducts(ctrl)

	t.Run("Should search with id", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		ctx := chi.NewRouteContext()
		rp.EXPECT().SearchProducts(gomock.Any())
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
		ctx.URLParams.Add("id", "1")
		expectedStatus := http.StatusOK

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
