package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	homeHandler(w, req)

	expected := "Welcome to Tigist's Home Page!\n"
	if w.Body.String() != expected {
		t.Errorf("Expected %q but got %q", expected, w.Body.String())
	}
}
