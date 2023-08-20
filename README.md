# GoP256KE
Implementation of the P-256 curve for key exchange in Go


https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
```go
func (g *Generator) GeneratePrimes(limit int) {
	if limit <= 3 {
		return
	}
	c := make([]bool, limit)
	
	for i := 1; i < int(1+(1.0*float64(limit-1)/2.0)); i++ {
		if !c[i] {
			p := 2*i + 1
			for j := p * p; j < limit; j += 2 * p {
				c[(j-1)/2] = true
			}
		}
	}
	//adding them to the list
	g.primes = append(g.primes, 2)
	for i := 1; i < limit/2; i++ {
		if !c[i] {
			g.primes = append(g.primes, 2*i+1)
		}
	}
}
```

![image](https://github.com/YABOIpy/-GoP256/assets/110062350/5f1edbba-7cf0-4afa-9434-f8cb27a73dee)


# Understanding The KeyExchange 
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
	// Taking the Curve parameter N and Generating a PseudoRandom Key That is used as a Private Key 
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
	sharedSecret1 := CalculateSecret(Curve, AlicePrivateKey, BobPublicKeyX, bobPublicKeyY)
	sharedSecret2 := CalculateSecret(Curve, BobPrivateKey, AlicePublicKeyX, alicePublicKeyY)
	// and there you have it Key Exchange done fast and efficiently

```
# BenchMarks
```md

```

https://www.youtube.com/watch?v=NF1pwjL9-DE
https://en.wikipedia.org/wiki/Prime_number
https://en.wikipedia.org/wiki/Elliptic-curve_cryptography
https://en.wikipedia.org/wiki/Curve25519

