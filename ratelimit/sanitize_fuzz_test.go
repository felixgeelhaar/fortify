package ratelimit

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func FuzzSanitizeLogKey(f *testing.F) {
	f.Add("user-123")
	f.Add("")
	f.Add("with\nnewline")
	f.Add("with\x00null")
	f.Add(string([]byte{0x7f}))
	f.Add("multi\nline\rbreaks\ttabs")
	f.Add("非ascii\nuser")

	f.Fuzz(func(t *testing.T, in string) {
		got := sanitizeLogKey(in)

		// Property 1: output never contains control characters.
		for i := 0; i < len(got); i++ {
			c := got[i]
			if c < 32 || c == 127 {
				t.Fatalf("control byte 0x%02x at index %d in output %q (input %q)", c, i, got, in)
			}
		}

		// Property 2: length is preserved (sanitize replaces with '_', not removes).
		if len(got) != len(in) {
			t.Fatalf("length changed: in=%d out=%d (in=%q out=%q)", len(in), len(got), in, got)
		}

		// Property 3: idempotent.
		if got2 := sanitizeLogKey(got); got2 != got {
			t.Fatalf("not idempotent: pass1=%q pass2=%q", got, got2)
		}

		// Property 4: clean ASCII inputs returned unchanged.
		clean := true
		for i := 0; i < len(in); i++ {
			c := in[i]
			if c < 32 || c == 127 {
				clean = false
				break
			}
		}
		if clean && got != in {
			t.Fatalf("clean input mutated: in=%q got=%q", in, got)
		}

		// Property 5: output is valid UTF-8 if input is. (Sanitize uses byte
		// replacement, so multi-byte UTF-8 sequences with no control bytes
		// should pass through unchanged.)
		if utf8.ValidString(in) && !utf8.ValidString(got) {
			// Only fail if the input contained no control bytes — the
			// byte-substitution is allowed to invalidate UTF-8 when it
			// has actually replaced something.
			if !strings.ContainsAny(in, "\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f\x7f") {
				t.Fatalf("output not valid UTF-8: in=%q got=%q", in, got)
			}
		}
	})
}
