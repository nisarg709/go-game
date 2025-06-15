package checksum

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"strconv"
	"strings"
)

func NewChecksum(key string) *Checksum {
	hasher := hmac.New(sha256.New, []byte(key))
	return &Checksum{hasher}
}

type Checksum struct {
	hasher hash.Hash
}

func (self *Checksum) Verify(checksum string, id string, success bool, score int, helps int) bool {

	message := self.GenerateMessage(id, success, score, helps)
	expected := self.ExpectedChecksum(message)

	return strings.ToLower(checksum) == strings.ToLower(hex.EncodeToString(expected))
}

func (self *Checksum) GenerateMessage(id string, success bool, score int, helps int) []byte {
	message := "i:" + id + "-r:" + strconv.FormatBool(success) + "-s:" + strconv.Itoa(score) + "-h:" + strconv.Itoa(helps);
	return []byte(message);
}

func (self *Checksum) ExpectedChecksum(message []byte) []byte {
	self.hasher.Write(message)

	return self.hasher.Sum(nil)
}
