package jwt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"
)

// Header représente la partie header du JWT
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

// Payload représente la partie payload du JWT
type Payload struct {
	Sub string `json:"sub"`
	Iat int64  `json:"iat"`
}

type JWT struct {
	Hd        Header
	Pld       Payload
}

type Key struct {
	Private *rsa.PrivateKey
	Public  *rsa.PublicKey
}

func (k *Key) GenerateKey() error {
	var err error
	k.Private, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	k.Public = &k.Private.PublicKey
	return nil
}

func (jwt *JWT) GenerateToken() string{
	jwt.Hd = Header{
		Alg: "RS256",
		Typ: "JWT",
	}

	jwt.Pld = Payload{
		Sub: "1234567890",
		Iat: time.Now().Unix(),
	}

	headerJSON := MarshalPart(jwt.Hd)
	payloadJSON := MarshalPart(jwt.Pld)
	message := PartEncode(headerJSON, payloadJSON)

	msgHash := sha256.New()
	_, err := msgHash.Write([]byte(message))
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)
	var key Key
	err = key.GenerateKey()

	if err != nil {
		panic(err)
	}

	signature, err := rsa.SignPSS(rand.Reader, key.Private, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		panic(err)
	}

	Token := message + "." + PartEncode(signature)
	return Token
}

func PartEncode(parts ...[]byte) string {
	var result []string
	for _, p := range parts {
		result = append(result, base64.StdEncoding.EncodeToString(p))

	}
	return strings.Join(result, ".")
}

func MarshalPart(data interface{}) []byte {
	partJSON, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return partJSON
}
