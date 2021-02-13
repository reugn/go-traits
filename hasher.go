package traits

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"hash/fnv"

	"github.com/reugn/go-traits/internal"
)

type hasher interface {
	Md5() [16]byte
	Md5Hex() string
	Sha256() [32]byte
	Sha256Hex() string
	Sha512() [64]byte
	Sha512Hex() string
	HashCode32() uint32
	HashCode64() uint64

	setHasher(i interface{})
}

// Hasher provides unique hash generators to an embedding struct.
type Hasher struct {
	self interface{}
}

var _ hasher = (*Hasher)(nil)

func (h *Hasher) setHasher(i interface{}) {
	h.self = i
}

// Md5 returns the MD5 checksum of an embedding struct.
func (h *Hasher) Md5() [16]byte {
	jsonBytes, err := json.Marshal(h.self)
	internal.Check(err)
	return md5.Sum(jsonBytes)
}

// Md5Hex returns the MD5 checksum hex string representation.
func (h *Hasher) Md5Hex() string {
	md5Bytes := h.Md5()
	return hex.EncodeToString(md5Bytes[:])
}

// Sha256 returns the SHA256 checksum of an embedding struct.
func (h *Hasher) Sha256() [32]byte {
	jsonBytes, err := json.Marshal(h.self)
	internal.Check(err)
	return sha256.Sum256(jsonBytes)
}

// Sha256Hex returns the SHA256 checksum hex string representation.
func (h *Hasher) Sha256Hex() string {
	sha256Bytes := h.Sha256()
	return hex.EncodeToString(sha256Bytes[:])
}

// Sha512 returns the SHA512 checksum of an embedding struct.
func (h *Hasher) Sha512() [64]byte {
	jsonBytes, err := json.Marshal(h.self)
	internal.Check(err)
	return sha512.Sum512(jsonBytes)
}

// Sha512Hex returns the SHA512 checksum hex string representation.
func (h *Hasher) Sha512Hex() string {
	sha512Bytes := h.Sha512()
	return hex.EncodeToString(sha512Bytes[:])
}

// HashCode32 generates a unique uint32 hash of an embedding struct.
func (h *Hasher) HashCode32() uint32 {
	jsonBytes, err := json.Marshal(h.self)
	internal.Check(err)

	h32 := fnv.New32a()
	_, err = h32.Write(jsonBytes)
	internal.Check(err)

	return h32.Sum32()
}

// HashCode64 generates a unique uint64 hash of an embedding struct.
func (h *Hasher) HashCode64() uint64 {
	jsonBytes, err := json.Marshal(h.self)
	internal.Check(err)

	h64 := fnv.New64a()
	_, err = h64.Write(jsonBytes)
	internal.Check(err)

	return h64.Sum64()
}
