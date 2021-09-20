package color_test

import (
	"testing"

	"github.com/nburmi/maroto/pkg/color"
	"github.com/stretchr/testify/assert"
)

func TestNewWhite(t *testing.T) {
	// Act
	white := color.NewWhite()

	// Assert
	assert.Equal(t, 255, white.Red)
	assert.Equal(t, 255, white.Green)
	assert.Equal(t, 255, white.Blue)
}

func TestColor_IsWhite(t *testing.T) {
	// Act
	white := color.NewWhite()

	// Assert
	assert.True(t, white.IsWhite())
}
