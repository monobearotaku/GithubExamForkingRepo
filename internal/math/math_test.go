package math

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestSinus(t *testing.T) {
	assert.Equal(t, math.Sin(1234.5678), Sinus(1234.5678))
}

func TestCosinus(t *testing.T) {
	assert.Equal(t, math.Cos(1234.5678), Cosinus(1234.5678))
}

func TestTangens(t *testing.T) {
	assert.Equal(t, math.Tan(1234.5678), Tangens(1234.5678))
}
