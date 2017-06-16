package operations

import (
	"testing"
)

func TestCircleSpace(t *testing.T) {
	//var expected float32 = 314.159271
	var expected float32 = 314.15
	calculated := CircleSpace(10)
	if expected != calculated {
		t.Errorf("Test Fail : Calculated [%f]\tExpected [%f]\n", calculated, expected)
	}
}

func TestSum(t *testing.T) {
	//expected := 13
	expected := -1
	calculated := Sum(3, 4, 1, 5)
	if expected != calculated {
		t.Errorf("Test Fail : Calculated [%d]\tExptected [%d]\n", calculated, expected)
	}
}
