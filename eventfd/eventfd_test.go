package eventfd

import (
	"testing"
)

func TestEventFDBasic(t *testing.T) {
	// Create new eventfd
	efd, err := Create()
	if err != nil {
		t.Fatalf("Failed to create eventfd: %v", err)
	}
	defer efd.Close()

	// Test initial value
	val, err := efd.Read()
	if err != nil {
		t.Errorf("Failed to read initial value: %v", err)
	}
	if val != 0 {
		t.Errorf("Expected initial value 0, got %d", val)
	}
}
