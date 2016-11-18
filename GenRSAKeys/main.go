package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"os"
	"runtime/debug"
)

func main() {
	reader := rand.Reader
	bitSize := 512
	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)
	fmt.Println("Private key primes", key.Primes[0].String(), key.Primes[1].String())
	fmt.Println("Private key exponent", key.D.String())

	publicKey := key.PublicKey
	fmt.Println("Public key modulus", publicKey.N.String())
	fmt.Println("Public key exponent", publicKey.E)
	saveGobKey("private.key", key)
	saveGobKey("public.key", publicKey)

	savePEMKey("private.pem", key)
}

func checkError(err error) {
	if nil != err {
		fmt.Println("Fatal error ", err.Error())
		debug.PrintStack()
		os.Exit(1)
	}
}

func saveGobKey(file string, key interface{}) {
	out, err := os.Create(file)
	checkError(err)
	defer out.Close()
	encoder := gob.NewEncoder(out)
	err = encoder.Encode(key)
	checkError(err)
}

func savePEMKey(file string, key *rsa.PrivateKey) {
	out, err := os.Create(file)
	checkError(err)
	defer out.Close()
	privateKey := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}
	pem.Encode(out, privateKey)
}
