package main

import (
  "crypto/rsa"
  "io/ioutil"
  "time"
  "log"

  jwt "github.com/dgrijalva/jwt-go"
)

const (
  privKeyPath = "../keys/app.rsa"     // openssl genrsa -out app.rsa keysize
  pubKeyPath  = "../keys/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
  signKey   *rsa.PrivateKey
  verifyKey *rsa.PublicKey
)

func CreateToken(payload string)(string, error) {
  signBytes, err := ioutil.ReadFile(privKeyPath)
  fatal(err)

  signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
  fatal(err)

  t := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
    "payload": payload,
    "exp": time.Now().Add(time.Hour * 24).Unix(),
    })
  tokenStr, err := t.SignedString(signKey)
  fatal(err)
  return tokenStr, err
}

func VerifyToken(tokenStr string)(bool, string) {
  verifyBytes, err := ioutil.ReadFile(pubKeyPath)
  fatal(err)

  verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
  fatal(err)

  keyLookupFn := func(token *jwt.Token) (interface{}, error) {
    return verifyKey, nil
  }

  token, err := jwt.Parse(tokenStr, keyLookupFn)
  if err != nil {
    return false, err.Error()
  } else if !token.Valid{
    return false, "Token invalid!"
  } else {
    return true, "OK"
  }
}

func fatal(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
