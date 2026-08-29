package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsType(t *testing.T) {
	tests := []struct {
		name  string
		types []string
		want  string
		match bool
	}{
		{
			name:  "matching type returns true",
			types: []string{"payment", "transfer", "contract"},
			want:  "payment",
			match: true,
		},
		{
			name:  "non-matching type returns false",
			types: []string{"payment", "transfer", "contract"},
			want:  "mint",
			match: false,
		},
		{
			name:  "empty slice returns false",
			types: []string{},
			want:  "payment",
			match: false,
		},
		{
			name:  "nil slice returns false",
			types: nil,
			want:  "payment",
			match: false,
		},
		{
			name:  "matching is exact not substring",
			types: []string{"payment", "transfer"},
			want:  "pay",
			match: false,
		},
		{
			name:  "single element slice matches",
			types: []string{"event"},
			want:  "event",
			match: true,
		},
		{
			name:  "empty want matches empty element",
			types: []string{""},
			want:  "",
			match: true,
		},
		{
			name:  "empty want in non-empty slice returns false",
			types: []string{"payment"},
			want:  "",
			match: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsType(tt.types, tt.want)
			assert.Equal(t, tt.match, got)
		})
	}
}
