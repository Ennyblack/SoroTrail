package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopicMatches(t *testing.T) {
	tests := []struct {
		name   string
		topics json.RawMessage
		topic  json.RawMessage
		want   bool
	}{
		{
			name:   "identical string topic matches",
			topics: json.RawMessage(`["payment"]`),
			topic:  json.RawMessage(`"payment"`),
			want:   true,
		},
		{
			name:   "different string topic does not match",
			topics: json.RawMessage(`["payment"]`),
			topic:  json.RawMessage(`"mint"`),
			want:   false,
		},
		{
			name:   "identical object topic matches",
			topics: json.RawMessage(`[{"type":"payment"}]`),
			topic:  json.RawMessage(`{"type":"payment"}`),
			want:   true,
		},
		{
			name:   "different object topic does not match",
			topics: json.RawMessage(`[{"type":"payment"}]`),
			topic:  json.RawMessage(`{"type":"mint"}`),
			want:   false,
		},
		{
			name:   "whitespace differences do not break matching",
			topics: json.RawMessage(`[{"type":"payment"}]`),
			topic:  json.RawMessage(`{  "type" : "payment" }`),
			want:   true,
		},
		{
			name:   "empty topics returns false",
			topics: json.RawMessage(`[]`),
			topic:  json.RawMessage(`"payment"`),
			want:   false,
		},
		{
			name:   "nil topics returns false",
			topics: nil,
			topic:  json.RawMessage(`"payment"`),
			want:   false,
		},
		{
			name:   "empty topic returns false",
			topics: json.RawMessage(`["payment"]`),
			topic:  json.RawMessage(``),
			want:   false,
		},
		{
			name:   "nil topic returns false",
			topics: json.RawMessage(`["payment"]`),
			topic:  nil,
			want:   false,
		},
		{
			name:   "malformed topics JSON is rejected without panic",
			topics: json.RawMessage(`not-json`),
			topic:  json.RawMessage(`"payment"`),
			want:   false,
		},
		{
			name:   "match found in multi-element array",
			topics: json.RawMessage(`["mint","payment","contract"]`),
			topic:  json.RawMessage(`"payment"`),
			want:   true,
		},
		{
			name:   "topic not present in multi-element array",
			topics: json.RawMessage(`["mint","contract"]`),
			topic:  json.RawMessage(`"payment"`),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topicMatches(tt.topics, tt.topic)
			assert.Equal(t, tt.want, got)
		})
	}
}
