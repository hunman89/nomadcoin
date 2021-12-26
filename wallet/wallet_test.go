package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"testing"
)

const (
	testKey     string = "307702010104205c09540e8c81cab147044bbd37187623e4c74e20a300f8ee8d1831a725f72d6aa00a06082a8648ce3d030107a14403420004d58f94e39917a976422a0f362a582aa21ca22cdcad400dea0f0a21f340c42b3f6e36680722542bedf7a03c2b4a05d87e9dec2abd7bc4d77d00eec45102d0252b"
	testPayload string = "07899002d762e2f1663b2f558fd23d2277dc3a59f02f0e9f0b9e8deeaeba6e83"
	testSig     string = "85e92e60538044da06c9087ebdb63f831cf84d1424da96883831bbef470b38552011da8196cb98314d78152531c51f8e52d4b893b16df6b1aebbf2cdfb96f122"
)

func makeTestWallet() *wallet {
	w := &wallet{}
	b, _ := hex.DecodeString(testKey)
	key, _ := x509.ParseECPrivateKey(b)
	w.privateKey = key
	w.Address = aFromK(key)
	return w
}

func TestSign(t *testing.T) {
	s := Sign(testPayload, makeTestWallet())
	_, err := hex.DecodeString(s)
	if err != nil {
		t.Errorf("Sign() should return a hex encoded string, got %s", s)
	}
}

func TestVerify(t *testing.T) {
	type test struct {
		input string
		ok    bool
	}
	tests := []test{
		{testPayload, true},
		{"17899002d762e2f1663b2f558fd23d2277dc3a59f02f0e9f0b9e8deeaeba6e83", false},
	}
	for _, tc := range tests {
		w := makeTestWallet()
		ok := Verify(testSig, tc.input, w.Address)
		if ok != tc.ok {
			t.Error("Verify() could not verify testSignature and testPayload")
		}
	}
}

func TestRestoreBigInts(t *testing.T) {
	_, _, err := restoreBigInts("xx")
	if err == nil {
		t.Error("RestoreBigInts should return error when payload is not hex")
	}
}
