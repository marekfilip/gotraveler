package maze

import (
	"encoding/json"
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"gotraveler/point"
	"gotraveler/queue"
	//"time"
)

const (
	// Depth-first search, przeszukiwanie wgłąb
	DFS_METHOD int = 0
	// Breadth-first search, przeszukiwanie wszerz
	BFS_METHOD int = 1
)

type Maze [][]*point.Point

func (m *Maze) GetMaxPointsY() int { return len(*m) }

func (m *Maze) GetMaxPointsX() int {
	var max int = 0

	for y := 0; y < len(*m); y++ {
		if len((*m)[y]) > max {
			max = len((*m)[y])
		}
	}

	return max
}

func (m *Maze) Solve(tab *gtk.Table, searchMethod int) error {
	if searchMethod != DFS_METHOD && searchMethod != BFS_METHOD {
		return fmt.Errorf("Wybrano niepoprawną metodę przeszukiwania")
	}

	var (
		que          queue.Queue
		currentPoint *point.Point = nil
		foundFinish  bool
		sp           *point.Point = m.GetStartPoint()
	)

	if sp != nil {
		que.Push(sp)

		// Przeszukiwanie wszerz (nowe na końcu)
		for !que.IsEmpty() {
			currentPoint = que.Pop()

			if currentPoint.IsFinish() {
				fmt.Println("Znalazłem koniec")
				foundFinish = true
				m.ColorWays(tab)
			} else {
				currentPoint.Visited = true
				// Szukamy punktów obok
				// W lewo
				if int(currentPoint.X)-1 >= 0 && !(*m)[currentPoint.Y][currentPoint.X-1].IsWall() && !(*m)[currentPoint.Y][currentPoint.X-1].Visited {
					(*m)[currentPoint.Y][currentPoint.X-1].Parent = currentPoint
					if searchMethod == DFS_METHOD {
						que.Push((*m)[currentPoint.Y][currentPoint.X-1])
					}
					if searchMethod == BFS_METHOD {
						que.Unshift((*m)[currentPoint.Y][currentPoint.X-1])
					}
				}
				// W górę
				if int(currentPoint.Y)+1 < len((*m)[currentPoint.Y]) && !(*m)[currentPoint.Y+1][currentPoint.X].IsWall() && !(*m)[currentPoint.Y+1][currentPoint.X].Visited {
					(*m)[currentPoint.Y+1][currentPoint.X].Parent = currentPoint
					if searchMethod == DFS_METHOD {
						que.Push((*m)[currentPoint.Y+1][currentPoint.X])
					}
					if searchMethod == BFS_METHOD {
						que.Unshift((*m)[currentPoint.Y+1][currentPoint.X])
					}
				}
				// W prawo
				if int(currentPoint.X)+1 < len((*m)[currentPoint.Y]) && !(*m)[currentPoint.Y][currentPoint.X+1].IsWall() && !(*m)[currentPoint.Y][currentPoint.X+1].Visited {
					(*m)[currentPoint.Y][currentPoint.X+1].Parent = currentPoint
					if searchMethod == DFS_METHOD {
						que.Push((*m)[currentPoint.Y][currentPoint.X+1])
					}
					if searchMethod == BFS_METHOD {
						que.Unshift((*m)[currentPoint.Y][currentPoint.X+1])
					}
				}
				// W dół
				if int(currentPoint.Y)-1 >= 0 && !(*m)[currentPoint.Y-1][currentPoint.X].IsWall() && !(*m)[currentPoint.Y-1][currentPoint.X].Visited {
					(*m)[currentPoint.Y-1][currentPoint.X].Parent = currentPoint
					if searchMethod == DFS_METHOD {
						que.Push((*m)[currentPoint.Y-1][currentPoint.X])
					}
					if searchMethod == BFS_METHOD {
						que.Unshift((*m)[currentPoint.Y-1][currentPoint.X])
					}
				}
			}
			currentPoint.ColorVisitedPoint(tab)
			//time.Sleep(time.Millisecond * 75)
		}

		if foundFinish == false {
			fmt.Println("Kolejka pusta. Koniec")
		}
	} else {
		fmt.Println("Brak punktu startowego. Koniec")
	}

	return nil
}

func (m *Maze) GenerateTable() *gtk.Table {
	var table *gtk.Table = gtk.NewTable(uint(len(*m)), uint(len(*m)), false)

	table.SetColSpacings(0)
	table.SetRowSpacings(0)

	for y := uint(0); y < uint(len(*m)); y++ {
		for x := uint(0); x < uint(len((*m)[y])); x++ {
			switch (*m)[y][x].Val {
			case point.WALL:
				(*m)[y][x].Img = gtk.NewImageFromFile("./images/wall.png")
				table.Attach((*m)[y][x].Img, x, x+1, y, y+1, gtk.FILL, gtk.FILL, 0, 0)

			case point.START:
				(*m)[y][x].Img = gtk.NewImageFromFile("./images/start.png")
				table.Attach((*m)[y][x].Img, x, x+1, y, y+1, gtk.FILL, gtk.FILL, 0, 0)

			case point.FINISH:
				(*m)[y][x].Img = gtk.NewImageFromFile("./images/stop.png")
				table.Attach((*m)[y][x].Img, x, x+1, y, y+1, gtk.FILL, gtk.FILL, 0, 0)
			}
		}
	}

	return table
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

func (m *Maze) GetStartPoint() *point.Point {
	for y := uint(0); y < uint(len(*m)); y++ {
		for x := uint(0); x < uint(len((*m)[y])); x++ {
			if (*m)[y][x].Val == point.START {
				return (*m)[y][x]
			}
		}
	}

	return nil
}

func (m *Maze) GetFinishPoint() *point.Point {
	for y := uint(0); y < uint(len(*m)); y++ {
		for x := uint(0); x < uint(len((*m)[y])); x++ {
			if (*m)[y][x].Val == point.FINISH {
				return (*m)[y][x]
			}
		}
	}

	return nil
}

func (m *Maze) ColorWays(tab *gtk.Table) {
	var fp *point.Point = m.GetFinishPoint()

	if fp != nil {
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
}
