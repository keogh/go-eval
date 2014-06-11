package main

import(
  "log"
  "bufio"
  "os"
  "fmt"
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
    fmt.Println(recoverData(line))
  }
}

func recoverData(line string) string {
  text, code := parseLine(line)
  normalize(&text, &code)
  return line
}

func parseLine(line string) ([]string, []string){
  elements := strings.Split(line, ";")
  return strings.Split(elements[0], " "), strings.Split(elements[1], " ")
}

func normalize(text *[]string, code *[]string) {
  lenText := len(*text)
  lenCode := len(*code)
  if lenText == lenCode {
    return
  }
  
}
