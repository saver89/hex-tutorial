package arithmetics

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Addition(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(2))
}

func TestSubstraction(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Substruction(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(0))
}

func TestMultiplication(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Mutliplication(2, 3)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(6))
}

func TestDivision(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Division(4, 2)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(2))
}
