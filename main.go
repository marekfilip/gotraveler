package main

import (
	"github.com/mattn/go-gtk/gtk"
	"gotraveler/maze"
	"io/ioutil"
	"os"
)

func main() {
	gtk.Init(&os.Args)

	var (
		lab       maze.Maze   = maze.GetDefaultMaze()
		window    *gtk.Window = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
		mainHBox  *gtk.HBox   = gtk.NewHBox(false, 1)
		table     *gtk.Table  = lab.GenerateTable()
		leftVBox  *gtk.VBox   = gtk.NewVBox(true, 0)
		rightVBox *gtk.VBox   = gtk.NewVBox(true, 0)
	)

	window.SetTitle("GO Traveler")
	window.Connect("destroy", gtk.MainQuit)
	window.SetResizable(false)
	window.SetSizeRequest(lab.GetMaxPointsY()*25+120, lab.GetMaxPointsY()*25)
	leftVBox.SetBorderWidth(5)
	rightVBox.SetBorderWidth(5)

	window.Add(mainHBox)
	mainHBox.Add(leftVBox)
	mainHBox.Add(rightVBox)
	leftVBox.Add(table)

	// Start button
	buttonbox := gtk.NewVBox(false, 1)
	startButton := gtk.NewToggleButtonWithLabel("Start")
	go startButton.Clicked(func() {
		go lab.Solve(table, maze.DFS_METHOD)
	})
	buttonbox.Add(startButton)

	// Ładowanie button
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

			table = templab.GenerateTable()

			window.SetSizeRequest(templab.GetMaxPointsY()*25+120, templab.GetMaxPointsX()*25)
			leftVBox.Add(table)
			leftVBox.ShowAll()
			lab = templab

			filechooserdialog.Destroy()
		})
		filechooserdialog.Run()
	})
	buttonbox.Add(dest)
	rightVBox.Add(buttonbox)

	// Radio
	radiobuttonbox := gtk.NewVBox(false, 1)
	radiofirst := gtk.NewRadioButtonWithLabel(nil, "DFS")
	radiobuttonbox.Add(radiofirst)
	radiobuttonbox.Add(gtk.NewRadioButtonWithLabel(radiofirst.GetGroup(), "BFS"))
	rightVBox.Add(radiobuttonbox)
	//radiobutton.SetMode(false);
	radiofirst.SetActive(true)

	window.ShowAll()
	gtk.Main()
}
