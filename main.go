package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

func CreatePrivateKey(curve elliptic.Curve) (*big.Int, error) {
	key, err := rand.Int(rand.Reader, curve.Params().N)
	if err != nil {
		return nil, err
	}
	key.Add(key, big.NewInt(1))
	return key, nil
}

func CalculateSecret(curve elliptic.Curve, privateKey *big.Int, pkX, pkY *big.Int) *big.Int {
	sharedX, _ := curve.ScalarMult(pkX, pkY, privateKey.Bytes())
	sharedX.Mod(sharedX, curve.Params().P)
	return sharedX
}

func CreateP256Curve() elliptic.Curve {
	return elliptic.P256()
}

func main() {
	Curve := CreateP256Curve()

	AlicePrivateKey, err := CreatePrivateKey(Curve)
	if err != nil {
		log.Println("Failed To Create Alice's private key:", err)
		return
	}

	BobPrivateKey, err := CreatePrivateKey(Curve)
	if err != nil {
		log.Println("Failed To Create Bob's private key:", err)
		return
	}

	AlicePublicKeyX, alicePublicKeyY := Curve.ScalarBaseMult(AlicePrivateKey.Bytes())
	BobPublicKeyX, bobPublicKeyY := Curve.ScalarBaseMult(BobPrivateKey.Bytes())

	Secret1 := CalculateSecret(Curve, AlicePrivateKey, BobPublicKeyX, bobPublicKeyY)
	Secret2 := CalculateSecret(Curve, BobPrivateKey, AlicePublicKeyX, alicePublicKeyY)

	fmt.Println("Alice:", AlicePublicKeyX)
	fmt.Println("Bob:", BobPublicKeyX)
	if Secret1.String() == Secret2.String() {
		fmt.Println("\nSecret Keys Match!")
		fmt.Printf("Shared Secret 1: %x\nShared Secret 2: %x\n", Secret1.Bytes(), Secret1.Bytes())
		fmt.Printf("Shared secret: %x\n", Secret1.Bytes())
	} else {
		log.Println("Shared secrets did not match!")
	}
}
