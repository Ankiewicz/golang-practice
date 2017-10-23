package main

import (
  "fmt"
  "os"
)

func check(e error){
  if e != nil {
    panic(e)
  }
}

func main() {
  // dat, err := ioutil.ReadFile("2015_Gaz_zcta_national.txt")
  // check(err)
  // fmt.Print(string(dat))

  f, err := os.Open("2015_Gaz_zcta_national.txt")
  check(err)

  // b1 := make([]byte, 5)
  // n1, err := f.Read(b1)
  // check(err)
  // fmt.Println("%d bytes: %s\n", n1, string(b1))

  o2, err := f.Seek(6,0)
  check(err)
  b2 := make([]byte, 2)
  n2, err := f.Read(b2)
  fmt.Println("%d bytes @ %s\n", n2, o2, string(b2))
}
