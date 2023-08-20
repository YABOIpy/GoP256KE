# Go25519
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
	g.primes = append(g.primes, 2)
	for i := 1; i < limit/2; i++ {
		if !c[i] {
			g.primes = append(g.primes, 2*i+1)
		}
	}
}
```

![image](https://github.com/YABOIpy/-GoP256/assets/110062350/5f1edbba-7cf0-4afa-9434-f8cb27a73dee)


```go
// P256 returns a Curve which implements NIST P-256 (FIPS 186-3, section D.2.3),
// also known as secp256r1 or prime256v1. The CurveParams.Name of this Curve is
// "P-256".


https://www.youtube.com/watch?v=NF1pwjL9-DE
https://en.wikipedia.org/wiki/Prime_number
https://en.wikipedia.org/wiki/Curve25519

