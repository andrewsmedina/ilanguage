package encoder

import "testing"

func TestEncode(t *testing.T) {
	sample := "o sapo nao lava o pe"
	translated := Encode(sample)
	expected := "i sipi nii livi i pi"

	if translated != expected {
		t.Errorf("Expected %s, got %s", expected, translated)
	}
}
