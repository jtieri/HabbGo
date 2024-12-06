package protocol

import "testing"

func TestVl64Decode(t *testing.T) {
	v := DecodeVl64([]byte("YdA"))
	t.Logf("Decoded VL64 value: %d \n", v)
}
