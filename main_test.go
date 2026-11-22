package main

import (
	"os"
	"testing"
)

func TestGetMessageDefault(t *testing.T) {
	os.Unsetenv("APP_MESSAGE")
	expected := "Hello, World!"
	actual := getMessage()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

func TestGetMessageCustom(t *testing.T) {
	os.Setenv("APP_MESSAGE", "Kubernetes is awesome")
	defer os.Unsetenv("APP_MESSAGE")

	expected := "Kubernetes is awesome"
	actual := getMessage()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
