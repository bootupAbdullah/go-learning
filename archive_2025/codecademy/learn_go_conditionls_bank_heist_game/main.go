package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  
  // --> print statment for Codecademy designation:
  fmt.Println("Codecademy, Learn Go: Conditionals")

  // --> 

  // --> tools:
  // -- 1. 'dialogue' designation
  var dialogue string
  dialogue = "\nDialogue: "

  // -- 3. 'test' designation
  var test string
  test = "\nTest: "

  // -- 2. toggle visbility of 'test' print statements 
  var printTestLine bool
  printTestLine = true

  // ----> end: tools


  // ---> start: main script start
  // #1. generate random numbers:
  rand.Seed(time.Now().UnixNano())

  // #2. declare isHeistOn variable
  var isHeistOn bool
  isHeistOn = true
  //----> TEST ----> type: var: 
  if printTestLine {
    fmt.Printf(test + "line 39, prints var 'isHeistOn': %t.", isHeistOn)
  }

  // #3. crete var eludedGuards
  var eludedGuards int 
  eludedGuards = rand.Intn(100)

  ///----> TEST ----> type: var:
  if printTestLine {
    fmt.Printf(test + "line 47, prints var 'eludedGuards': %d\n", eludedGuards)
  }

  //--> #4. first conditional
  if eludedGuards <= 50 {
    //----> TEST ----> type: var:
    if printTestLine {
      fmt.Println(eludedGuards)
    }
    fmt.Println("\n[Game Start]")
    fmt.Printf(dialogue + "Begin Heist!\n")
    fmt.Println(dialogue + "Looks like you made it past the guards! They were more interested in the watching the game. You've made it to the next round.")

  // --> #5. add else condition, change isHesitOn to false
  } else {
    isHeistOn = false
    fmt.Println(dialogue + "Heist over before it began!")
    fmt.Println(dialogue + "You didn't make it past the guards!They got you big dawg, you gotta be quick the moves.")
  } // --> end first conditional, starts line 28

  //#7. declare openedVault var:
  var openedVault int
  openedVault = rand.Intn(100)

  //----> TEST ----> type: var:
  if printTestLine {
    fmt.Printf(test + "line 70, print var 'openedVault': %d", openedVault)
  }

  //#8. check if isHeistOn = true && openedVault >= 70:
  if isHeistOn && openedVault >=70 {
    fmt.Println(dialogue + "The vault is open, grab and go!")
  //#9. add an else if statement to check isHesitOn:
  } else if isHeistOn && openedVault < 70 {
    isHeistOn = false
    fmt.Printf(dialogue + "That vault can't be opened!")
  }
  

  // --> #6. purposefully placed at bottom of code. View current status of isHeistOn:
  //----> TEST ----> type: bool:
  if printTestLine {
    fmt.Printf(test + "line 87, Current status of isHeistOn: %t", isHeistOn)
  }

} // --> end main func
