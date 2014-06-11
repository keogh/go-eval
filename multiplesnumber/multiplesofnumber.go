package main

import (
  "fmt"
  "log"
  "bufio"
  "os"
  "strconv"
  "strings"
)

func main () {
  file, err := os.Open(os.Args[1])
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    //'scanner.Text()' represents the test case, do something with it
    line := scanner.Text()
    x, n := getNumbers(line)
    fmt.Println(MultiplesOfANumber(x, n))
  }
}

func MultiplesOfANumber(x int, n int) int {
  res := n
  for res < x {
    res = res + n
  }
  return res
}

func getNumbers(line string) (int, int) {
  elements := strings.Split(line, ",")
  x, _ := strconv.Atoi(elements[0])
  n, _ := strconv.Atoi(elements[1])
  return x, n
}
