package main

import "testing"

func TestGetHeadChar(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected rune
		wantErr  bool
	}{
		{"should return head character of given line", "abc", 'a', false},
		{"should return err when encounter EOS", "", 0, true},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getHeadChar(tt.input)
			if tt.wantErr && err == nil {
				t.Error("should return err, but didnt")
			}

			if tt.expected != got {
				t.Errorf("should return %d, but got %d\n", tt.expected, got)
			}
		})
	}
}
