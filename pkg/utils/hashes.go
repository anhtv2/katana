package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/mfonda/simhash"
)

func Md5(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func Sha1(data []byte) string {
	hash := sha1.Sum(data)
	return hex.EncodeToString(hash[:])
}

func Sha256(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func Sha512(data []byte) string {
	hash := sha512.Sum512(data)
	return hex.EncodeToString(hash[:])
}

func Simhash(data []byte) string {
	hash := simhash.Simhash(simhash.NewWordFeatureSet(data))
	return fmt.Sprintf("%d", hash)
}

func CalculateHashes(body []byte, headers []byte) map[string]interface{} {
	hashes := make(map[string]interface{})
	if len(body) > 0 {
		hashes["body_md5"] = Md5(body)
		hashes["body_sha1"] = Sha1(body)
		hashes["body_sha256"] = Sha256(body)
		hashes["body_sha512"] = Sha512(body)
		hashes["body_simhash"] = Simhash(body)
	}
	if len(headers) > 0 {
		hashes["header_md5"] = Md5(headers)
		hashes["header_sha1"] = Sha1(headers)
		hashes["header_sha256"] = Sha256(headers)
		hashes["header_sha512"] = Sha512(headers)
		hashes["header_simhash"] = Simhash(headers)
	}
	return hashes
}
