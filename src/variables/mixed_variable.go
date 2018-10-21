package main

import "fmt"

func main(){
  var a, b, c = 3, 4, "foo"

  fmt.Println(a, b, c)
  fmt.Printf("%T %T %T \n", a, b, c)
}
