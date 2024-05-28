package main

import(
  "fmt"
)

var gameData [3][3]int8

  /*
  x | o | x 
 ---|---|---
  o | o | x 
 ---|---|---
  x | x | o 
*/

func whatToPrint(o int8) string{
  toPrint := ""
  if o == 0 {
    toPrint += fmt.Sprintf("   ")
  }else if o == 1{
    toPrint += fmt.Sprintf(" x ")
  }else{ 
    toPrint += fmt.Sprintf(" o ")
  }
  return toPrint
}


func buildGame() {
  for x := 0; x <= 2; x++{
    var toPrint string
    for y := 0; y <= 2; y++{
      toPrint += whatToPrint(gameData[x][y])
      if y != 2{
        toPrint += "|"
      }
    }
    fmt.Println(toPrint)
    if x != 2{
      fmt.Println("---|---|---")
    }
  }     
}

func main(){
  gameData[1][1] = 2
  buildGame()
}
