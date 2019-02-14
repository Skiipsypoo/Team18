package ascii 

import (
	"testing"
	"unicode"
)


func TestGreetingsASCII(t *testing.T) {
	var str string 
	str =	GreetingASCII()
	for _, c := range str {
		if c > unicode.MaxASCII {
			t.Error("Is not ASCII")
		}
	}
}