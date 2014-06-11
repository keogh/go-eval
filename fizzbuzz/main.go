package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"

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
    elements := strings.Split(line, " ")
    result := generateFizzBuzz(elements)
    fmt.Println(result)
  }
}

func generateFizzBuzz(elements []string) string {
  a, _ := strconv.Atoi(elements[0])
  b, _ := strconv.Atoi(elements[1])
  n, _ := strconv.Atoi(elements[2])
  var res []string
  for i := 1; i <= n; i++ {
    res = append(res, parseFB(i, a, b))
  }
  return strings.Join(res, " ")
}

func parseFB(i int, a int, b int) string {
  if i % a == 0 && i % b == 0 {
    return "FB"
  } else if i % a == 0 {
    return "F"
  } else if i % b == 0 {
    return "B"
  }
  return strconv.Itoa(i)
}
