package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ksusonic/goshorter/internal/utils"
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
			t.Parallel()

			got := utils.GenerateHash(tt.url)
			assert.Equal(t, tt.expected, got)
		})
	}
}
