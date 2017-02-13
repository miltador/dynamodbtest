package dynamodbtest

import (
	"fmt"
	"testing"
)

// TestAll tests public API
func TestAll(t *testing.T) {
	// Log output to aid debugging
	LogOutput = true

	// Start a new test process
	db, err := New()
	if err != nil {
		t.Fatal(err)
	}

	url := db.URL()
	if url != "http://localhost:8000" {
		t.Error("URL is not correct")
	}

	_err := db.Close()
	if _err != nil {
		t.Fatal(_err)
	}
}

func Example() {
	// Log output to aid debugging
	LogOutput = true

	// Start a new test process
	db, err := New()
	if err != nil {
		fmt.Errorf("Couldn't start DynamoDB local, %v", err)
	}
	defer db.Close()

	// Choice of client is up to you, but you will need to point it at db.URL
	//client := NewDynamoClient(...)
	//client.URL = db.URL()
}
