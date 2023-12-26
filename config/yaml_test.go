package config

import "testing"

func TestInterpretYaml(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected *Config
		success  bool
	}{
		{
			name:    "parse valid yaml 1",
			content: "helloWorld: bro...",
			expected: &Config{
				HelloWorld: "bro...",
			},
			success: true,
		},
		{
			name:    "parse invalid yaml 1",
			content: "{'helloWorld': 'bro...'}",
			success: false,
		},
	}

	interpreter := NewYamlInterpreter()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := interpreter.Interpret([]byte(test.content))
			if !test.success && err != nil {
				t.Fail()
			}

			if test.success && *result != *test.expected {
				t.Fail()
			}
		})
	}
}
