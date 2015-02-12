# dynamodbtest

Package for testing golang programs that use DynamoDB.

## Install

	$ go get github.com/groupme/dynamodbtest

## Usage

```go
package foo

import "github.com/groupme/dynamodbtest"

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