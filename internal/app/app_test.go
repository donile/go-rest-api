package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApp_GetRoot(t *testing.T) {
	app := Create()

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)

	app.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Errorf("expected '%d' but received '%d'", 200, w.Code)
	}
}
