package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"gotraveler/point"
	"gotraveler/queue"
	"os"
	"strings"
	"time"
	"unsafe"
)

var maze [][]*point.Point

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
	maze = make([][]*point.Point, 10)
	maze[0] = make([]*point.Point, 10)
	maze[0][0] = point.New(0, 0, 2)
	maze[0][1] = point.New(1, 0, 1)
	maze[0][2] = point.New(2, 0, 0)
	maze[0][3] = point.New(3, 0, 0)
	maze[0][4] = point.New(4, 0, 0)
	maze[0][5] = point.New(5, 0, 0)
	maze[0][6] = point.New(6, 0, 0)
	maze[0][7] = point.New(7, 0, 1)
	maze[0][8] = point.New(8, 0, 1)
	maze[0][9] = point.New(9, 0, 1)
	maze[1] = make([]*point.Point, 10)
	maze[1][0] = point.New(0, 1, 0)
	maze[1][1] = point.New(1, 1, 0)
	maze[1][2] = point.New(2, 1, 0)
	maze[1][3] = point.New(3, 1, 1)
	maze[1][4] = point.New(4, 1, 1)
	maze[1][5] = point.New(5, 1, 1)
	maze[1][6] = point.New(6, 1, 0)
	maze[1][7] = point.New(7, 1, 0)
	maze[1][8] = point.New(8, 1, 1)
	maze[1][9] = point.New(9, 1, 1)
	maze[2] = make([]*point.Point, 10)
	maze[2][0] = point.New(0, 2, 1)
	maze[2][1] = point.New(1, 2, 1)
	maze[2][2] = point.New(2, 2, 0)
	maze[2][3] = point.New(3, 2, 0)
	maze[2][4] = point.New(4, 2, 0)
	maze[2][5] = point.New(5, 2, 1)
	maze[2][6] = point.New(6, 2, 1)
	maze[2][7] = point.New(7, 2, 0)
	maze[2][8] = point.New(8, 2, 0)
	maze[2][9] = point.New(9, 2, 0)
	maze[3] = make([]*point.Point, 10)
	maze[3][0] = point.New(0, 3, 0)
	maze[3][1] = point.New(1, 3, 0)
	maze[3][2] = point.New(2, 3, 0)
	maze[3][3] = point.New(3, 3, 1)
	maze[3][4] = point.New(4, 3, 0)
	maze[3][5] = point.New(5, 3, 0)
	maze[3][6] = point.New(6, 3, 1)
	maze[3][7] = point.New(7, 3, 1)
	maze[3][8] = point.New(8, 3, 1)
	maze[3][9] = point.New(9, 3, 1)
	maze[4] = make([]*point.Point, 10)
	maze[4][0] = point.New(0, 4, 1)
	maze[4][1] = point.New(1, 4, 1)
	maze[4][2] = point.New(2, 4, 1)
	maze[4][3] = point.New(3, 4, 1)
	maze[4][4] = point.New(4, 4, 1)
	maze[4][5] = point.New(5, 4, 0)
	maze[4][6] = point.New(6, 4, 0)
	maze[4][7] = point.New(7, 4, 0)
	maze[4][8] = point.New(8, 4, 0)
	maze[4][9] = point.New(9, 4, 0)
	maze[5] = make([]*point.Point, 10)
	maze[5][0] = point.New(0, 5, 0)
	maze[5][1] = point.New(1, 5, 0)
	maze[5][2] = point.New(2, 5, 0)
	maze[5][3] = point.New(3, 5, 1)
	maze[5][4] = point.New(4, 5, 1)
	maze[5][5] = point.New(5, 5, 1)
	maze[5][6] = point.New(6, 5, 1)
	maze[5][7] = point.New(7, 5, 1)
	maze[5][8] = point.New(8, 5, 0)
	maze[5][9] = point.New(9, 5, 1)
	maze[6] = make([]*point.Point, 10)
	maze[6][0] = point.New(0, 6, 0)
	maze[6][1] = point.New(1, 6, 1)
	maze[6][2] = point.New(2, 6, 0)
	maze[6][3] = point.New(3, 6, 1)
	maze[6][4] = point.New(4, 6, 0)
	maze[6][5] = point.New(5, 6, 0)
	maze[6][6] = point.New(6, 6, 0)
	maze[6][7] = point.New(7, 6, 1)
	maze[6][8] = point.New(8, 6, 0)
	maze[6][9] = point.New(9, 6, 1)
	maze[7] = make([]*point.Point, 10)
	maze[7][0] = point.New(0, 7, 0)
	maze[7][1] = point.New(1, 7, 1)
	maze[7][2] = point.New(2, 7, 0)
	maze[7][3] = point.New(3, 7, 0)
	maze[7][4] = point.New(4, 7, 0)
	maze[7][5] = point.New(5, 7, 1)
	maze[7][6] = point.New(6, 7, 0)
	maze[7][7] = point.New(7, 7, 0)
	maze[7][8] = point.New(8, 7, 0)
	maze[7][9] = point.New(9, 7, 1)
	maze[8] = make([]*point.Point, 10)
	maze[8][0] = point.New(0, 8, 0)
	maze[8][1] = point.New(1, 8, 1)
	maze[8][2] = point.New(2, 8, 1)
	maze[8][3] = point.New(3, 8, 1)
	maze[8][4] = point.New(4, 8, 1)
	maze[8][5] = point.New(5, 8, 1)
	maze[8][6] = point.New(6, 8, 1)
	maze[8][7] = point.New(7, 8, 1)
	maze[8][8] = point.New(8, 8, 1)
	maze[8][9] = point.New(9, 8, 1)
	maze[9] = make([]*point.Point, 10)
	maze[9][0] = point.New(0, 9, 0)
	maze[9][1] = point.New(1, 9, 0)
	maze[9][2] = point.New(2, 9, 0)
	maze[9][3] = point.New(3, 9, 0)
	maze[9][4] = point.New(4, 9, 0)
	maze[9][5] = point.New(5, 9, 0)
	maze[9][6] = point.New(6, 9, 0)
	maze[9][7] = point.New(7, 9, 0)
	maze[9][8] = point.New(8, 9, 3)
	maze[9][9] = point.New(9, 9, 1)
}

