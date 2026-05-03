package grpc

import (
	"strings"
	"testing"
)

func FuzzSanitizeMetadataKey(f *testing.F) {
	f.Add("api-key-abc123")
	f.Add("")
	f.Add("with\nnewline")
	f.Add("with\x00null")
	f.Add(strings.Repeat("a", 1024)) // long input
	f.Add("non-ascii-éüñ")

	f.Fuzz(func(t *testing.T, in string) {
		got := sanitizeMetadataKey(in)

		// Property 1: length capped at metadataKeyMaxLen.
		if len(got) > metadataKeyMaxLen {
			t.Fatalf("output length %d > cap %d (in=%q got=%q)", len(got), metadataKeyMaxLen, in, got)
		}

		// Property 2: no control bytes (ASCII < 32 or 127).
		for i := 0; i < len(got); i++ {
			c := got[i]
			if c < 32 || c == 127 {
				t.Fatalf("control byte 0x%02x at %d in output %q (in=%q)", c, i, got, in)
			}
		}

		// Property 3: idempotent.
		if got2 := sanitizeMetadataKey(got); got2 != got {
			t.Fatalf("not idempotent: pass1=%q pass2=%q", got, got2)
		}

		// Property 4: clean ASCII inputs of OK length pass through unchanged.
		if len(in) <= metadataKeyMaxLen {
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
		}
	})
}
