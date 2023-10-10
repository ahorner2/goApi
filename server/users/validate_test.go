package users

import (
	"regexp"
	"testing"
)

func IsValidUsername(username string) bool {
	if len(username) < 4 || len(username) > 20 {
		return false
	}

	isAlphaNumeric := regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString

	return isAlphaNumeric(username)
}

func TestIsValidUsername(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	} {
		{"Al", false},
		{"Alec", true},
		{"alec!", false}, 
		{"Alec H", false}, 
		{"AReallyLongUsernameThatExceedsTwentyCharacters", false},
	}

	for _, test := range tests { 
		got := IsValidUsername(test.input)
		if got != test.want {
			t.Errorf("IsValidUsername(%q) = %v; want %v", test.input, got, test.want)
		}
	}
}

