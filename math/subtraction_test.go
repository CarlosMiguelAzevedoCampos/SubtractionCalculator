package math

import (
	"testing"
)
func TestSubtraction(t *testing.T){
	result:= Subtract(1,2)
	if result != 1{
		t.Errorf("The result was %v", result)
	}
}