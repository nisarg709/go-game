package service

import (
	"time"
	"fmt"
	"crypto/md5"
)

type Buffer struct {
	data   []byte
	etag   string
	expire time.Time
}

func (b *Buffer) IsBufferedDataFresh() bool {
	if b.expire.IsZero() {
		return false
	}

	return b.expire.After(time.Now())
}

func (b *Buffer) isEtagLatest(etag string) bool {
	return etag == b.etag
}

func (b *Buffer) ExtendExpiration(etag string) {
	if etag == b.etag {
		b.expire = time.Now().Add(10 * time.Second)
	} else {
		b.expire = time.Now().Add(-1 * time.Hour)
	}
}

func (b *Buffer) GetBufferedData() []byte {
	return b.data
}

func (b *Buffer) SetBufferedData(data []byte) {
	b.data = data
	b.etag = fmt.Sprintf("%x", md5.Sum(data))
	b.expire = time.Now().Add(10 * time.Second)
}
