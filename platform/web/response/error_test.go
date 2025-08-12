package response_test

import (
	"app/platform/web/response"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for Error
func TestError(t *testing.T) {
	t.Run("case 1: should return status code 500 - invalid code", func(t *testing.T) {
		// arrange
		// ...

		// act
		rr := httptest.NewRecorder()
		code := 0
		message := "error message"
		response.Error(rr, code, message)

		// assert
		expectedCode := http.StatusInternalServerError
		expectedBody := `{"status":"Internal Server Error","message":"error message"}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, rr.Code)
		require.Equal(t, expectedBody, rr.Body.String())
		require.Equal(t, expectedHeaders, rr.Header())
	})

	t.Run("case 2: should return status code 400", func(t *testing.T) {
		// arrange
		// ...

		// act
		rr := httptest.NewRecorder()
		code := http.StatusBadRequest
		message := "error message"
		response.Error(rr, code, message)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := `{"status":"Bad Request","message":"error message"}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, rr.Code)
		require.Equal(t, expectedBody, rr.Body.String())
		require.Equal(t, expectedHeaders, rr.Header())
	})

	t.Run("should throw marshal error", func(t *testing.T) {
		realMarshal := response.JsonMarshal
		defer func() { response.JsonMarshal = realMarshal }()

		response.JsonMarshal = func(v interface{}) ([]byte, error) {
			return nil, errors.New("marshal error")
		}

		rr := httptest.NewRecorder()
		response.Error(rr, http.StatusBadRequest, "some error")

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("expected status code 500, got %d", rr.Code)
		}

		if rr.Body.Len() != 0 {
			t.Errorf("expected empty body on marshal error, got %s", rr.Body.String())
		}
	})
}

// Tests for Errorf
func TestErrorf(t *testing.T) {
	t.Run("case 1: should return status code 500 - invalid code", func(t *testing.T) {
		// arrange
		// ...

		// act
		rr := httptest.NewRecorder()
		code := 0
		format := "error message %s"
		args := []interface{}{"arg"}
		response.Errorf(rr, code, format, args...)

		// assert
		expectedCode := http.StatusInternalServerError
		expectedBody := `{"status":"Internal Server Error","message":"error message arg"}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, rr.Code)
		require.Equal(t, expectedBody, rr.Body.String())
		require.Equal(t, expectedHeaders, rr.Header())
	})

	t.Run("case 2: should return status code 400", func(t *testing.T) {
		// arrange
		// ...

		// act
		rr := httptest.NewRecorder()
		code := http.StatusBadRequest
		format := "error message %s"
		args := []interface{}{"arg"}
		response.Errorf(rr, code, format, args...)

		// assert
		expectedCode := http.StatusBadRequest
		expectedBody := `{"status":"Bad Request","message":"error message arg"}`
		expectedHeaders := http.Header{"Content-Type": []string{"application/json"}}
		require.Equal(t, expectedCode, rr.Code)
		require.Equal(t, expectedBody, rr.Body.String())
		require.Equal(t, expectedHeaders, rr.Header())
	})
}
