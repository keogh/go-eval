package main

import "fmt"
import "math"

func main() {
  total := 0
  totalPrimes := 0
  for i := 1; totalPrimes < 1000; i++ {
    if isPrime(i) {
      total += i
      totalPrimes++
    }
  }
  fmt.Println(total)
}

func isPrime(n int) bool {
  if n <= 1 {
    return false
  }
  upto := int(math.Sqrt(float64(n)))
  for i := 2; i <= upto; i++ {
    if n % i == 0 {
      return false
    }
  }
  return true
}