func main() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("GO Traveler")
	window.Connect("destroy", gtk.MainQuit)

	// Labirynt
	size := len(maze)
	var startPoint *point.Point
	table, startPoint, err := generateTable(uint(size))
	if err != nil {
		panic(err)
	}
	// Start button
	startButton := gtk.NewToggleButtonWithLabel("Start")
	go startButton.Clicked(func() {
		go findWay(table, startPoint, size)
	})

	// Dnd
	targets := []gtk.TargetEntry{
		{"text/uri-list", 0, 0},
		{"STRING", 0, 1},
		{"text/plain", 0, 2},
	}
	dest := gtk.NewLabel("D&D")
	dest.DragDestSet(
		gtk.DEST_DEFAULT_MOTION|
			gtk.DEST_DEFAULT_HIGHLIGHT|
			gtk.DEST_DEFAULT_DROP,
		targets,
		gdk.ACTION_COPY)
	dest.DragDestAddUriTargets()
	dest.Connect("drag-data-received", func(ctx *glib.CallbackContext) {
		sdata := gtk.NewSelectionDataFromNative(unsafe.Pointer(ctx.Args(3)))
		if sdata != nil {
			a := (*[2000]uint8)(sdata.GetData())
			files := strings.Split(string(a[0:sdata.GetLength()-1]), "\n")
			fmt.Println(files)
			if len(files) > 1 {
				dialog := gtk.NewMessageDialog(
					window,
					gtk.DIALOG_MODAL,
					gtk.MESSAGE_INFO,
					gtk.BUTTONS_OK,
					"Tylko 1 na raz")
				dialog.SetTitle("D&D")
				dialog.Response(func() {
					dialog.Destroy()
				})
				dialog.Run()
			} else {

			}
			/*for i := range files {
				filename, _, _ := glib.FilenameFromUri(files[i])
				files[i] = filename
			}*/
		}
	})

	leftVBox := gtk.NewVBox(true, 0)
	leftVBox.SetBorderWidth(5)
	leftVBox.Add(table)

	rightVBox := gtk.NewVBox(true, 0)
	rightVBox.SetBorderWidth(5)
	rightVBox.Add(startButton)
	rightVBox.Add(dest)

	mainHBox := gtk.NewHBox(false, 1)
	mainHBox.Add(leftVBox)
	mainHBox.Add(rightVBox)

	window.Add(mainHBox)
	window.SetResizable(false)
	window.SetSizeRequest(size*50+100, size*50)
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
			if int(currentPoint.Y)-1 >= 0 && !maze[currentPoint.Y-1][currentPoint.X].IsWall() && !maze[currentPoint.Y-1][currentPoint.X].Visited {
				maze[currentPoint.Y-1][currentPoint.X].Parent = currentPoint
				que.Push(maze[currentPoint.Y-1][currentPoint.X]) // W dół
			}
			if int(currentPoint.Y)+1 < size && !maze[currentPoint.Y+1][currentPoint.X].IsWall() && !maze[currentPoint.Y+1][currentPoint.X].Visited {
				maze[currentPoint.Y+1][currentPoint.X].Parent = currentPoint
				que.Push(maze[currentPoint.Y+1][currentPoint.X]) // W górę
			}
			if int(currentPoint.X)-1 >= 0 && !maze[currentPoint.Y][currentPoint.X-1].IsWall() && !maze[currentPoint.Y][currentPoint.X-1].Visited {
				maze[currentPoint.Y][currentPoint.X-1].Parent = currentPoint
				que.Push(maze[currentPoint.Y][currentPoint.X-1]) // W lewo
			}
			if int(currentPoint.X)+1 < size && !maze[currentPoint.Y][currentPoint.X+1].IsWall() && !maze[currentPoint.Y][currentPoint.X+1].Visited {
				maze[currentPoint.Y][currentPoint.X+1].Parent = currentPoint
				que.Push(maze[currentPoint.Y][currentPoint.X+1]) // W prawo
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

func generateTable(size uint) (*gtk.Table, *point.Point, error) {
	var table *gtk.Table = gtk.NewTable(size, size, false)
	var startPoint *point.Point
	table.SetColSpacings(0)
	table.SetRowSpacings(0)

	if uint(len(maze)) != size || uint(len(maze[0])) != size {
		return nil, nil, fmt.Errorf("Nieprawidłowy rozmiar tablicy!")
	}

	for y := uint(0); y < size; y++ {
		for x := uint(0); x < size; x++ {
			switch maze[y][x].Val {
			case 1:
				maze[y][x].Img = gtk.NewImageFromFile("./images/wall.png")
				table.Attach(maze[y][x].Img, x, x+1, y, y+1, gtk.FILL, gtk.FILL, 0, 0)
			case 2:
				maze[y][x].Img = gtk.NewImageFromFile("./images/start.png")
				startPoint = maze[y][x]
				table.Attach(maze[y][x].Img, x, x+1, y, y+1, gtk.FILL, gtk.FILL, 0, 0)
			case 3:
				maze[y][x].Img = gtk.NewImageFromFile("./images/stop.png")
				table.Attach(maze[y][x].Img, x, x+1, y, y+1, gtk.FILL, gtk.FILL, 0, 0)
			}
		}
	}

	return table, startPoint, nil
}
