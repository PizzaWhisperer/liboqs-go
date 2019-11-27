// signature Go example
package main

import (
	"fmt"
	"github.com/open-quantum-safe/liboqs-go/oqs"
)

func main() {
	fmt.Println("Supported signatures:")
	fmt.Println(oqs.SupportedSigs())

	fmt.Println("\nEnabled signatures:")
	fmt.Println(oqs.EnabledSigs())

	sigName := "DEFAULT"
	signer := oqs.Signature{}
	defer signer.Clean() // clean up even in case of panic

	signer.Init(sigName, nil)
	fmt.Println("\nSignature details:")
	fmt.Println(signer.Details())

	msg := []byte("This is the message to sign")
	pubKey := signer.GenerateKeyPair()
	fmt.Printf("\nSigner public key:\n% X ... % X\n", pubKey[0:8],
		pubKey[len(pubKey)-8:])

	signature := signer.Sign(msg)
	fmt.Printf("\nSignature:\n% X ... % X\n", signature[0:8],
		signature[len(signature)-8:])

	verifier := oqs.Signature{}
	defer verifier.Clean() // clean up even in case of panic

	verifier.Init(sigName, nil)
	isValid := verifier.Verify(msg, signature, pubKey)

	fmt.Println("\nValid signature? ", isValid)
}
