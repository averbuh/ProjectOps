package main

import (
	"net/http"
	"testing"
)

func TestRecipesHandler_ServeHTTP(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *RecipesHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.ServeHTTP(tt.args.w, tt.args.r)
		})
	}
}
