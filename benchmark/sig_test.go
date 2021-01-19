package main

import (
	cRand "crypto/rand"
	"io"
	"testing"

	"github.com/open-quantum-safe/liboqs-go/oqs"
)

var sigName = "DILITHIUM_3"

func BenchmarkKeyGen(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		var signer oqs.Signature
		b.StartTimer()
		_ = signer.Init(sigName, nil)
		signer.GenerateKeyPair()
	}
}

func BenchmarkSign(b *testing.B) {
	var msg [59]byte
	rand := cRand.Reader
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		io.ReadFull(rand, msg[:])
		var signer oqs.Signature
		// ignore potential errors everywhere
		_ = signer.Init(sigName, nil)
		signer.GenerateKeyPair()
		b.StartTimer()
		signer.Sign(msg[:])
	}
}

func BenchmarkVerify(b *testing.B) {
	var msg [59]byte
	rand := cRand.Reader
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		io.ReadFull(rand, msg[:])
		var signer, verifier oqs.Signature
		// ignore potential errors everywhere
		_ = signer.Init(sigName, nil)
		_ = verifier.Init(sigName, nil)
		pubKey, _ := signer.GenerateKeyPair()
		signature, _ := signer.Sign(msg[:])
		b.StartTimer()
		verifier.Verify(msg[:], signature, pubKey)
	}
}
