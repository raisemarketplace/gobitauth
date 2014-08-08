package bitauth

import (
	"testing"
)

func TestGetSINFromPublicKey(t *testing.T) {
	var keys = map[string]SIN{
		"02F840A04114081690223B7069071A70D6DABB891763B638CC20C7EC3BD58E6C86": SIN("TfG4ScDgysrSpodWD4Re5UtXmcLbY5CiUHA"),
		"023f5d74e874b2f7c784729fc93b7d38a3c28129d27321b2e4f7cde09d7609adff": SIN("TeyP7pTjXKMoZuHPvDpoUfopfjrHCosFB7W"),
	}

	for key, sin := range keys {
		gen, _ := GetSINFromPublicKey(key)
		if gen != sin {
			t.Errorf("Invalid sin. Expected %s, generated %s from key %s", sin, gen, key)
			return
		}
	}
}

func TestGetPublicKeyFromPrivateKey(t *testing.T) {
	// map[private]public{}
	// Maps private keys (as the map key) to public keys (value)
	var keys = map[string]string{
		"1217a587403374c5b21897bda9eaaa88f7315c931e4cb1d40f24fcf8f4d34419": "023f5d74e874b2f7c784729fc93b7d38a3c28129d27321b2e4f7cde09d7609adff",
	}

	for private, public := range keys {
		if generated, _ := GetPublicKeyFromPrivateKey(private); generated != public {
			t.Errorf("Invalid public key. Expected %s, generated %s from private key %s", public, generated, private)
			return
		}
	}
}

func TestSINGeneration(t *testing.T) {
	sininfo, _ := GenerateSIN()

	if expected, _ := GetSINFromPublicKey(sininfo.PublicKey); expected != sininfo.SIN {
		t.Errorf("Generated SIN/Public key do not match")
		return
	}
}