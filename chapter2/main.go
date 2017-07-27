package main

import "fmt"

//this is a comment

func main() {
  // var y string = "Hello World"
  // fmt.Println(y)
  //
  // var x string
  // x = "first "
  // fmt.Println(x)
  // x = x + "second"
  // fmt.Println(x)

  fmt.Print("Enter a number: ")
   var input float64
   fmt.Scanf("%f", &input)

   output := input * 2

   fmt.Println(output)

}
