package main

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "fmt"
    "log"
    "os"
)

func GenerateRsaKeyPair() (*rsa.PrivateKey, *rsa.PublicKey) {
    privkey, _ := rsa.GenerateKey(rand.Reader, 4096)
    return privkey, &privkey.PublicKey
}

func ExportRsaPrivateKeyAsPemStr(privkey *rsa.PrivateKey) string {
    privkey_bytes := x509.MarshalPKCS1PrivateKey(privkey)
    privkey_pem := pem.EncodeToMemory(
            &pem.Block{
                    Type:  "RSA PRIVATE KEY",
                    Bytes: privkey_bytes,
            },
    )
    return string(privkey_pem)
}

func ExportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) (string, error) {
    pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
    if err != nil {
            return "", err
    }
    pubkey_pem := pem.EncodeToMemory(
            &pem.Block{
                    Type:  "RSA PUBLIC KEY",
                    Bytes: pubkey_bytes,
            },
    )

    return string(pubkey_pem), nil
}

func main() {

    // Create the keys
    priv, pub := GenerateRsaKeyPair()

    // Export the keys to pem string
    priv_pem_str := ExportRsaPrivateKeyAsPemStr(priv)
    pub_pem_str, _ := ExportRsaPublicKeyAsPemStr(pub)

    writeToFile("priv_key", priv_pem_str)
    writeToFile("pub_key", pub_pem_str)
}

func writeToFile(fileName string, text string){
  file, err := os.Create(fileName)
  if err != nil {
      log.Fatal("Cannot create file", err)
  }
  defer file.Close()

  fmt.Fprintf(file, text)
}
