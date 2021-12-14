package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/hunman89/nomadcoin/utils"
)

const (
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
	privateKey    string = "3077020101042095ff23f66ac771cfc2e335fa06d6c355de036eeb95e33db102b5b2931b523f08a00a06082a8648ce3d030107a14403420004a71c398dd854f9b7ebf0bb073bcf406fb6944d324618e2c61b50ca3116708fa827271d8646754e09a8f5d922e1e0790413d37931e83d5f22b3547205ef00df4f"
	signature     string = "4fa31be12c6f87c658f004a58c8e7247b3a7c9eac8d1d6ab4a2f6cf5cbdb397e1605218984cce59c471e43a90cdecc771e1d782177499c852c80b9be4ba83a"
)

func Start() {

	privBytes, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	_, err = x509.ParseECPrivateKey(privBytes)
	utils.HandleErr(err)

	sigBytes, err := hex.DecodeString(signature)

	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}
	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)
	fmt.Println(bigR, bigS)
}
