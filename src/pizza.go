package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)


func check(e error) {
  if e != nil {
      fmt.Println(e)
  }
}

func main() {
  // flour input
  flourQ := bufio.NewReader(os.Stdin)
  fmt.Print("\nHow many grams of flour are you going to use? > ")
  flour, err := flourQ.ReadString('\n')
  check(err)
  flourFloat, err = strconv.ParseFloat(flour)
  check(err)

  // room temperature input
  roomTempQ := bufio.NewReader(os.Stdin)
  fmt.Print("\nWhat is the actual room temperature? > ")
  roomTemp, err := roomTempQ.ReadString('\n')
  check(err)
  roomTempFloat, err = strconv.ParseFloat(roomTemp)
  check(err)

  loop:
  for {
    // pizza type input
    pizzaTypeQ := bufio.NewReader(os.Stdin)
    fmt.Print("\nDo you want make C(lassic) Pizza or B(iga) Pizza? >")
    pizzaType, err := pizzaTypeQ.ReadString('\n')
    check(err)
    // switch expression
    switch pizzaType {
    case "C\n":
      pizza(flourFloat, roomTempFloat)
      break loop
    case "B\n":
      biga(flourFloat, roomTempFloat)
      break loop
    default:
      fmt.Println("\nPlease type \"C\" for Classic Piza or \"B\" for Biga Pizza")
    }
  }
}

func pizza(f,rt float64) {
  // temperature variables - Napels style
  fmt.Println(f,rt)
//
//   // dough ingredients quantity
//   doug_g_flour = flour
//   doug_g_water = flour * (66/100)
//   doug_g_yeast = doug_g_flour * (0.1/100)
//   doug_g_salt = flour * (2/100)
}
//
func biga(f,rt float64) {
  fmt.Println(f,rt)
}
