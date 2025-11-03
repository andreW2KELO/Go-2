package tasks

import "testing"

func TestAreAnagrams(t *testing.T) {
    tests := []struct {
        str1, str2 string
        want       bool
        name       string
    }{
        {"listen", "silent", true, "простая анаграмма"},
        {"Listen", "Silent", true, "анаграмма с разным регистром"},
        {"evil", "vile", true, "анаграмма коротких слов"},
        {"triangle", "integral", true, "анаграмма длинных слов"},
        {"apple", "papel", true, "анаграмма со схожими буквами"},
        {"rat", "car", false, "разные слова"},
        {"abc", "abcc", false, "разная длина"},
        {"hello", "holle", true, "неанаграмма с одинаковой длиной"},
        {"", "", true, "две пустые строки"},
        {"a", "A", true, "одна буква в разном регистре"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := AreAnagrams(tt.str1, tt.str2)
            if got != tt.want {
                t.Errorf("AreAnagrams(%q, %q) = %v; want %v", tt.str1, tt.str2, got, tt.want)
            }
        })
    }
}