package base62

import (
	"math"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	testCases := []struct {
		input uint64
		want  string
	}{
		{0, "0"},
		{1, "1"},
		{10, "A"},
		{35, "Z"},
		{36, "a"},
		{61, "z"},
		{62, "10"},
		{1000, "g8"},
		{999999, "4c91"},
		{math.MaxUint64, "LygHa16AHYF"},
	}

	for _, tc := range testCases {
		t.Run(tc.want, func(t *testing.T) {
			encoded := Encode(tc.input)
			if encoded != tc.want {
				t.Errorf("Encode(%d) = %s; want %s", tc.input, encoded, tc.want)
			}

			decoded, err := Decode(encoded)
			if err != nil {
				t.Errorf("Unexpected error decoding %s: %v", encoded, err)
			}
			if decoded != tc.input {
				t.Errorf("Decode(%s) = %d; want %d", encoded, decoded, tc.input)
			}
		})
	}
}

func TestDecodeInvalid(t *testing.T) {
	invalidInputs := []string{
		"",
		" ",
		"!",
		"ABC!123",
		"你好",
	}

	for _, input := range invalidInputs {
		t.Run(input, func(t *testing.T) {
			_, err := Decode(input)
			if err == nil {
				t.Errorf("Expected error decoding %q, but got none", input)
			}
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Encode(uint64(i))
	}
}

func BenchmarkDecode(b *testing.B) {
	encoded := Encode(math.MaxUint64)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Decode(encoded)
	}
}
