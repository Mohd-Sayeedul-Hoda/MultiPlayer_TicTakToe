package main

import(
  "fmt"
)

type play struct{
  gameData [3][3]int8
  turn bool
  playAble bool
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
  var row, col int8
  for{
    _, err := fmt.Scan(&input)
    if err != nil{
      fmt.Println("Please enter the number")
      continue
    }
    if input >= 0 && input <= 9{
      row, col = p.inputToGameData(input - 1)
      if p.alreadyInput(row, col){
        fmt.Println("Input already taken")
        continue
      }
      break
    }
  }
  if p.turn{
    p.gameData[row][col] = 1
  }else{
    p.gameData[row][col] = 2
  }
  p.turn = !p.turn
  p.gameWon()
}

func (p *play) inputToGameData(input int8) (int8, int8){
  row := (input / 3) 
  col := (input % 3) 
  return row, col
}

func (p *play) alreadyInput(row, col int8) bool{
  if p.gameData[row][col] == 1 || p.gameData[row][col] == 2{
    return true
  }
  return false
}

func (p *play) gameWon() {
  for i := 0; i < 3; i++{
    if p.gameData[i][0] != 0 && p.gameData[i][0] == p.gameData[i][1] && p.gameData[i][1] == p.gameData[i][2]{
      p.playAble = false
      if p.gameData[i][0] == 1 {
        fmt.Println("X won the game")
      }else{ 
        fmt.Println("O won the game")
      }
    }else if p.gameData[0][i] != 0 && p.gameData[0][i] == p.gameData[1][i] && p.gameData[1][i] == p.gameData[2][i]{
      p.playAble = false
      if p.gameData[0][i] == 1 {
        fmt.Println("X won the game")
      }else{ 
        fmt.Println("O won the game")
      }
    }
  }
  // X this line checking
  if p.gameData[0][0] != 0 && p.gameData[0][0] == p.gameData[1][1] && p.gameData[1][1] == p.gameData[2][2]{
      p.playAble = false
      if p.gameData[0][0] == 1 {
        fmt.Println("X won the game")
      }else{ 
        fmt.Println("O won the game")
      }
  } 

  if p.gameData[0][2] != 0 && p.gameData[0][2] == p.gameData[1][1] && p.gameData[1][1] == p.gameData[2][0]{
      p.playAble = false
      if p.gameData[0][2] == 1 {
        fmt.Println("X won the game")
      }else{ 
        fmt.Println("O won the game")
      }
  } 


}

func main(){
  p := &play{
    turn: true,
    playAble: true,
  }
  for{
    if p.playAble{
      p.inputTurn()
    }else{
      break
    }
  }
}
