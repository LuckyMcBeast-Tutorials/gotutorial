package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOddEven(t *testing.T){
	newOddEven := newOddEven()

	assert.Equal(t, 11, len(newOddEven))
	assert.ElementsMatch(t, []int(newOddEven), []int{0,1,2,3,4,5,6,7,8,9,10})
}

func TestIsEven(t *testing.T){
	isEightEven := isEven(8)
	isThreeEven := isEven(3)

	assert.Equal(t, true, isEightEven)
	assert.Equal(t, false, isThreeEven)
}

func TestOddOrEvenOutput(t *testing.T){
	eightOutput := oddOrEvenOutput(8)
	threeOutput := oddOrEvenOutput(3)

	assert.Equal(t, "8 is even", eightOutput)
	assert.Equal(t, "3 is odd", threeOutput)
}
