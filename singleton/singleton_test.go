package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {

	counter1 := GetInstance()
	if counter1 == nil {
		t.Errorf("Expected pointer to a singleton after calling GetInstance() not nil")
	}

	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("after calliing the count the value should be 1 but got %d", currentCount)
	}

}
