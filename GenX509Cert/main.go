package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {
	random := rand.Reader
	now := time.Now()
	then := now.Add(60 * 60 * 24 * 365 * 1000 * 1000 * 1000) //一年
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "jan.newmarch.name", Organization: []string{"Jan Newmarch"}},
		// NotBefore:time.Unix(now,0).UTC(),
		// NotAfter:time.Unix(now+60*60*24*365,0).UTC()
		NotBefore:             now,
		NotAfter:              then,
		SubjectKeyId:          []byte{1, 2, 3, 4},
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:     true,
		DNSNames: []string{"jan.newmarch.name", "localhost"},
	}
	var key rsa.PrivateKey
	loadKey("../GenRSAKeys/private.key", &key)
	derBytes, err := x509.CreateCertificate(random, &template, &template, &key.PublicKey, &key)
	checkError(err)
	certCerFile, err := os.Create("jan.newmarch.name.cer")
	checkError(err)
	//defer certCerFile.Close()
	certCerFile.Write(derBytes)
	certPEMFile, err := os.Create("jan.newmarch.name.pem")
	checkError(err)
	pem.Encode(certPEMFile, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	keyPEMFile, err := os.Create("private.pem")
	checkError(err)
	pem.Encode(keyPEMFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(&key)})
}

func loadKey(file string, key *rsa.PrivateKey) {
	in, err := os.Open(file)
	checkError(err)
	decoder := gob.NewDecoder(in)
	err = decoder.Decode(key)
	checkError(err)
}

func checkError(err error) {
	if nil != err {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
