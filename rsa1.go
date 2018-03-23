package main

import (
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "fmt"
    "os"
    "log"
)

func main(){
  mariaPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
  if err != nil {
      fmt.Println(err.Error)
      os.Exit(1)
  }
  mariaPublicKey := &mariaPrivateKey.PublicKey
  raulPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
  if err != nil {
      fmt.Println(err.Error)
      os.Exit(1)
  }
  raulPublicKey := &raulPrivateKey.PublicKey

  file, err := os.Create("result.txt")
    if err != nil {
        log.Fatal("Cannot create file", err)
    }
    defer file.Close()

/*----------------------------------------------------------*/
    fmt.Fprintf(file, "maria private key : ")
    fmt.Fprintf(file, "\n")
    fmt.Fprintf(file, "\n")
    fmt.Fprintf(file, fmt.Sprint(mariaPrivateKey))
    fmt.Fprintf(file, "\n")
    fmt.Fprintf(file, "\n")
    fmt.Fprintf(file, "maria public key : ")
    fmt.Fprintf(file, "\n")
    fmt.Fprintf(file, "\n")
    fmt.Fprintf(file, fmt.Sprint(mariaPublicKey))

    fmt.Fprintf(file, "\n")
    fmt.Fprintf(file, "\n")
    fmt.Fprintf(file, "\n")

    fmt.Fprintf(file, "raul private key : ")
    fmt.Fprintf(file, "\n")
    fmt.Fprintf(file, "\n")
    fmt.Fprintf(file, fmt.Sprint(raulPrivateKey))
    fmt.Fprintf(file, "\n")
    fmt.Fprintf(file, "\n")
    fmt.Fprintf(file, "raul public key : ")
    fmt.Fprintf(file, "\n")
    fmt.Fprintf(file, "\n")
    fmt.Fprintf(file, fmt.Sprint(raulPublicKey))

  message := "the code must be like a piece of music"

  ciphertext := getCypherTextWithPubKey(message, raulPublicKey)

  fmt.Printf("OAEP encrypted [%s] to \n[%x]\n", message, ciphertext)

  signature := getSignatureWithPrivKey(message, mariaPrivateKey)

  fmt.Printf("PSS Signature : %x\n", signature)

  plainText := getPlainTextWithPrivateKey(ciphertext, raulPrivateKey)

  fmt.Printf("OAEP decrypted [%x] to \n[%s]\n", ciphertext, plainText)

  verifySignatureWithPublicKey(fmt.Sprintf("%s", plainText), signature, mariaPublicKey)
}

func verifySignatureWithPublicKey(message string, signature []byte, key *rsa.PublicKey){
  newhash := crypto.SHA256
  var opts rsa.PSSOptions
  opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
  pssh := newhash.New()
  pssh.Write([]byte(message))
  hashed := pssh.Sum(nil)
  err := rsa.VerifyPSS(
    key,
    newhash,
    hashed,
    signature,
    &opts)
  if err != nil {
      fmt.Println("Who are U? Verify Signature failed")
      os.Exit(1)
  } else {
      fmt.Println("Verify Signature successful")
  }
}

func getPlainTextWithPrivateKey(ciphertext []byte, key *rsa.PrivateKey) []byte{
  hash := sha256.New()
  label := []byte("")
  plainText, err := rsa.DecryptOAEP(
    hash,
    rand.Reader,
    key,
    ciphertext,
    label)
if err != nil {
    fmt.Println(err)
    os.Exit(1)
}
return plainText
}

func getSignatureWithPrivKey(message string, key *rsa.PrivateKey) []byte{
  var opts rsa.PSSOptions
  opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
  PSSmessage := message
  newhash := crypto.SHA256
  pssh := newhash.New()
  pssh.Write([]byte(PSSmessage))
  hashed := pssh.Sum(nil)
  signature, err := rsa.SignPSS(
      rand.Reader,
      key,
      newhash,
      hashed,
      &opts)
  if err != nil {
      fmt.Println(err)
      os.Exit(1)
  }

  return signature
}

func getCypherTextWithPubKey(msg string, key *rsa.PublicKey) []byte{
  message := []byte(msg)
  label := []byte("")
  hash := sha256.New()
  ciphertext, err := rsa.EncryptOAEP(
      hash,
      rand.Reader,
      key,
      message,
      label)

  if err != nil {
      fmt.Println(err)
      os.Exit(1)
  }

  return ciphertext
}
