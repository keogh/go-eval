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
    line := strings.Trim(scanner.Text(), " ")
    if line == "" {
      continue
    }
    fmt.Println(reverseWords(line))
  }
}

func reverseWords(line string) string {
  words := strings.Split(line, " ")
  var reversed []string
  maxIndex := len(words) - 1
  for i := maxIndex; i >= 0; i-- {
    reversed = append(reversed, words[i])
  }
  return strings.Join(reversed, " ")
}
