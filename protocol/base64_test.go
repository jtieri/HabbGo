package protocol

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_Base64Encode asserts that the Base64 encoding function is working as intended.
func Test_Base64Encode(t *testing.T) {
	var tests = []struct {
		name         string
		decodedValue int
		encodedValue []byte
		shouldEqual  bool
	}{
		{"1 should encode to @A", 1, []byte("@A"), true},
		{"2 should not encode to @A", 2, []byte("@A"), false},
		{"257 should encode to DA", 257, []byte("DA"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodedValue := EncodeB64(tt.decodedValue, 2)

			if tt.shouldEqual {
				require.Equal(
					t,
					tt.encodedValue,
					encodedValue,
					"encoded value incorrect, expected(%s) got(%s)", tt.encodedValue, encodedValue,
				)
			} else {
				require.NotEqual(
					t,
					tt.encodedValue,
					encodedValue,
					"encoded value correct, expected(%s) got(%s)", tt.encodedValue, encodedValue,
				)
			}
		})
	}
}

// Test_Base64Decode asserts that the Base64 decoding function is working as intended.
func Test_Base64Decode(t *testing.T) {
	var tests = []struct {
		name         string
		decodedValue int
		encodedValue []byte
		shouldEqual  bool
	}{
		{"A should decode to 1", 1, []byte("A"), true},
		{"@@ should decode to 0", 0, []byte("@@"), true},
		{"@ should decode to 0", 0, []byte("@"), true},
		{"@@@ should decode to 0", 0, []byte("@@@"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decodedValue := DecodeB64(tt.encodedValue)

			if tt.shouldEqual {
				require.Equal(
					t,
					tt.decodedValue,
					decodedValue,
					"decoded value incorrect, expected(%s) got(%s)", tt.decodedValue, decodedValue,
				)
			} else {
				require.NotEqual(
					t,
					tt.decodedValue,
					decodedValue,
					"decoded value correct, expected(%s) got(%s)", tt.decodedValue, decodedValue,
				)
			}
		})
	}
}

// TOPO: delete this later, using to get encoded headers from IDs
func TestEncodingHeaders(t *testing.T) {
	value := EncodeB64(257, 2)
	t.Logf("Encoded header %s \n", value)
}

func TestDecodingHeaders(t *testing.T) {
	value := DecodeB64([]byte("A"))
	t.Logf("Decoded header %d \n", value)
}
