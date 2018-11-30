package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "time"
  "strings"
  "log"
)

func PercentInt(pcent int, all int) float64{
  percent := ((float64(all) * float64(pcent)) / float64(100))
  return percent
}

func PercentFloat(pcent float64, all int) float64{
  percent := ((float64(all) * pcent) / float64(100))
  return percent
}

func main() {
  // flour input
  flourQ := bufio.NewReader(os.Stdin)
  fmt.Print("\nHow many grams of flour are you going to use? > ")
  flour, err := flourQ.ReadString('\n')
  if err != nil {
    log.Fatal(fmt.Println("ERROR - Cannot read from stdin"))
  }

  flourParsed := strings.Replace(flour, "\n", "", -1)
  flourInt, err := strconv.Atoi(flourParsed)
  if err != nil {
    log.Fatal(fmt.Println("ERROR - Value must be an integer"))
  }

  // room temperature input
  roomTempQ := bufio.NewReader(os.Stdin)
  fmt.Print("\nWhat is the actual room temperature? > ")
  roomTemp, err := roomTempQ.ReadString('\n')
  if err != nil {
    log.Fatal(fmt.Println("ERROR - Cannot read from stdin"))
  }

  roomTempParsed := strings.Replace(roomTemp, "\n", "", -1)
  roomTempInt, err := strconv.Atoi(roomTempParsed)
  if err != nil {
    log.Fatal(fmt.Println("ERROR - Value must be an integer"))
  }

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
      pizza(flourInt, roomTempInt)
      break loop
    case "B\n":
      biga(flourInt, roomTempInt)
      break loop
    default:
      fmt.Println("\nPlease type \"C\" for Classic Piza or \"B\" for Biga Pizza")
    }
  }
}

func pizza(f,rt int) {
  // temperature variables - Napels style
  ingredientsTemp := rt - 1
  constant := 9
  waterTemp := 75 - rt - ingredientsTemp - constant
  timePizzaReady := time.Now().Add(6*time.Hour)

  // dough ingredients quantity
  dougFlour := f
  var dougWater float64 = PercentInt(66, dougFlour)
  var dougYeast float64 = PercentFloat(0.1, dougFlour)
  var dougSalt float64 = PercentInt(2, dougFlour)
  starterFlour := PercentInt(10, dougFlour)

  // procedure
  fmt.Printf(`
    ########################################################################
    ### THIS IS THE PERFECT RECIPE FOR MAKING A PERFECT PIZZA AT HOME   ###
    ########################################################################

    INTRO:
    All ingredients are expressed in GRAMS.
    For liquids like water, oil, malt you can safely use MILLILITERS.
    Use fresh yeast, not dry yeast.
    Flour must have the following requirement: 0,40 < P/L < 0,60; W > 300

    If you want a good result, you need to use good quality ingredients. This
    is the secret. If you buy first price stuff, probably your pizza would
    taste like the Pizza Hut/Papa Johns/Domino ones.

    N.B
    Water temperature must be %v

    PIZZA DOUGH INGREDIENTS:
    Flour: %v
    Salt: %v
    Water: %v
    Yeast: %v

    PREPARATION:
    For PIZZA DOUGH: melt the yeast into the water add roughly %v grams of flour and knead it.
    Add the salt and make sure it get completely absorbed then add flour and knead the ingredients for 5 minutes.

    Leave the dough rest for 4-6 hours at 15-16°C. Pizza dough will be ready %v Higher the temperature, lower will be the resting time.
    Cook in the oven at 350°C - static.
    `, waterTemp, dougFlour, dougSalt, dougWater, dougYeast, starterFlour, timePizzaReady)
}

func biga(f,rt int) {
  fmt.Println(f,rt)
}
