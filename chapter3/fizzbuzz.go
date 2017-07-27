package main

import "fmt"

func main() {
  x := 100

  for i := 0; i <= x; i++ {
    fmt.Println(i)
     if i % 3 == 0 && i % 5 == 0{
      fmt.Println("FIZZBUZZ -- Number: ", i, " is a mulitple of 3 & 5")
    } else if i % 5 == 0{
      fmt.Println("BUZZ -- Number: ", i, " is a mulitple of 5")
    } else if i % 3 == 0{
      fmt.Println("FIZZ -- Number: ", i,  " is a mulitple of 3")
    }
  }
}
