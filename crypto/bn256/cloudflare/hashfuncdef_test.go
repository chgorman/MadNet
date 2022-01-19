package cloudflare

import (
	"bytes"
	"fmt"
	"testing"
)

func TestHashFunc256(t *testing.T) {
	emptyByte := []byte{}
	res := HashFunc256(emptyByte)

	emptyByteHash := make([]byte, 32)

	emptyByteHash[0] = 197
	emptyByteHash[1] = 210
	emptyByteHash[2] = 70
	emptyByteHash[3] = 1
	emptyByteHash[4] = 134
	emptyByteHash[5] = 247
	emptyByteHash[6] = 35
	emptyByteHash[7] = 60

	emptyByteHash[8] = 146
	emptyByteHash[9] = 126
	emptyByteHash[10] = 125
	emptyByteHash[11] = 178
	emptyByteHash[12] = 220
	emptyByteHash[13] = 199
	emptyByteHash[14] = 3
	emptyByteHash[15] = 192

	emptyByteHash[16] = 229
	emptyByteHash[17] = 0
	emptyByteHash[18] = 182
	emptyByteHash[19] = 83
	emptyByteHash[20] = 202
	emptyByteHash[21] = 130
	emptyByteHash[22] = 39
	emptyByteHash[23] = 59

	emptyByteHash[24] = 123
	emptyByteHash[25] = 250
	emptyByteHash[26] = 216
	emptyByteHash[27] = 4
	emptyByteHash[28] = 93
	emptyByteHash[29] = 133
	emptyByteHash[30] = 164
	emptyByteHash[31] = 112

	if !bytes.Equal(res, emptyByteHash) {
		t.Fatal("HashFunc256 changed; invalid hash result for empty byte slice")
	}
}

func TestKMAC1(t *testing.T) {
	trueMsg := []byte{}
	trueMsg = append(trueMsg, []byte("KMAC")...)
	rate := 136
	zeroPad := make([]byte, 132)
	trueMsg = append(trueMsg, zeroPad...)
	if len(trueMsg) != rate {
		t.Fatal("invalid padding")
	}

	emptyKMAC := make([]byte, 32)

	emptyKMAC[0] = 115
	emptyKMAC[1] = 217
	emptyKMAC[2] = 78
	emptyKMAC[3] = 34
	emptyKMAC[4] = 83
	emptyKMAC[5] = 52
	emptyKMAC[6] = 47
	emptyKMAC[7] = 90

	emptyKMAC[8] = 27
	emptyKMAC[9] = 31
	emptyKMAC[10] = 121
	emptyKMAC[11] = 200
	emptyKMAC[12] = 39
	emptyKMAC[13] = 121
	emptyKMAC[14] = 49
	emptyKMAC[15] = 130

	emptyKMAC[16] = 69
	emptyKMAC[17] = 231
	emptyKMAC[18] = 227
	emptyKMAC[19] = 91
	emptyKMAC[20] = 81
	emptyKMAC[21] = 19
	emptyKMAC[22] = 83
	emptyKMAC[23] = 44

	emptyKMAC[24] = 228
	emptyKMAC[25] = 11
	emptyKMAC[26] = 53
	emptyKMAC[27] = 242
	emptyKMAC[28] = 98
	emptyKMAC[29] = 41
	emptyKMAC[30] = 95
	emptyKMAC[31] = 59

	rethash := kmac(nil, nil, nil)
	if !bytes.Equal(rethash, emptyKMAC) {
		t.Fatal("invalid KMAC")
	}
}

