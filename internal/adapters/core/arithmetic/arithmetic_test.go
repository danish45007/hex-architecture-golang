package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()
	ans, err := arith.Addition(1, 1)
	if err != nil {
		t.Fatalf("expected %v but got %v", nil, err)
	}

	require.Equal(t, ans, int32(2))
}

func TestSubtraction(t *testing.T) {
	arith := NewAdapter()
	ans, err := arith.Subtraction(2, 1)
	if err != nil {
		t.Fatalf("expected %v but got %v", nil, err)
	}

	require.Equal(t, ans, int32(1))
}

func TestMultiplication(t *testing.T) {
	arith := NewAdapter()
	ans, err := arith.Multiplication(2, 2)
	if err != nil {
		t.Fatalf("expected %v but got %v", nil, err)
	}

	require.Equal(t, ans, int32(4))
}

func TestDivision(t *testing.T) {
	arith := NewAdapter()
	ans, err := arith.Division(10, 5)
	if err != nil {
		t.Fatalf("expected %v but got %v", nil, err)
	}

	require.Equal(t, ans, int32(2))
}
