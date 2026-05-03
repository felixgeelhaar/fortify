package http

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func FuzzSanitizeKey(f *testing.F) {
	f.Add("user-123", 256)
	f.Add("", 256)
	f.Add("with\nnewline", 100)
	f.Add("with\x00null", 100)
	f.Add(string([]byte{0x7f}), 100)
	f.Add("ünïcödé", 100)
	f.Add("very long key that exceeds maxLen", 5)
	f.Add(strings.Repeat("ä", 50), 10)

	f.Fuzz(func(t *testing.T, in string, maxLen int) {
		// Bound maxLen to a sane range so the test doesn't OOM.
		if maxLen < 0 {
			maxLen = -maxLen
		}
		if maxLen > 1024 {
			maxLen = 1024
		}

		got := SanitizeKey(in, maxLen)

		// Property 1: output is valid UTF-8 (NFC normalization + Unicode
		// rune-aware truncation guarantees this).
		if !utf8.ValidString(got) {
			t.Fatalf("not valid UTF-8: in=%q got=%q maxLen=%d", in, got, maxLen)
		}

		// Property 2: rune count never exceeds the requested cap (after
		// defaulting). When maxLen is 0, the function uses DefaultMaxKeyLength.
		effectiveMax := maxLen
		if effectiveMax <= 0 {
			effectiveMax = DefaultMaxKeyLength
		}
		if utf8.RuneCountInString(got) > effectiveMax {
			t.Fatalf("rune count %d > maxLen %d (in=%q got=%q)",
				utf8.RuneCountInString(got), effectiveMax, in, got)
		}

		// Property 3: no control characters in output.
		for _, r := range got {
			if r < 32 || r == 127 {
				t.Fatalf("control rune %d in output %q (in=%q)", r, got, in)
			}
		}

		// Property 4: idempotent (calling SanitizeKey on its own output
		// produces the same string).
		got2 := SanitizeKey(got, maxLen)
		if got2 != got {
			t.Fatalf("not idempotent: pass1=%q pass2=%q", got, got2)
		}
	})
}
