package checksum

import (
	"encoding/hex"
	"testing"
)

var key = "qwerty"

func TestWithMixedLowerAndUpper(t *testing.T) {

	var checksum = "e4958Cd97077CB3421112db7354F56A4A39E1F8E9cd2CC5A222800E4DB355192"

	c := NewChecksum(key)
	if !c.Verify(checksum, "5cf621cbb1a9e668525e250e", true, 1001, 0) {
		t.Errorf("Verification of checksum failed")
	}
}

func TestGeneratedMessage(t *testing.T) {

	var expected = "i:5cf621cbb1a9e668525e250e-r:true-s:1001-h:0"

	c := NewChecksum(key)
	r := c.GenerateMessage("5cf621cbb1a9e668525e250e", true, 1001, 0)

	if string(r) != expected {
		t.Errorf("Expected message %s, got %s", expected, string(r))
	}
}

func TestGeneratedChecksum(t *testing.T) {

	expected := "2fc95299861b2416dd2a012ff155d46b85de7e1f4382c6c2719f59beb86ab62e"

	c := NewChecksum(key)
	r := c.ExpectedChecksum([]byte("i:5cf621cbb1a9e668525e250e-r:true-s:1001-h:0"))

	result := hex.EncodeToString(r)

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)

	}
}
