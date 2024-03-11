// to run the test, use "go test -v ."

package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}