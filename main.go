package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"gtktest/point"
	"gtktest/queue"
	"os"
)

var maze [][]point.Point
var que queue.Queue

func init() {
	/*
		[
			0[2,1,0,0,0,0,0,1,1,1],
			1[0,0,0,1,1,1,0,0,1,1],
			2[1,1,0,0,0,1,1,0,0,0],
			3[0,0,0,1,0,0,1,1,1,1],
			4[1,1,1,1,1,0,0,0,0,0],
			5[0,0,0,1,1,1,1,1,0,1],
			6[0,1,0,1,0,0,0,1,0,1],
			7[0,1,0,0,0,1,0,0,0,1],
			8[0,1,1,1,1,1,1,1,1,1],
			9[0,0,0,0,0,0,0,0,3,1],
		]
	*/
	maze = make([][]point.Point, 10)
	maze[0] = make([]point.Point, 10)
	maze[0][0] = point.Point{0, 0, 2}
	maze[0][1] = point.Point{1, 0, 1}
	maze[0][2] = point.Point{2, 0, 0}
	maze[0][3] = point.Point{3, 0, 0}
	maze[0][4] = point.Point{4, 0, 0}
	maze[0][5] = point.Point{5, 0, 0}
	maze[0][6] = point.Point{6, 0, 0}
	maze[0][7] = point.Point{7, 0, 1}
	maze[0][8] = point.Point{8, 0, 1}
	maze[0][9] = point.Point{9, 0, 1}
	maze[1] = make([]point.Point, 10)
	maze[1][0] = point.Point{0, 1, 0}
	maze[1][1] = point.Point{1, 1, 0}
	maze[1][2] = point.Point{2, 1, 0}
	maze[1][3] = point.Point{3, 1, 1}
	maze[1][4] = point.Point{4, 1, 1}
	maze[1][5] = point.Point{5, 1, 1}
	maze[1][6] = point.Point{6, 1, 0}
	maze[1][7] = point.Point{7, 1, 0}
	maze[1][8] = point.Point{8, 1, 1}
	maze[1][9] = point.Point{9, 1, 1}
	maze[2] = make([]point.Point, 10)
	maze[2][0] = point.Point{0, 2, 1}
	maze[2][1] = point.Point{1, 2, 1}
	maze[2][2] = point.Point{2, 2, 0}
	maze[2][3] = point.Point{3, 2, 0}
	maze[2][4] = point.Point{4, 2, 0}
	maze[2][5] = point.Point{5, 2, 1}
	maze[2][6] = point.Point{6, 2, 1}
	maze[2][7] = point.Point{7, 2, 0}
	maze[2][8] = point.Point{8, 2, 0}
	maze[2][9] = point.Point{9, 2, 0}
	maze[3] = make([]point.Point, 10)
	maze[3][0] = point.Point{0, 3, 0}
	maze[3][1] = point.Point{1, 3, 0}
	maze[3][2] = point.Point{2, 3, 0}
	maze[3][3] = point.Point{3, 3, 1}
	maze[3][4] = point.Point{4, 3, 0}
	maze[3][5] = point.Point{5, 3, 0}
	maze[3][6] = point.Point{6, 3, 1}
	maze[3][7] = point.Point{7, 3, 1}
	maze[3][8] = point.Point{8, 3, 1}
	maze[3][9] = point.Point{9, 3, 1}
	maze[4] = make([]point.Point, 10)
	maze[4][0] = point.Point{0, 4, 1}
	maze[4][1] = point.Point{1, 4, 1}
	maze[4][2] = point.Point{2, 4, 1}
	maze[4][3] = point.Point{3, 4, 1}
	maze[4][4] = point.Point{4, 4, 1}
	maze[4][5] = point.Point{5, 4, 0}
	maze[4][6] = point.Point{6, 4, 0}
	maze[4][7] = point.Point{7, 4, 0}
	maze[4][8] = point.Point{8, 4, 0}
	maze[4][9] = point.Point{9, 4, 0}
	maze[5] = make([]point.Point, 10)
	maze[5][0] = point.Point{0, 5, 0}
	maze[5][1] = point.Point{1, 5, 0}
	maze[5][2] = point.Point{2, 5, 0}
	maze[5][3] = point.Point{3, 5, 1}
	maze[5][4] = point.Point{4, 5, 1}
	maze[5][5] = point.Point{5, 5, 1}
	maze[5][6] = point.Point{6, 5, 1}
	maze[5][7] = point.Point{7, 5, 1}
	maze[5][8] = point.Point{8, 5, 0}
	maze[5][9] = point.Point{9, 5, 1}
	maze[6] = make([]point.Point, 10)
	maze[6][0] = point.Point{0, 6, 0}
	maze[6][1] = point.Point{1, 6, 1}
	maze[6][2] = point.Point{2, 6, 0}
	maze[6][3] = point.Point{3, 6, 1}
	maze[6][4] = point.Point{4, 6, 0}
	maze[6][5] = point.Point{5, 6, 0}
	maze[6][6] = point.Point{6, 6, 0}
	maze[6][7] = point.Point{7, 6, 1}
	maze[6][8] = point.Point{8, 6, 0}
	maze[6][9] = point.Point{9, 6, 1}
	maze[7] = make([]point.Point, 10)
	maze[7][0] = point.Point{0, 7, 0}
	maze[7][1] = point.Point{1, 7, 1}
	maze[7][2] = point.Point{2, 7, 0}
	maze[7][3] = point.Point{3, 7, 0}
	maze[7][4] = point.Point{4, 7, 0}
	maze[7][5] = point.Point{5, 7, 1}
	maze[7][6] = point.Point{6, 7, 0}
	maze[7][7] = point.Point{7, 7, 0}
	maze[7][8] = point.Point{8, 7, 0}
	maze[7][9] = point.Point{9, 7, 1}
	maze[8] = make([]point.Point, 10)
	maze[8][0] = point.Point{0, 8, 0}
	maze[8][1] = point.Point{1, 8, 1}
	maze[8][2] = point.Point{2, 8, 1}
	maze[8][3] = point.Point{3, 8, 1}
	maze[8][4] = point.Point{4, 8, 1}
	maze[8][5] = point.Point{5, 8, 1}
	maze[8][6] = point.Point{6, 8, 1}
	maze[8][7] = point.Point{7, 8, 1}
	maze[8][8] = point.Point{8, 8, 1}
	maze[8][9] = point.Point{9, 8, 1}
	maze[9] = make([]point.Point, 10)
	maze[9][0] = point.Point{0, 9, 0}
	maze[9][1] = point.Point{1, 9, 0}
	maze[9][2] = point.Point{2, 9, 0}
	maze[9][3] = point.Point{3, 9, 0}
	maze[9][4] = point.Point{4, 9, 0}
	maze[9][5] = point.Point{5, 9, 0}
	maze[9][6] = point.Point{6, 9, 0}
	maze[9][7] = point.Point{7, 9, 0}
	maze[9][8] = point.Point{8, 9, 3}
	maze[9][9] = point.Point{9, 9, 1}
}

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("GO Traveler")
	window.Connect("destroy", gtk.MainQuit)

	// Labirynt
	size := uint(len(maze))
	table, err := generateTable(size)
	if err != nil {
		panic(err)
	}
	// Start button
	startButton := gtk.NewToggleButtonWithLabel("Start")

	mainHBox := gtk.NewHBox(false, 1)
	leftVBox := gtk.NewVBox(false, 1)
	rightVBox := gtk.NewVBox(false, 1)

	leftVBox.Add(table)
	rightVBox.Add(startButton)
	mainHBox.Add(leftVBox)
	mainHBox.Add(rightVBox)

	window.Add(mainHBox)
	window.SetResizable(false)
	window.SetSizeRequest(600, 500)
	window.ShowAll()

	gtk.Main()
}

func generateTable(size uint) (*gtk.Table, error) {
	table := gtk.NewTable(size, size, false)
	table.SetColSpacings(0)
	table.SetRowSpacings(0)

	if uint(len(maze)) != size || uint(len(maze[0])) != size {
		return nil, fmt.Errorf("Nieprawidłowy rozmiar tablicy!")
	}

	for y := uint(0); y < size; y++ {
		for x := uint(0); x < size; x++ {
			switch maze[y][x].Val {
			case 0:
				table.Attach(gtk.NewImage(), x, x+1, y, y+1, gtk.FILL, gtk.FILL, 0, 0)
			case 1:
				table.Attach(gtk.NewImageFromFile("./images/wall.png"), x, x+1, y, y+1, gtk.FILL, gtk.FILL, 0, 0)
			case 2:
				table.Attach(gtk.NewImageFromFile("./images/start.png"), x, x+1, y, y+1, gtk.FILL, gtk.FILL, 0, 0)
			case 3:
				table.Attach(gtk.NewImageFromFile("./images/stop.png"), x, x+1, y, y+1, gtk.FILL, gtk.FILL, 0, 0)
			}
		}
	}

	return table, nil
}
