package eventfd

import (
	"testing"
)

func TestEventFDNonBlocking(t *testing.T) {
	// Create new non-blocking eventfd
	efd, err := Create()
	if err != nil {
		t.Fatalf("Failed to create non-blocking eventfd: %v", err)
	}
	defer efd.Close()

	// Test initial read (should not block and return 0)
	val, err := efd.Read()
	if err != nil {
		t.Errorf("Failed to read initial value: %v", err)
	}
	if val != 0 {
		t.Errorf("Expected initial value 0, got %d", val)
	}

	// Write a value and read it back
	err = efd.Write(42)
	if err != nil {
		t.Fatalf("Failed to write value: %v", err)
	}

	val, err = efd.Read()
	if err != nil {
		t.Errorf("Failed to read value after write: %v", err)
	}
	if val != 42 {
		t.Errorf("Expected value 42, got %d", val)
	}

	// Test read with no data available (should not block and return 0)
	val, err = efd.Read()
	if err != nil {
		t.Errorf("Failed to read value when no data is available: %v", err)
	}
	if val != 0 {
		t.Errorf("Expected value 0 when no data is available, got %d", val)
	}
}
