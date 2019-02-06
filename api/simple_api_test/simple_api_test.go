package simple

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// START OMIT
func TestTodosHandler(t *testing.T) {
	t.Run("GET /todos", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/todos", nil) // HL
		rr := httptest.NewRecorder()                    // HL
		handler := http.HandlerFunc(todosHandler)

		handler.ServeHTTP(rr, req) // HL

		status := rr.Code
		if status != http.StatusOK { // HL
			t.Errorf("should response status code %v but got %v \n", status, http.StatusOK)
		}

		resp := `hello GET todos`
		if rr.Body.String() != resp { // HL
			t.Errorf("should response body %q but got %q \n", resp, rr.Body.String())
		}
	})
}

// END OMIT
