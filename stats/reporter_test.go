package stats

import "testing"

func TestAddTagsToName(t *testing.T) {
	tests := []struct {
		name     string
		tags     map[string]string
		expected string
	}{
		{
			name:     "recvd",
			tags:     nil,
			expected: "recvd.no-endpoint.no-os.no-browser",
		},
		{
			name: "recvd",
			tags: map[string]string{
				"endpoint": "advance",
				"os":       "OS X",
				"browser":  "Chrome",
			},
			expected: "recvd.advance.OS-X.Chrome",
		},
		{
			name: "r.call",
			tags: map[string]string{
				"host":     "my-host-name",
				"endpoint": "advance",
				"os":       "OS{}/\tX",
				"browser":  "Chro\\:me",
			},
			expected: "r.call.my-host-name.advance.OS----X.Chro--me",
		},
	}

	for _, tt := range tests {
		got := addTagsToName(tt.name, tt.tags)
		if got != tt.expected {
			t.Errorf("addTagsToName(%v, %v) got %v, expected %v",
				tt.name, tt.tags, got, tt.expected)
		}
	}
}
