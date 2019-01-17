// ==== RULES ====
// Any live cell with fewer than two live neighbors dies, as if by underpopulation.
// Any live cell with two or three live neighbors lives on to the next generation.
// Any live cell with more than three live neighbors dies, as if by overpopulation.
// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

/*
Implementation of the game of life by Zetsuban
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	termbox "github.com/nsf/termbox-go"
	"github.com/zetsuban/game-of-life/entities"
)

var col = flag.Int("c", -1, "number of columns")
var row = flag.Int("r", -1, "number of rows")
var glider = flag.Bool("g", false, "use the glider as start")
var tumbler = flag.Bool("t", false, "use the tumbler as start")
var exploder = flag.Bool("e", false, "use the exploder as start")
var line = flag.Bool("l", false, "use the line as start")
var step = flag.Bool("S", false, "step by step")
var sleepTime = flag.Int("s", 250, "ms between each generations")
var cmd = flag.Bool("cmd", false, "display without using termbox-go")
var help = flag.Bool("h", false, "displays help")

var infinit = true

// Usage displays helpfull message
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: life [options]\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func displayCmd(world entities.World, gen int) {
	for i := 0; i < 200; i++ {
		fmt.Println()
	}

	fmt.Print("Generation : ", gen, "\n\n")
	for i := 0; i < world.Row; i++ {
		for j := 0; j < world.Col; j++ {
			fmt.Print(world.Cells[i][j].Str() + " ")
		}
		fmt.Println()
	}
}

// displayWorld : displays world using termbox
func displayWorld(world entities.World) {
	for i := 0; i < world.Row; i++ {
		for j := 0; j < world.Col; j++ {
			if world.Cells[i][j].Alive() {
				termbox.SetCell(j, i, world.Cells[i][j].Rune(), termbox.ColorDefault, termbox.ColorMagenta)
			} else {
				termbox.SetCell(j, i, world.Cells[i][j].Rune(), termbox.ColorDefault, termbox.ColorDefault)
			}
		}
	}

	termbox.Flush()
}

// Sleep : delay the next generation
func Sleep() {
	if *step {
		for {
			ev := termbox.PollEvent()
			if ev.Ch == 'p' {
				return
			}
			if ev.Key == termbox.KeySpace {
				return
			}
		}
	} else {
		time.Sleep(time.Duration(*sleepTime) * time.Millisecond)
	}
}

func main() {
	flag.Usage = Usage
	flag.Parse()

	if *help {
		Usage()
	}

	go inputLoop()

	if *cmd {
		CmdMain()
	} else {
		TBMain()
	}

}

// CmdMain : Cmd main loop
func CmdMain() {
	gen := 0

	// Default value if user didn't specify size
	if *row < 0 {
		*row = 50
	}
	if *col < 0 {
		*col = 50
	}

	// Generate new world and base alive cells
	world := entities.NewWorld(*row, *col)
	if *glider {
		world.Glider()
	} else if *tumbler {
		world.Tumbler()
	} else if *exploder {
		world.Exploder()
	} else if *line {
		world.Line()
	} else {
		world.Start()
	}

	Play(world, gen)

}

// TBMain : Termbox main loop
func TBMain() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	// var
	gen := 0

	// Get screen dimensions
	termCols, termRows := termbox.Size()

	// Use term size if user didn't specify
	if *row < 0 {
		*row = termRows
	}
	if *col < 0 {
		*col = termCols
	}

	// Generate new world and base alive cells
	world := entities.NewWorld(*row, *col)
	if *glider {
		world.Glider()
	} else if *tumbler {
		world.Tumbler()
	} else if *exploder {
		world.Exploder()
	} else {
		world.Start()
	}

	termbox.SetInputMode(termbox.InputEsc)
	// termbox.HideCursor()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	Play(world, gen)

}

// Play : Main play loop
func Play(world entities.World, gen int) {
	// TODO : Revisite loop

	for infinit {
		if *cmd {
			displayCmd(world, gen)
		} else {
			displayWorld(world)
		}

		world = world.NextGen()

		Sleep()

		gen++
	}
}
