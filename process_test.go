package main

import "testing"

func TestProcess(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"reference 1", "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?", "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"},
		{"reference 2", "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure", "I have to pack 5 outfits. Packed 26 just to be sure"},
		{"reference 3", "Don not be sad ,because sad backwards is das . And das not good", "Don not be sad, because sad backwards is das. And das not good"},
		{"reference 4", "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '", "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"},
		{"reference 5", "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68."},
		{"reference 6", "Punctuation tests are ... kinda boring ,what do you think ?", "Punctuation tests are... kinda boring, what do you think?"},
		{"hexadecimal", "1E (hex) files were added", "30 files were added"},
		{"binary", "It has been 10 (bin) years", "It has been 2 years"},
		{"uppercase", "Ready, set, go (up) !", "Ready, set, GO!"},
		{"lowercase", "I should stop SHOUTING (low)", "I should stop shouting"},
		{"capitalize", "Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"several uppercase words", "This is so exciting (up, 2)", "This is SO EXCITING"},
		{"several lowercase words", "ONE TWO THREE (low, 2)", "ONE two three"},
		{"several capitalized words", "hELLO wORLD (cap, 2)", "Hello World"},
		{"attached marker", "word(up)", "WORD"},
		{"attached counted marker", "one two(cap,2)!", "One Two!"},
		{"marker whitespace", "one two(up,\t2)", "ONE TWO"},
		{"marker at start", "(up) hello", "hello"},
		{"only markers", "(up)(cap, 2)", ""},
		{"large count", "go now (up, 9)", "GO NOW"},
		{"overflowing count", "go now (up, 999999999999999999999999999999)", "GO NOW"},
		{"zero count", "HELLO (low, 0)", "HELLO"},
		{"negative count", "HELLO (low, -1)", "HELLO (low, -1)"},
		{"unknown marker", "hello (other)", "hello (other)"},
		{"consecutive markers", "go (up)(low)", "go"},
		{"punctuation not counted", "one , ! two (cap, 2)", "One,! Two"},
		{"attached punctuation not counted", "one,two(up,2)", "ONE, TWO"},
		{"quotes not counted", "' hello world ' (up, 2)", "'HELLO WORLD'"},
		{"hexadecimal with punctuation", "1E, (hex)", "30,"},
		{"binary with quotes", "' 101 ' (bin)", "'5'"},
		{"invalid hexadecimal", "oops (hex)", "oops"},
		{"invalid binary", "102 (bin)", "102"},
		{"number overflow", "FFFFFFFFFFFFFFFF (hex)", "FFFFFFFFFFFFFFFF"},
		{"unicode capitalization", "éLAN (cap)", "Élan"},
		{"unicode case", "été (up) ÜBER (low)", "ÉTÉ über"},
		{"punctuation spacing", "hello ,world ! yes ? fine :okay ;done .", "hello, world! yes? fine: okay; done."},
		{"punctuation groups", "Really !? Yes ... absolutely !!", "Really!? Yes... absolutely!!"},
		{"separated punctuation", "Wait . . . what ! ?", "Wait... what!?"},
		{"punctuation alone", " . ! ? ", ".!?"},
		{"quotes", "they describe me: ' awesome '", "they describe me: 'awesome'"},
		{"several quote pairs", "' hello ' and ' world '", "'hello' and 'world'"},
		{"attached quotes", "He said:'hello.'", "He said: 'hello.'"},
		{"apostrophes", "don't change someone's words", "don't change someone's words"},
		{"apostrophes inside quotes", "' don't stop . '", "'don't stop.'"},
		{"vowels and h", "a apple a egg a idea a orange a umbrella a house", "an apple an egg an idea an orange an umbrella an house"},
		{"uppercase article", "There it was. A amazing rock! A HOUR", "There it was. An amazing rock! An HOUR"},
		{"article before quotes", "a ' amazing ' idea", "an 'amazing' idea"},
		{"article inside quotes", "' a apple '", "'an apple'"},
		{"article after marker", "A (low) apple", "an apple"},
		{"article unchanged", "a car A bike a", "a car A bike a"},
		{"article before punctuation", "a, apple", "a, apple"},
		{"empty input", "", ""},
		{"whitespace", " \t ", ""},
		{"repeated whitespace", "  hello\t world  ", "hello world"},
		{"decimal punctuation rule", "3.14 1,000", "3. 14 1, 000"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := Process(c.in); got != c.want {
				t.Errorf("Process(%q)\n got: %q\nwant: %q", c.in, got, c.want)
			}
		})
	}
}
