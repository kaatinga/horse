package main

import (
	"fmt"
	"log"
	"strconv"
)

type board struct {
	pieces map[*cell]*piece
}

func (b *board) addPiece(p *piece, pos cell) {
	if p.pieceType != "Horse" {
		log.Println("Incorrect piece type")
		return
	}
	b.pieces[&pos] = p
	b.pieces[&pos].cell = &pos
}

func (b *board) print() {

	tmpMap := [8][8]*piece{}
	possibleMoves := [8][8]string{}
	for key, value := range b.pieces {
		tmpMap[key.x][key.y] = value
		{
			// быстрый колхоз для отображения возможных шагов

			for _, value := range (*value).findSteps() {
				possibleMoves[value.x][value.y] = "."
			}
		}
	}

	fmt.Println("┌───┬───┬───┬───┬───┬───┬───┬───┬───┐")
	fmt.Println("│   │ a │ b │ c │ d │ e │ f │ g │ h │")
	fmt.Println("├───┼───┼───┼───┼───┼───┼───┼───┼───┤")
	for keyY := 0; keyY < 8; keyY++ {
		fmt.Printf("│ %d │", keyY+1)

		for keyX := 0; keyX < 8; keyX++ {
			if tmpMap[keyX][keyY] != nil {
				fmt.Printf(" %s ", tmpMap[keyX][keyY].name)
			} else if possibleMoves[keyX][keyY] != "" {
				fmt.Printf(" %s ", possibleMoves[keyX][keyY])
			} else {
				fmt.Print("   ")
			}
			fmt.Print("│")
			if keyX == 7 && keyY != 7 {
				fmt.Println()
				fmt.Println("├───┼───┼───┼───┼───┼───┼───┼───┼───┤")
			}
		}
	}
	fmt.Println()
	fmt.Println("└───┴───┴───┴───┴───┴───┴───┴───┴───┘")
}

func (b *board) init() {
	for keyY := 0; keyY < 8; keyY++ {
		for keyX := 0; keyX < 8; keyX++ {
			(*b).pieces[&cell{keyX, keyY}] = &piece{name: " ", pieceType: "None"}
		}
	}
}

type cell struct {
	x int
	y int
}

type piece struct {
	name        string
	pieceType   string
	jumpPattern []cell
	*cell
}

func (p *piece) newPiece(pieceType string) {
	p.name = string(pieceType[0])
	p.pieceType = pieceType
	p.jumpPattern = jumps[pieceType]
}

func (b *piece) findSteps() (possibleJumps []cell) {
	var target cell
	for _, value := range b.jumpPattern {
		target.x = (*b).x + value.x
		target.y = (*b).y + value.y
		if !(target.x < 0 || target.x > 7 || target.y < 0 || target.y > 7) {
			possibleJumps = append(possibleJumps, target)
		}
	}
	return
}

var jumps map[string][]cell

func (target *piece) move(x, y int) {
	newTarget, ok := target.check(x, y)
	if ok {
		*(*target).cell = newTarget
		return
	}
	fmt.Println("Impossible to move there...")
}

func (target *piece) check(x, y int) (where cell, ok bool) {
	newX := (*target).x + x
	newY := (*target).y + y

	if !(newX < 0 || newX > 7 || newY < 0 || newY > 7) {
		return cell{newX, newY}, true
	}
	return
}

func (p *piece) print() {

	fmt.Printf("The current position is %s\n", mapping[(*p).x]+strconv.Itoa((*p).y+1))
	fmt.Print("The possible moves are: ")
	for _, value := range p.findSteps() {

		fmt.Printf("%s%s ", mapping[value.x], strconv.Itoa(value.y+1))
	}
	fmt.Println()
}

var mapping [8]string

func main() {

	mapping = [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}

	p := fmt.Println
	p("Welcome to chess possible step finder ")

	testBoard := board{pieces: make(map[*cell]*piece, 64)}
	//testBoard.init()

	jumps = make(map[string][]cell, 6)

	jumps["Horse"] = []cell{
		cell{1, 2},
		cell{2, 1},
		cell{2, -1},
		cell{1, -2},
		cell{-1, -2},
		cell{-2, -1},
		cell{-2, 1},
		cell{-1, 2},
	}

	jumps["Pown"] = []cell{
		cell{1, -1},
		cell{1, 0},
		cell{1, 1},
		cell{0, -1},
		cell{0, 1},
		cell{-1, -1},
		cell{-1, 0},
		cell{-1, 1},
	}

	p("The application has added a horse on the board")

	var myHorse piece
	myHorse.newPiece("Horse")
	testBoard.addPiece(&myHorse, cell{1, 1})

	testBoard.print()
	myHorse.print()

	// вводим действие
	var action string
	for {
		p("Please, enter 'a', 'd', 'w' or 's' in order to move the horse (H) on the board. Enter 'exit' if you want to exit from the application")
		fmt.Scan(&action)
		switch action {
		case "a":
			p("You are moving the horse left")
			myHorse.move(-1, 0)
		case "w":
			p("You are moving the horse up")
			myHorse.move(0, -1)
		case "s":
			p("You are moving the horse down")
			myHorse.move(0, 1)
		case "d":
			p("You are moving the horse right")
			myHorse.move(1, 0)
		default:
			p("The data input has been wrong. Please, repeat:")
		}
		testBoard.print()
		myHorse.print()
	}
}
