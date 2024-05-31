package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type play struct{
  gameData [3][3]uint8
  turn bool
  playAble bool
  host bool
}

const (
  EmptyCell uint8 = iota
  PlayerX 
  PlayerO 
)

  /*
  x | o | x 
 ---|---|---
  o | o | x 
 ---|---|---
  x | x | o 
*/

func whatToPrint(o uint8) string{
  toPrint := ""
  if o == EmptyCell {
    toPrint += fmt.Sprintf("   ")
  }else if o == PlayerX{
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

func (p *play) startGame(conn net.Conn){
  for {
    if p.playAble{
      fmt.Println(p.whoesTrun(), "player enter 1 to 9 to input")
      p.buildBord()
      row, col := p.inputTurn(conn)
      if p.turn{
        p.gameData[row][col] = PlayerX
      }else{
        p.gameData[row][col] = PlayerO
      }
      p.turn = !p.turn
      p.gameWon()
    }else{
      p.buildBord()
      break
    }
  }
}

func (p *play) inputTurn(conn net.Conn) (uint8, uint8){
  var input uint8
  var row, col uint8
  var err error
  reader := bufio.NewReader(conn)
  for{
    if p.host && p.turn{
      _, err = fmt.Scan(&input)
      if err != nil{
        fmt.Println("Please enter the number")
        continue
      }
      conn.Write([]byte{input})
    }else if p.host && !p.turn{
      fmt.Println("Wating for player O")
      input, err = reader.ReadByte()
      if err != nil{
        fmt.Println("Cannot read from the network terminating game")
        p.playAble = false
        return 0, 0
      }
    }else if !p.host && p.turn{
      fmt.Println("Wating for player X")
      input, err = reader.ReadByte()
      if err != nil{
        fmt.Println("Cannot read from the network terminating game")
        p.playAble = false
        return 0, 0
      }
    }else{
      _, err = fmt.Scan(&input)
      if err != nil{
        fmt.Println("Please enter the number")
        continue
      }
      conn.Write([]byte{input})
    }
    if input >= 0 && input <= 9{
      row, col = p.inputToGameData(input - 1)
      if p.alreadyInput(row, col){
        fmt.Println("Input already taken")
        continue
      }
    }
    break
  }
  return row, col
}

func (p *play) inputToGameData(input uint8) (uint8, uint8){
  row := (input / 3) 
  col := (input % 3) 
  return row, col
}

func (p *play) alreadyInput(row, col uint8) bool{
  if p.gameData[row][col] != EmptyCell{
    return true
  }
  return false
}

// I can use just a loop here. Noting but overthinking 
func (p *play) gameWon() {
  for i := 0; i < 3; i++{
    if p.gameData[i][0] != EmptyCell && p.gameData[i][0] == p.gameData[i][1] && p.gameData[i][1] == p.gameData[i][2]{
      p.playAble = false
      if p.gameData[i][0] == PlayerX{
        fmt.Println("X won the game")
      }else{ 
        fmt.Println("O won the game")
      }
    }else if p.gameData[0][i] != EmptyCell && p.gameData[0][i] == p.gameData[1][i] && p.gameData[1][i] == p.gameData[2][i]{
      p.playAble = false
      if p.gameData[0][i] == PlayerX{
        fmt.Println("X won the game")
      }else{ 
        fmt.Println("O won the game")
      }
    }
  }
  // X this line checking
  if p.gameData[0][0] != EmptyCell && p.gameData[0][0] == p.gameData[1][1] && p.gameData[1][1] == p.gameData[2][2]{
      p.playAble = false
      if p.gameData[0][0] == PlayerX{
        fmt.Println("X won the game")
      }else{ 
        fmt.Println("O won the game")
      }
  } 

  if p.gameData[0][2] != EmptyCell && p.gameData[0][2] == p.gameData[1][1] && p.gameData[1][1] == p.gameData[2][0]{
      p.playAble = false
      if p.gameData[0][2] == PlayerX{
        fmt.Println("X won the game")
      }else{ 
        fmt.Println("O won the game")
      }
  } 
}

func createTCP() (net.Conn, error){
  listen, err := net.Listen("tcp", ":7000")
  if err != nil{
    return nil, err
  }
  conn, err := listen.Accept()
  if err != nil{
    return nil, err
  }
  return conn, nil
}

func joinTCP() (net.Conn, error){
  conn, err := net.Dial("tcp", ":7000")
  if err != nil{
    return nil, err
  }
  return conn, nil
}

func readFromSocket(reader *bufio.Reader)(uint8, error){
  input, err := reader.ReadByte()
  if err != nil{
    return input, err
  }
  return input, nil
}

func WriteToSocket(conn net.Conn, value uint8) (error){
  _, err := conn.Write([]byte{value})
  if err != nil{
    return err
  }
  return nil
}

func main(){
  host := flag.Bool("host", false, "To host the game")
  flag.Parse()
  var conn net.Conn
  var err error
  if *host{
    conn, err = createTCP()
    if err != nil{
      fmt.Println("Unable to host the game: " , err)
      return
    }
  }else{
    conn, err = joinTCP()
    if err != nil{
      fmt.Println("Unable to join the host game: " , err)
      return
    }
  }

  game := &play{
    turn: true,
    playAble: true,
    host: *host,
  }
  game.startGame(conn)
}
