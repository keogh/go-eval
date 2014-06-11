package main

import (
  "testing"
)

type testpair struct {
  a1 []string
  a2 []string
  result []string
}

var tests = []testpair{
  {[]string{"1","2","3","4"}, []string{"4","5","6"}, []string{"4"}},
  {[]string{"20","21","22"}, []string{"45","46","47"}, []string{}},
  {[]string{"7","8","9"}, []string{"8","9","10","11","12"}, []string{"8","9"}},
}

func TestSetIntersection(t *testing.T) {
  for _, pair := range tests {
    v := SetIntersection(pair.a1, pair.a2)
    if !StringsEquals(v, pair.result) {
      t.Error(
        "For", pair.a1, ",", pair.a2,
        "expected", pair.result,
        "got", v,
      )
    }
  }
}

func StringsEquals(a, b []string) bool {
  if len(a) != len(b) {
    return false
  }
  for i, v := range a {
    if v != b[i] {
      return false
    }
  }
  return true
}
