package main

import (
  "fmt"
  "log"
  "bufio"
  "os"
  "strings"
)

func main() {
  file, err := os.Open(os.Args[1])
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    //'scanner.Text()' represents the test case, do something with it
    line := scanner.Text()
    a1, a2 := parseLine(line)
    result := SetIntersection(a1, a2)
    fmt.Println(strings.Join(result, ","))
  }
}

func SetIntersection(a1 []string, a2 []string) []string {
  mapped := toMap(a2)
  var result []string
  for _, v := range a1 {
    _, ok := mapped[v]
    if ok {
      result = append(result, v)
    }
  }
  return result
}

func toMap(a [] string) map[string]string {
  res := make(map[string]string)
  for _, v := range a {
    res[v] = v
  }
  return res
}

func parseLine(line string) ([]string, []string) {
  elements := strings.Split(line, ";")
  el1 := strings.Split(elements[0], ",")
  el2 := strings.Split(elements[1], ",")
  return el1, el2
}
