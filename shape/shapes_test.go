package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsSquareShape(t *testing.T){
	assert.Implements(t, (*shape)(nil), new(square))
}

func TestIsTriangleShape(t *testing.T){
	assert.Implements(t, (*shape)(nil), new(triangle))
}

func TestGetAreaOfSquare(t *testing.T){
	areaOfSquare := square{sideLen: 5}.getArea()

	assert.Equal(t, float64(25), areaOfSquare)
}

func TestGetAreaOfTriangle(t *testing.T){
	areaOfTriangle := triangle{base: 6, height: 5}.getArea()

	assert.Equal(t, float64(15), areaOfTriangle)
}