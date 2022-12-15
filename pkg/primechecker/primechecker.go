package primechecker

import "math/big"

// PrimeChecker check if number is a prime or not
type PrimeChecker struct{}

func (p *PrimeChecker) Check(number int) bool {
	n := int64(number)

	return big.NewInt(n).ProbablyPrime(0)
}
