package maze

import (
	"encoding/json"
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"gotraveler/point"
)

type Maze [][]*point.Point

func (m *Maze) GenerateTable() (*gtk.Table, *point.Point, int, error) {
	var size uint = uint(len(*m))
	var table *gtk.Table = gtk.NewTable(size, size, false)
	var startPoint *point.Point
	table.SetColSpacings(0)
	table.SetRowSpacings(0)

	if uint(len(*m)) != size || uint(len((*m)[0])) != size {
		return nil, nil, 0, fmt.Errorf("Nieprawidłowy rozmiar tablicy!")
	}

	for y := uint(0); y < size; y++ {
		for x := uint(0); x < size; x++ {
			switch (*m)[y][x].Val {
			case 1:
				(*m)[y][x].Img = gtk.NewImageFromFile("./images/wall.png")
				table.Attach((*m)[y][x].Img, x, x+1, y, y+1, gtk.FILL, gtk.FILL, 0, 0)
			case 2:
				(*m)[y][x].Img = gtk.NewImageFromFile("./images/start.png")
				startPoint = (*m)[y][x]
				table.Attach((*m)[y][x].Img, x, x+1, y, y+1, gtk.FILL, gtk.FILL, 0, 0)
			case 3:
				(*m)[y][x].Img = gtk.NewImageFromFile("./images/stop.png")
				table.Attach((*m)[y][x].Img, x, x+1, y, y+1, gtk.FILL, gtk.FILL, 0, 0)
			}
		}
	}

	return table, startPoint, int(size), nil
}

func (m *Maze) UnmarshalJSON(data []byte) error {
	tab := &[][]int{}

	if err := json.Unmarshal(data, tab); err != nil {
		return err
	}

	if *m == nil {
		*m = make([][]*point.Point, len(*tab))
	}

	for y := 0; y < len(*tab); y++ {

		for x := 0; x < len((*tab)[y]); x++ {
			(*m)[y] = append((*m)[y], point.New(uint(x), uint(y), uint((*tab)[y][x])))
		}
	}

	return nil
}
