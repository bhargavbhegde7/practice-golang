package main

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "errors"
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

func ParseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
    block, _ := pem.Decode([]byte(privPEM))
    if block == nil {
            return nil, errors.New("failed to parse PEM block containing the key")
    }

    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
            return nil, err
    }

    return priv, nil
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

func ParseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
    block, _ := pem.Decode([]byte(pubPEM))
    if block == nil {
            return nil, errors.New("failed to parse PEM block containing the key")
    }

    pub, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
            return nil, err
    }

    switch pub := pub.(type) {
    case *rsa.PublicKey:
            return pub, nil
    default:
            break // fall through
    }
    return nil, errors.New("Key type is not RSA")
}

func main() {

    // Create the keys
    priv, pub := GenerateRsaKeyPair()

    // Export the keys to pem string
    priv_pem_str := ExportRsaPrivateKeyAsPemStr(priv)
    pub_pem_str, _ := ExportRsaPublicKeyAsPemStr(pub)

    writeToFile("priv_key.txt", priv_pem_str)
    writeToFile("pub_key.txt", pub_pem_str)

    // Import the keys from pem string
    priv_parsed, _ := ParseRsaPrivateKeyFromPemStr(readFromFile("priv_key.txt"))
    pub_parsed, _ := ParseRsaPublicKeyFromPemStr(readFromFile("pub_key.txt"))

    // Export the newly imported keys
    priv_parsed_pem := ExportRsaPrivateKeyAsPemStr(priv_parsed)
    pub_parsed_pem, _ := ExportRsaPublicKeyAsPemStr(pub_parsed)

    fmt.Println(priv_parsed_pem)
    fmt.Println(pub_parsed_pem)

}

func writeToFile(fileName string, text string){
  file, err := os.Create(fileName)
  if err != nil {
      log.Fatal("Cannot create file", err)
  }
  defer file.Close()

  fmt.Fprintf(file, text)
}

func readFromFile(fileName string) string{
  file, err := os.Open(fileName)
  if err != nil {
    fmt.Println(err)
    //return
  }
  defer file.Close()

  fileinfo, err := file.Stat()
  if err != nil {
    fmt.Println(err)
    //return
  }

  filesize := fileinfo.Size()
  buffer := make([]byte, filesize)

  _, err = file.Read(buffer)
  if err != nil {
    fmt.Println(err)
    //return
  }

  //fmt.Println("bytes read: ", bytesread)
  //fmt.Println("bytestream to string: ", string(buffer))

  result := string(buffer)

  return result
}
