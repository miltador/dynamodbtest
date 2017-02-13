# dynamodbtest

Package for testing Go language programs that use DynamoDB.

Runs DynamoDB local server.

## Install

	$ go get github.com/miltador/dynamodbtest

## Usage

```go
package foo

import (
    "github.com/miltador/dynamodbtest"
    "testing"
)

func TestFoo(t *testing.T) {
	// Log output to aid debugging
	dynamodbtest.LogOutput = true

	// Start a new test process
	db, err := dynamodbtest.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Choice of client is up to you, but you will need to point it at db.URL
	client := NewDynamoClient(...)
	client.URL = db.URL()
}

```