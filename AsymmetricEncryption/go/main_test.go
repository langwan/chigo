package main

import (
	"crypto/rand"
	"crypto/rsa"
	"math/big"
	"testing"
)

func TestGenKey(t *testing.T) {
	_, err := GenKey(100, true)
	if err != nil {
		return
	}
}

func TestFromGenKey(t *testing.T) {
	key, err := GenKey(10000, true)
	if err != nil {
		t.Error(err)
		return
	}

	msg := big.NewInt(3)

	cipher := Encrypt(key, msg)
	decryptMsg := Decrypt(key, cipher)
	t.Log("msg", msg)
	t.Log("decryptMsg", decryptMsg)
	if msg.Cmp(decryptMsg) != 0 {
		t.Fail()
	}
}

func TestFromRsaGen(t *testing.T) {
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}
	key := &Key{
		N: rsaKey.N,
		E: big.NewInt(int64(rsaKey.E)),
		D: rsaKey.D,
	}
	msg := "chihuo"
	msgInt := new(big.Int).SetBytes([]byte(msg))
	cipher := Encrypt(key, msgInt)
	decryptMsgInt := Decrypt(key, cipher)
	decryptMsg := string(decryptMsgInt.Bytes())
	t.Log("msg", msg)
	t.Log("decryptMsg", decryptMsg)
	if msg != decryptMsg {
		t.Fail()
	}
}
