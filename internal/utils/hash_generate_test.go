package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateShort(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		url      string
		expected string
	}{
		{
			name:     "success",
			url:      "https://google.com",
			expected: "BQRvJsg-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, GenerateHash(tt.url))
		})
	}
}
