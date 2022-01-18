package cloudflare

import (
	eth "github.com/ethereum/go-ethereum/crypto"
)

// HashFunc256 is our universal hash function with 32-byte digest.
// Changing this function will change *all* locations.
func HashFunc256(msg ...[]byte) []byte {
	return eth.Keccak256(msg...)
}

func kmac(key []byte, msg []byte, customization []byte) []byte {
	rate := 136
	newMsg := []byte{}
	newMsg = append(newMsg, []byte("KMAC")...)
	newMsg = append(newMsg, key...)
	newMsg = append(newMsg, customization...)
	numZeros := rate - (len(newMsg) % rate)
	padZeros := make([]byte, numZeros)
	newMsg = append(newMsg, padZeros...)
	newMsg = append(newMsg, msg...)
	return HashFunc256(newMsg)
}
