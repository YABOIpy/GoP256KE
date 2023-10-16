# GoP256KE
Implementation of the P-256 curve for key exchange in Go


![image](https://github.com/YABOIpy/-GoP256/assets/110062350/5f1edbba-7cf0-4afa-9434-f8cb27a73dee)


# Understanding The Public Key Exchange 
```go
// Gop256KE uses the p256 EC("elliptic Curve") for the Key Exchange aslo known as ECDH("Elliptic Curve Diffie Hellman") 

// P256 returns a Curve which implements NIST P-256 (FIPS 186-3, section D.2.3),
// also known as secp256r1 or prime256v1. The CurveParams.Name of this Curve is
// "P-256".

Curve := elliptic.P256()
Curve.Params().N  // returns the starting point of the elliptic curve
Curve.Params().P // returns the Prime Modulus of the elliptic curve
Curve.Params().Gx 
Curve.Params().Gy
//Gx and Gy are the points on the elliptic curve

// Generating the Keys
	// Private Keys
	// Taking the Curve parameter N and Generating a 256-Bit PseudoRandom Key That is used as a Private Key 
	func CreatePrivateKey(curve elliptic.Curve) (*big.Int, error) {
		key, err := rand.Int(rand.Reader, curve.Params().N)
		if err != nil {
			return nil, err
		}
		key.Add(key, big.NewInt(1))
		return key, nil
	}
	
	// Public Shared Keys
	// Scalar multiplication takes a point from the base point of the curve("Curve.Params().N") and multiplying it by the Private Key
	AlicePublicKeyX, AlicePublicKeyY := Curve.ScalarBaseMult(alicePrivateKey.Bytes())
	BobPublicKeyX, BobPublicKeyY := Curve.ScalarBaseMult(bobPrivateKey.Bytes())


// Calculating the Secrets
	// Taking the EC(p256 Curve) & The PublicKeys X and Y and Returning The Shared X Secret of the keys
	func CalculateSecret(curve elliptic.Curve, privateKey *big.Int, pkX, pkY *big.Int) *big.Int {
		sharedX, _ := curve.ScalarMult(pkX, pkY, privateKey.Bytes())
		sharedX.Mod(sharedX, curve.Params().P)
		return sharedX
	}
	// pkX is Alice's PublicKeyX and PkY is the PublicKeyY
	// we Do the same for the other Secret but switching the Keys around and using Bobs Private Key
	// These are the 256-Bit Shared Secrets between Alice And Bob
	Secret1 := CalculateSecret(Curve, AlicePrivateKey, BobPublicKeyX, bobPublicKeyY)
	Secret2 := CalculateSecret(Curve, BobPrivateKey, AlicePublicKeyX, alicePublicKeyY)
	// and there you have it Key Exchange done fast and efficiently

```

# BenchMarks
```md

> 1000 iterations done within 167
Benchmark finished in 166.5962ms
Benchmark finished in 166.93ms
Benchmark finished in 164.0544ms
Benchmark finished in 164.7368ms
Benchmark finished in 164.5691ms
Benchmark finished in 164.8557ms
Benchmark finished in 165.3431ms
Benchmark finished in 155.8193ms
Benchmark finished in 158.9972ms
Benchmark finished in 165.3473ms
Benchmark finished in 164.9142ms
Benchmark finished in 163.3916ms
Benchmark finished in 166.0719ms
AVG: 163.9079ms

```go
func main() {
	s := time.Now()
	
	for i := 0; i < 1000; i++ {
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
		if Secret1.Cmp(Secret2) != 0 {
			log.Println("Shared secrets do not match!")
		}
	}
	fmt.Printf("1000 runs in %s\n", time.Since(s))
}


```
# Visual Example
![image](https://github.com/YABOIpy/GoP256KE/assets/110062350/3b0c9a68-41bf-4a6c-a3e6-bfd4f0847c75)


# Other Uses of the Curve
![image](https://github.com/YABOIpy/GoP256KE/assets/110062350/f3c95a7d-dd78-4bd4-a72c-0588ad27db9c)
https://www.cem.me/20170410-ecc-1.html
_______

https://www.youtube.com/watch?v=NF1pwjL9-DE

https://en.wikipedia.org/wiki/Prime_number

https://en.wikipedia.org/wiki/Elliptic-curve_cryptography

https://en.wikipedia.org/wiki/Curve25519