func TestKMAC2(t *testing.T) {
	trueMsg := []byte{}
	trueMsg = append(trueMsg, []byte("KMAC")...)
	emptyKey := make([]byte, 32)
	trueMsg = append(trueMsg, emptyKey...)
	emptyCustom := make([]byte, 32)
	trueMsg = append(trueMsg, emptyCustom...)
	rate := 136
	zeroPad := make([]byte, rate-len(trueMsg))
	trueMsg = append(trueMsg, zeroPad...)
	if len(trueMsg) != rate {
		t.Fatal("invalid padding")
	}

	trueMsgHash := HashFunc256(trueMsg)

	emptyKMAC := make([]byte, 32)

	emptyKMAC[0] = 115
	emptyKMAC[1] = 217
	emptyKMAC[2] = 78
	emptyKMAC[3] = 34
	emptyKMAC[4] = 83
	emptyKMAC[5] = 52
	emptyKMAC[6] = 47
	emptyKMAC[7] = 90

	emptyKMAC[8] = 27
	emptyKMAC[9] = 31
	emptyKMAC[10] = 121
	emptyKMAC[11] = 200
	emptyKMAC[12] = 39
	emptyKMAC[13] = 121
	emptyKMAC[14] = 49
	emptyKMAC[15] = 130

	emptyKMAC[16] = 69
	emptyKMAC[17] = 231
	emptyKMAC[18] = 227
	emptyKMAC[19] = 91
	emptyKMAC[20] = 81
	emptyKMAC[21] = 19
	emptyKMAC[22] = 83
	emptyKMAC[23] = 44

	emptyKMAC[24] = 228
	emptyKMAC[25] = 11
	emptyKMAC[26] = 53
	emptyKMAC[27] = 242
	emptyKMAC[28] = 98
	emptyKMAC[29] = 41
	emptyKMAC[30] = 95
	emptyKMAC[31] = 59

	if !bytes.Equal(trueMsgHash, emptyKMAC) {
		t.Fatal("invalid KMAC value")
	}

	rethash := kmac(nil, nil, nil)
	if !bytes.Equal(rethash, emptyKMAC) {
		t.Fatal("invalid KMAC")
	}
}

func TestKMAC3(t *testing.T) {
	trueMsg := []byte{}
	trueMsg = append(trueMsg, []byte("KMAC")...)
	key := make([]byte, 32)
	for k := 0; k < len(key); k++ {
		key[k] = 239
	}
	trueMsg = append(trueMsg, key...)
	custom := make([]byte, 32)
	for k := 0; k < len(custom); k++ {
		custom[k] = 241
	}
	trueMsg = append(trueMsg, custom...)
	rate := 136
	zeroPad := make([]byte, rate-len(trueMsg))
	trueMsg = append(trueMsg, zeroPad...)
	if len(trueMsg) != rate {
		t.Fatal("invalid padding")
	}
	trueMsgHash := HashFunc256(trueMsg)

	trueKMAC := make([]byte, 32)

	trueKMAC[0] = 70
	trueKMAC[1] = 165
	trueKMAC[2] = 102
	trueKMAC[3] = 206
	trueKMAC[4] = 229
	trueKMAC[5] = 255
	trueKMAC[6] = 204
	trueKMAC[7] = 198

	trueKMAC[8] = 168
	trueKMAC[9] = 249
	trueKMAC[10] = 10
	trueKMAC[11] = 128
	trueKMAC[12] = 63
	trueKMAC[13] = 22
	trueKMAC[14] = 120
	trueKMAC[15] = 222

	trueKMAC[16] = 228
	trueKMAC[17] = 15
	trueKMAC[18] = 244
	trueKMAC[19] = 69
	trueKMAC[20] = 59
	trueKMAC[21] = 47
	trueKMAC[22] = 212
	trueKMAC[23] = 232

	trueKMAC[24] = 86
	trueKMAC[25] = 124
	trueKMAC[26] = 76
	trueKMAC[27] = 71
	trueKMAC[28] = 72
	trueKMAC[29] = 152
	trueKMAC[30] = 213
	trueKMAC[31] = 96

	if !bytes.Equal(trueMsgHash, trueKMAC) {
		t.Fatal("invalid KMAC value")
	}

	rethash := kmac(key, nil, custom)
	if !bytes.Equal(rethash, trueKMAC) {
		fmt.Println(rethash)
		t.Fatal("invalid KMAC")
	}
}
