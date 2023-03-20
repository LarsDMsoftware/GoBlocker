package util

import (
	randc "crypto/rand"
	"io"
	"math/rand"
	"time"

	"github.com/LarsDMsoftware/GoBlocker/proto"
)

func RandomHash() []byte {
	hash := make([]byte, 32)
	io.ReadFull(randc.Reader, hash)
	return hash
}
func RandomBlock() *proto.Block {
	header := &proto.Header{
		Version:   1,
		Height:    int32(rand.Intn(1000 + 1)),
		PrefHash:  RandomHash(),
		RootHash:  RandomHash(),
		TimeStamp: time.Now().UnixNano(),
	}
	return &proto.Block{
		Header: header,
	}
}
