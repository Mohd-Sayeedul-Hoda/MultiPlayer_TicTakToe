package main

import(
  "fmt"
)

type play struct{
  gameData [3][3]int8
  turn bool
}
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
    toPrint += fmt.Sprintf(" X ")
  }else{ 
    toPrint += fmt.Sprintf(" O ")
  }
  return toPrint
}


func (p *play) buildBord() {
  for x := 0; x <= 2; x++{
    var toPrint string
    for y := 0; y <= 2; y++{
      toPrint += whatToPrint(p.gameData[x][y])
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

func (p *play) whoesTrun() string{
  if p.turn {
    return "X"
  }else{
    return "O"
  } 
}

func (p *play) inputTurn() {
  var input int8
  fmt.Println(p.whoesTrun(), "player enter 1 to 9 to input")
  p.buildBord()
  for{
    _, err := fmt.Scan(&input)
    if err != nil{
      fmt.Println("Please enter the number")
      continue
    }
    if input >= 0 && input <= 9{
      break
    }
    fmt.Println("Enter the input between 1 to 9 ")
  }
  row, col := p.inputToGameData(input - 1)
  if p.turn{
    p.gameData[row][col] = 1
  }else{
    p.gameData[row][col] = 2
  }
  p.turn = !p.turn
}

func (p *play) inputToGameData(input int8) (int8, int8){
  row := (input / 3) 
  col := (input % 3) 
  return row, col
}

func main(){
  p := &play{
    turn: true,
  }
  for{
    p.inputTurn()
  }
}
