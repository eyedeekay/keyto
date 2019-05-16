package i2pdest

import (
	"encoding/base32"
	"encoding/base64"
	"crypto/sha256"
    "encoding/binary"
    "strings"
)

var (
	i2pB64enc *base64.Encoding = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-~")
	i2pB32enc *base32.Encoding = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567")
)

// Base32 returns the base32 address corresponding to the destination key
func Base32(c string) (string, error) {
	hash := sha256.New()
    b64, err := Base64(c)
	if err != nil {
        return "", err
    }
    sb64, err := i2pB64enc.DecodeString(b64)
	if err != nil {
        return "", err
    }
    hash.Write([]byte(sb64))
	return strings.ToLower(strings.Replace(i2pB32enc.EncodeToString(hash.Sum(nil)), "=", "", -1)), nil
}

// Base64 returns the base64 address corresponding to the destination key
func Base64(c string) (string, error) {
	s, err := i2pB64enc.DecodeString(c)
	if err != nil {
        return "", err
    }
    alen := binary.BigEndian.Uint16(s[385:387])
	return i2pB64enc.EncodeToString(s[:387+alen]), nil
}
