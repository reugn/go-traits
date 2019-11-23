package traits

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"hash/fnv"
)

type hash interface {
	Md5() [16]byte
	Md5Hex() string
	Sha256() [32]byte
	Sha256Hex() string
	HashCode() uint32
	setHasher(i interface{})
}

// A Hash trait provides unique hash generators
type Hash struct {
	self interface{}
}

func (h *Hash) setHasher(i interface{}) {
	h.self = i
}

// Md5 byte array for self struct
func (h *Hash) Md5() [16]byte {
	jsonBytes, _ := json.Marshal(h.self)
	return md5.Sum(jsonBytes)
}

// Md5Hex string for self struct
func (h *Hash) Md5Hex() string {
	md5Bytes := h.Md5()
	return hex.EncodeToString(md5Bytes[:])
}

// Sha256 byte array for self struct
func (h *Hash) Sha256() [32]byte {
	jsonBytes, _ := json.Marshal(h.self)
	return sha256.Sum256(jsonBytes)
}

// Sha256Hex string for self struct
func (h *Hash) Sha256Hex() string {
	sha256Bytes := h.Sha256()
	return hex.EncodeToString(sha256Bytes[:])
}

// HashCode uint32 for self struct
func (h *Hash) HashCode() uint32 {
	jsonBytes, _ := json.Marshal(h.self)
	h32 := fnv.New32a()
	h32.Write(jsonBytes)
	return h32.Sum32()
}
