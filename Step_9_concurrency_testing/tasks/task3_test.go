package tasks

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"четное кол-во символов", "parrap", "parrap"},
		{"нечетное кол-во символов", "rapraprap", "parparpar"},
		{"цифары", "123472", "274321"},
		{"пустая строка", "", ""},
		{"юникод смайлики", "😊 Go 🚀", "🚀 oG 😊"},
		{"строка из одного символа", "a", "a"},
		{"cтрока из двух символов", "ab", "ba"},
		{"разный регистр", "ПРивет", "тевиРП"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			res := ReverseString(tt.input)
			if res != tt.output {
				t.Errorf("got %s, want %s", res, tt.output)
			}
		})
	}

}