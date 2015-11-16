package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"gotraveler/maze"
	"gotraveler/point"
	"gotraveler/queue"
	"io/ioutil"
	"os"
	"time"
)

var lab maze.Maze
var startPoint *point.Point
var table *gtk.Table
var size int

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
	lab = make([][]*point.Point, 10)
	lab[0] = make([]*point.Point, 10)
	lab[0][0] = point.New(0, 0, 2)
	lab[0][1] = point.New(1, 0, 1)
	lab[0][2] = point.New(2, 0, 0)
	lab[0][3] = point.New(3, 0, 0)
	lab[0][4] = point.New(4, 0, 0)
	lab[0][5] = point.New(5, 0, 0)
	lab[0][6] = point.New(6, 0, 0)
	lab[0][7] = point.New(7, 0, 1)
	lab[0][8] = point.New(8, 0, 1)
	lab[0][9] = point.New(9, 0, 1)
	lab[1] = make([]*point.Point, 10)
	lab[1][0] = point.New(0, 1, 0)
	lab[1][1] = point.New(1, 1, 0)
	lab[1][2] = point.New(2, 1, 0)
	lab[1][3] = point.New(3, 1, 1)
	lab[1][4] = point.New(4, 1, 1)
	lab[1][5] = point.New(5, 1, 1)
	lab[1][6] = point.New(6, 1, 0)
	lab[1][7] = point.New(7, 1, 0)
	lab[1][8] = point.New(8, 1, 1)
	lab[1][9] = point.New(9, 1, 1)
	lab[2] = make([]*point.Point, 10)
	lab[2][0] = point.New(0, 2, 1)
	lab[2][1] = point.New(1, 2, 1)
	lab[2][2] = point.New(2, 2, 0)
	lab[2][3] = point.New(3, 2, 0)
	lab[2][4] = point.New(4, 2, 0)
	lab[2][5] = point.New(5, 2, 1)
	lab[2][6] = point.New(6, 2, 1)
	lab[2][7] = point.New(7, 2, 0)
	lab[2][8] = point.New(8, 2, 0)
	lab[2][9] = point.New(9, 2, 0)
	lab[3] = make([]*point.Point, 10)
	lab[3][0] = point.New(0, 3, 0)
	lab[3][1] = point.New(1, 3, 0)
	lab[3][2] = point.New(2, 3, 0)
	lab[3][3] = point.New(3, 3, 1)
	lab[3][4] = point.New(4, 3, 0)
	lab[3][5] = point.New(5, 3, 0)
	lab[3][6] = point.New(6, 3, 1)
	lab[3][7] = point.New(7, 3, 1)
	lab[3][8] = point.New(8, 3, 1)
	lab[3][9] = point.New(9, 3, 1)
	lab[4] = make([]*point.Point, 10)
	lab[4][0] = point.New(0, 4, 1)
	lab[4][1] = point.New(1, 4, 1)
	lab[4][2] = point.New(2, 4, 1)
	lab[4][3] = point.New(3, 4, 1)
	lab[4][4] = point.New(4, 4, 1)
	lab[4][5] = point.New(5, 4, 0)
	lab[4][6] = point.New(6, 4, 0)
	lab[4][7] = point.New(7, 4, 0)
	lab[4][8] = point.New(8, 4, 0)
	lab[4][9] = point.New(9, 4, 0)
	lab[5] = make([]*point.Point, 10)
	lab[5][0] = point.New(0, 5, 0)
	lab[5][1] = point.New(1, 5, 0)
	lab[5][2] = point.New(2, 5, 0)
	lab[5][3] = point.New(3, 5, 1)
	lab[5][4] = point.New(4, 5, 1)
	lab[5][5] = point.New(5, 5, 1)
	lab[5][6] = point.New(6, 5, 1)
	lab[5][7] = point.New(7, 5, 1)
	lab[5][8] = point.New(8, 5, 0)
	lab[5][9] = point.New(9, 5, 1)
	lab[6] = make([]*point.Point, 10)
	lab[6][0] = point.New(0, 6, 0)
	lab[6][1] = point.New(1, 6, 1)
	lab[6][2] = point.New(2, 6, 0)
	lab[6][3] = point.New(3, 6, 1)
	lab[6][4] = point.New(4, 6, 0)
	lab[6][5] = point.New(5, 6, 0)
	lab[6][6] = point.New(6, 6, 0)
	lab[6][7] = point.New(7, 6, 1)
	lab[6][8] = point.New(8, 6, 0)
	lab[6][9] = point.New(9, 6, 1)
	lab[7] = make([]*point.Point, 10)
	lab[7][0] = point.New(0, 7, 0)
	lab[7][1] = point.New(1, 7, 1)
	lab[7][2] = point.New(2, 7, 0)
	lab[7][3] = point.New(3, 7, 0)
	lab[7][4] = point.New(4, 7, 0)
	lab[7][5] = point.New(5, 7, 1)
	lab[7][6] = point.New(6, 7, 0)
	lab[7][7] = point.New(7, 7, 0)
	lab[7][8] = point.New(8, 7, 0)
	lab[7][9] = point.New(9, 7, 1)
	lab[8] = make([]*point.Point, 10)
	lab[8][0] = point.New(0, 8, 0)
	lab[8][1] = point.New(1, 8, 1)
	lab[8][2] = point.New(2, 8, 1)
	lab[8][3] = point.New(3, 8, 1)
	lab[8][4] = point.New(4, 8, 1)
	lab[8][5] = point.New(5, 8, 1)
	lab[8][6] = point.New(6, 8, 1)
	lab[8][7] = point.New(7, 8, 1)
	lab[8][8] = point.New(8, 8, 1)
	lab[8][9] = point.New(9, 8, 1)
	lab[9] = make([]*point.Point, 10)
	lab[9][0] = point.New(0, 9, 0)
	lab[9][1] = point.New(1, 9, 0)
	lab[9][2] = point.New(2, 9, 0)
	lab[9][3] = point.New(3, 9, 0)
	lab[9][4] = point.New(4, 9, 0)
	lab[9][5] = point.New(5, 9, 0)
	lab[9][6] = point.New(6, 9, 0)
	lab[9][7] = point.New(7, 9, 0)
	lab[9][8] = point.New(8, 9, 3)
	lab[9][9] = point.New(9, 9, 1)
}

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("GO Traveler")
	window.Connect("destroy", gtk.MainQuit)
	leftVBox := gtk.NewVBox(true, 0)
	rightVBox := gtk.NewVBox(true, 0)

	// Labirynt
	table, startPoint, size, err := lab.GenerateTable()
	if err != nil {
		panic(err)
	}
	// Start button
	startButton := gtk.NewToggleButtonWithLabel("Start")
	go startButton.Clicked(func() {
		go findWay(table, startPoint, size)
	})

	// Ładowanie lab
	dest := gtk.NewButtonWithLabel("Załaduj labirynt")
	dest.Clicked(func() {
		filechooserdialog := gtk.NewFileChooserDialog(
			"Wybierz plik...",
			dest.GetTopLevelAsWindow(),
			gtk.FILE_CHOOSER_ACTION_OPEN,
			gtk.STOCK_OK,
			gtk.RESPONSE_ACCEPT)
		filter := gtk.NewFileFilter()
		filter.AddPattern("*.json")
		filechooserdialog.AddFilter(filter)
		filechooserdialog.Response(func() {
			var templab maze.Maze
			leftVBox.Remove(table)

			dat, err := ioutil.ReadFile(filechooserdialog.GetFilename())
			if err != nil {
				panic(err)
			}
			err = templab.UnmarshalJSON(dat)
			if err != nil {
				panic(err)
			}

			table, startPoint, size, err = templab.GenerateTable()
			if err != nil {
				panic(err)
			}
			window.SetSizeRequest(size*50+100, size*50)
			leftVBox.Add(table)
			leftVBox.ShowAll()
			lab = templab

			filechooserdialog.Destroy()
		})
		filechooserdialog.Run()
	})

	leftVBox.SetBorderWidth(5)
	leftVBox.Add(table)

	rightVBox.SetBorderWidth(5)
	rightVBox.Add(startButton)
	rightVBox.Add(dest)

	mainHBox := gtk.NewHBox(false, 1)
	mainHBox.Add(leftVBox)
	mainHBox.Add(rightVBox)

	window.Add(mainHBox)
	window.SetResizable(false)
	window.SetSizeRequest(size*50+120, size*50)
	window.ShowAll()

	gtk.Main()
}

