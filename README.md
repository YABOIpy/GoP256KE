# Go25519
Implementation of the Ed25519 Public-Key Cryptography Algorithm in Go



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


1. https://en.wikipedia.org/wiki/Prime_number
2. https://en.wikipedia.org/wiki/Curve25519
3. https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
