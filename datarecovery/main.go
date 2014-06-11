package main

import(
  "log"
  "bufio"
  "os"
  "fmt"
  "strings"
  "strconv"
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
  mapped := mapData(text, code)
  var res []string
  for i:= 1; i <= len(text); i++ {
    res = append(res, mapped[strconv.Itoa(i)])
  }
  return strings.Join(res, " ")
}

func parseLine(line string) ([]string, []string){
  elements := strings.Split(line, ";")
  return strings.Split(elements[0], " "), strings.Split(elements[1], " ")
}

func normalize(text *[]string, code *[]string) {
  lenText := len(*text)
  lenCode := len(*code)
  if lenText <= lenCode {
    return
  }
  for i := 1; i <= lenText; i++ {
    if (inArray(*code, strconv.Itoa(i))) {
      continue
    }
    *code = append(*code, strconv.Itoa(i))
    if lenText == len(*code) {
      break
    }
  }
}

func inArray(haystack [] string, needle string) bool {
  for _, v := range haystack {
    if needle == v {
      return true
    }
  }
  return false
}

func mapData(text []string, code []string) map[string]string {
  mapped := make(map[string]string)
  for i := 0; i < len(text); i++ {
    mapped[code[i]] = text[i]
  }
  return mapped
}
