package tests

import (
	"automated_tests/math"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	// arrange
	t.Helper()
	want := 10

	// act
	got := math.Add(4, 6)

	// assert
	require.Equal(t, got, want)	
}

func TestSubtract(t *testing.T) {
	// arrange
	t.Helper()
	want := 2
	
	// act
	got := math.Subtract(6, 4)
	
	// assert
	require.Equal(t, got, want)
}
