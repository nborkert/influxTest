package main

import "testing"


func TestGetHTTPCode(t *testing.T) {
	page := "http://www.example.com"
	if x := GetHTTPCode(page); x == "" {
		t.Errorf("Bad responseCode from ", page)
	}
}
