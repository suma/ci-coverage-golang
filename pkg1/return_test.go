package pkg1

import (
	"testing"
)

func TestReturnString(t *testing.T) {
	if returnString() != "abc" {
		t.Errorf("it expected 'abc' but %v", returnString())
	}
	errorShouldBeTheLastType()
}
