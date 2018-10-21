package main

import "fmt"

func main(){
  var x float64 = 20.0

  // variable init and decl short version //
  // valid only within main func //
  // variable type not defined. Iterpreter //
  // will define type during the compile //
  y := 42

  fmt.Println(x)
  fmt.Printf("x is type %T\n", x)
  fmt.Printf("y is type %T\n", y)
}
