package utils

import (
	"os"
	"encoding/json"
)


type JwtSignKey struct {
	K string `json:"k"`
}

var JWK []byte

func LoadJwtSecretKey() error {
	file, err := os.Open("secret.jwk")
	if err != nil {
		Infof("Err while opening JWK: %s", err.Error()); return err
	}

	var jwtSignkey JwtSignKey
	if err := json.NewDecoder(file).Decode(&jwtSignkey); err != nil {
		Infof("Err while decoding JWK: %s", err.Error()); return err
	}

	JWK = []byte(jwtSignkey.K)
	return nil
}
