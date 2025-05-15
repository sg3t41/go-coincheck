package client

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := New("", "")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if client == nil {
		t.Errorf("expected client, got nil")
	}
}
