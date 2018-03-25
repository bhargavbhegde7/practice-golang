package main

import(
  "fmt"
  "os"
  "crypto/rand"
  "crypto/rsa"
  "crypto/sha256"
  "crypto/x509"
  "encoding/pem"
  "errors"
)

//send command line arguments:
//1. filename of the public key text file
//2. string to encrypt
func main(){
  pub_fileName := os.Args[1]
  priv_fileName := os.Args[2]

  message := os.Args[3]

/* ------------------------------------------------------------------------ */
  pubKey_str := readKeyFromFile(pub_fileName)
  pub_key, err := convertStringPubKeyToRsaKey(pubKey_str)

  if err != nil{
    fmt.Println(err)
  }
/* ------------------------------------------------------------------------- */

/* ------------------------------------------------------------------------ */
  privKey_str := readKeyFromFile(priv_fileName)
  priv_key, err := convertStringPrivKeyToRsaKey(privKey_str)

  if err != nil{
    fmt.Println(err)
  }
/* ------------------------------------------------------------------------- */

  encrypted := getEncrypted(message, pub_key)

  //fmt.Printf("\n\n encrypted : \n%s\n", encrypted)
  fmt.Println(fmt.Sprintf("%s", encrypted))

  fmt.Printf("\n\n decrypted : \n%s\n", getDecrypted(encrypted, priv_key))

}

func convertStringPrivKeyToRsaKey(privPEM string) (*rsa.PrivateKey, error) {
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

//decrypter
func getDecrypted(ciphertext []byte, key *rsa.PrivateKey) []byte{
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

//encrypter
func getEncrypted(msg string, key *rsa.PublicKey) []byte{
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

func readKeyFromFile(fileName string) string{
  file, err := os.Open(fileName)
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  fileinfo, err := file.Stat()
  if err != nil {
    fmt.Println(err)
  }

  filesize := fileinfo.Size()
  buffer := make([]byte, filesize)

  _, err = file.Read(buffer)
  if err != nil {
    fmt.Println(err)
  }

  result := string(buffer)

  return result
}

func convertStringPubKeyToRsaKey(pubPEM string) (*rsa.PublicKey, error) {
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
