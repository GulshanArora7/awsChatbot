package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//StringInSlice checks if a slice contains a specific string
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Md5Hash encodes a string using the MD5 algorithm
func Md5Hash(hashData string) string {
	hash := md5.Sum([]byte(hashData))
	hashStr := hex.EncodeToString(hash[:])
	return hashStr
}
