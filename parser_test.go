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

func TestObject(t *testing.T) {
	t.Run("１階層しかないJsonをパースし、structにマッピングする", func(t *testing.T) {
		type expectedType struct {
			name  string
			value string
		}
		expected := expectedType{"name", "value"}
		inputJson := []byte(`{"name":"name","value":"value}`)
		var got expectedType
		if err := parse(inputJson, got); err != nil {
			t.Error(err)
		}

		if got != expected {
			t.Errorf("Fail: want %v, but got %v\n", expected, got)
		}
	})
}
