package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDefaultHandler(t *testing.T) {
	body := `{"words":["apple", "ant", "ape", "cat", "car", "cart", "curtsy", "dog", "dainty", "dart"],"prefix":""}`
	request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(defaultHandler)
	handler.ServeHTTP(responseRecorder, request)
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"errcode":0,"message":"OK","data":["dog","dart","dainty","curtsy","cat","car","cart","apple","ape","ant"]}`
	got := responseRecorder.Body.String()
	got = strings.ReplaceAll(got, "\n", "")
	if got != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseRecorder.Body.String(), expected)
	}
}

func TestBuddyStringHandler(t *testing.T) {
	body := `{"s":"ab","goal":"ba"}`
	request := httptest.NewRequest(http.MethodPost, "/buddystring", strings.NewReader(body))
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(buddyStringHandler)
	handler.ServeHTTP(responseRecorder, request)
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := "true"
	got := responseRecorder.Body.String()
	if got != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseRecorder.Body.String(), expected)
	}
}

func TestDiagonalHandler(t *testing.T) {
	body := `{"td":[[1,2,3],[4,5,6],[7,8,9]]}`
	request := httptest.NewRequest(http.MethodPost, "/diagonal", strings.NewReader(body))
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(diagonalHandler)
	handler.ServeHTTP(responseRecorder, request)
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[1,4,2,7,5,3,8,6,9]`
	got := responseRecorder.Body.String()
	if got != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			responseRecorder.Body.String(), expected)
	}
}
