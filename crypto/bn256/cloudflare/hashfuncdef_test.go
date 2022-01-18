package cloudflare

import (
	"bytes"
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

func TestKMAC(t *testing.T) {
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
