// Package oqstests provides unit testing for the oqs Go package.
package oqstests

import (
	"testing"

	"github.com/open-quantum-safe/liboqs-go/oqs"
)

var kem_name = "KYBER_768"

func BenchmarkKeyGenKEM(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		var client oqs.KeyEncapsulation
		b.StartTimer()
		_ = client.Init(kemName, nil)
		clientPublicKey, _ := client.GenerateKeyPair()
	}
}

func BenchmarkSign(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		var client oqs.KeyEncapsulation
		_ = client.Init(kemName, nil)
		clientPublicKey, _ := client.GenerateKeyPair()
		b.StartTimer()
		ciphertext, sharedSecretServer, _ := server.EncapSecret(clientPublicKey)
	}
}

func BenchmarkVerify(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		var client, server oqs.KeyEncapsulation
		_ = client.Init(kemName, nil)
		_ = server.Init(kemName, nil)
		clientPublicKey, _ := client.GenerateKeyPair()
		ciphertext, sharedSecretServer, _ := server.EncapSecret(clientPublicKey)
		b.StartTimer()
		sharedSecretClient, _ := client.DecapSecret(ciphertext)
	}
}
