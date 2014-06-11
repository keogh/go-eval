package main

import "testing"

type testpair struct {
  x int
  n int
  multiple int
}

var tests = []testpair{
  {13, 8, 16},
  {17, 16, 32},
  {1731, 64, 1792},
}

func TestMultiplesOfANumber(t *testing.T) {
  for _, pair := range tests {
    v := MultiplesOfANumber(pair.x, pair.n)
    if v != pair.multiple {
      t.Error(
        "For", pair.x, ",", pair.n,
        "expected", pair.multiple,
        "got ", v,
      )
    }
  }
}