func findWay(tab *gtk.Table, sp *point.Point, size int) {
	var que queue.Queue
	que.Push(sp)
	var currentPoint *point.Point = nil
	var foundFinish bool

	// Przeszukiwanie wszerz (nowe na końcu)
	for !que.IsEmpty() {
		currentPoint = que.Shift()

		if currentPoint.IsFinish() {
			fmt.Println("Znalazłem koniec")
			foundFinish = true
			colorWays(currentPoint, tab)
		} else {
			currentPoint.Visited = true
			// Szukamy punktów obok
			if int(currentPoint.Y)-1 >= 0 && !lab[currentPoint.Y-1][currentPoint.X].IsWall() && !lab[currentPoint.Y-1][currentPoint.X].Visited {
				lab[currentPoint.Y-1][currentPoint.X].Parent = currentPoint
				que.Push(lab[currentPoint.Y-1][currentPoint.X]) // W dół
			}
			if int(currentPoint.Y)+1 < size && !lab[currentPoint.Y+1][currentPoint.X].IsWall() && !lab[currentPoint.Y+1][currentPoint.X].Visited {
				lab[currentPoint.Y+1][currentPoint.X].Parent = currentPoint
				que.Push(lab[currentPoint.Y+1][currentPoint.X]) // W górę
			}
			if int(currentPoint.X)-1 >= 0 && !lab[currentPoint.Y][currentPoint.X-1].IsWall() && !lab[currentPoint.Y][currentPoint.X-1].Visited {
				lab[currentPoint.Y][currentPoint.X-1].Parent = currentPoint
				que.Push(lab[currentPoint.Y][currentPoint.X-1]) // W lewo
			}
			if int(currentPoint.X)+1 < size && !lab[currentPoint.Y][currentPoint.X+1].IsWall() && !lab[currentPoint.Y][currentPoint.X+1].Visited {
				lab[currentPoint.Y][currentPoint.X+1].Parent = currentPoint
				que.Push(lab[currentPoint.Y][currentPoint.X+1]) // W prawo
			}
		}
		go colorVisitedPoint(tab, currentPoint)
		time.Sleep(time.Millisecond * 50)
	}

	if foundFinish == false {
		fmt.Println("Kolejka pusta. Koniec")
	}
}

func colorWays(fp *point.Point, tab *gtk.Table) {
	var tp *point.Point = fp.Parent

	for tp != nil {
		if tp.IsStart() || tp.IsFinish() {
			continue
		}
		tab.Remove(tp.Img)
		tp.Img = gtk.NewImageFromFile("./images/greendot.png")
		tab.Attach(tp.Img, tp.X, tp.X+1, tp.Y, tp.Y+1, gtk.FILL, gtk.FILL, 0, 0)
		tp.Img.Show()
		tp = tp.Parent
	}
}

func colorVisitedPoint(tab *gtk.Table, p *point.Point) {
	p.Img = gtk.NewImageFromFile("./images/reddot.png")
	tab.Attach(p.Img, p.X, p.X+1, p.Y, p.Y+1, gtk.FILL, gtk.FILL, 0, 0)
	p.Img.Show()
}
