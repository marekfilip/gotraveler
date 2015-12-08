package maze

import "gotraveler/point"

func GetDefaultMaze() Maze {
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

	var lab Maze = make(Maze, 10)

	lab[0] = make([]*point.Point, 10)

	lab[0][0] = point.New(0, 0, point.START)
	lab[0][1] = point.New(1, 0, point.WALL)
	lab[0][2] = point.New(2, 0, point.WAY)
	lab[0][3] = point.New(3, 0, point.WAY)
	lab[0][4] = point.New(4, 0, point.WAY)
	lab[0][5] = point.New(5, 0, point.WAY)
	lab[0][6] = point.New(6, 0, point.WAY)
	lab[0][7] = point.New(7, 0, point.WALL)
	lab[0][8] = point.New(8, 0, point.WALL)
	lab[0][9] = point.New(9, 0, point.WALL)
	lab[1] = make([]*point.Point, 10)

	lab[1][0] = point.New(0, 1, point.WAY)
	lab[1][1] = point.New(1, 1, point.WAY)
	lab[1][2] = point.New(2, 1, point.WAY)
	lab[1][3] = point.New(3, 1, point.WALL)
	lab[1][4] = point.New(4, 1, point.WALL)
	lab[1][5] = point.New(5, 1, point.WALL)
	lab[1][6] = point.New(6, 1, point.WAY)
	lab[1][7] = point.New(7, 1, point.WAY)
	lab[1][8] = point.New(8, 1, point.WALL)
	lab[1][9] = point.New(9, 1, point.WALL)
	lab[2] = make([]*point.Point, 10)

	lab[2][0] = point.New(0, 2, point.WALL)
	lab[2][1] = point.New(1, 2, point.WALL)
	lab[2][2] = point.New(2, 2, point.WAY)
	lab[2][3] = point.New(3, 2, point.WAY)
	lab[2][4] = point.New(4, 2, point.WAY)
	lab[2][5] = point.New(5, 2, point.WALL)
	lab[2][6] = point.New(6, 2, point.WALL)
	lab[2][7] = point.New(7, 2, point.WAY)
	lab[2][8] = point.New(8, 2, point.WAY)
	lab[2][9] = point.New(9, 2, point.WAY)
	lab[3] = make([]*point.Point, 10)

	lab[3][0] = point.New(0, 3, point.WAY)
	lab[3][1] = point.New(1, 3, point.WAY)
	lab[3][2] = point.New(2, 3, point.WAY)
	lab[3][3] = point.New(3, 3, point.WALL)
	lab[3][4] = point.New(4, 3, point.WAY)
	lab[3][5] = point.New(5, 3, point.WAY)
	lab[3][6] = point.New(6, 3, point.WALL)
	lab[3][7] = point.New(7, 3, point.WALL)
	lab[3][8] = point.New(8, 3, point.WALL)
	lab[3][9] = point.New(9, 3, point.WALL)
	lab[4] = make([]*point.Point, 10)

	lab[4][0] = point.New(0, 4, point.WALL)
	lab[4][1] = point.New(1, 4, point.WALL)
	lab[4][2] = point.New(2, 4, point.WALL)
	lab[4][3] = point.New(3, 4, point.WALL)
	lab[4][4] = point.New(4, 4, point.WALL)
	lab[4][5] = point.New(5, 4, point.WAY)
	lab[4][6] = point.New(6, 4, point.WAY)
	lab[4][7] = point.New(7, 4, point.WAY)
	lab[4][8] = point.New(8, 4, point.WAY)
	lab[4][9] = point.New(9, 4, point.WAY)
	lab[5] = make([]*point.Point, 10)

	lab[5][0] = point.New(0, 5, point.WAY)
	lab[5][1] = point.New(1, 5, point.WAY)
	lab[5][2] = point.New(2, 5, point.WAY)
	lab[5][3] = point.New(3, 5, point.WALL)
	lab[5][4] = point.New(4, 5, point.WALL)
	lab[5][5] = point.New(5, 5, point.WALL)
	lab[5][6] = point.New(6, 5, point.WALL)
	lab[5][7] = point.New(7, 5, point.WALL)
	lab[5][8] = point.New(8, 5, point.WAY)
	lab[5][9] = point.New(9, 5, point.WALL)
	lab[6] = make([]*point.Point, 10)

	lab[6][0] = point.New(0, 6, point.WAY)
	lab[6][1] = point.New(1, 6, point.WALL)
	lab[6][2] = point.New(2, 6, point.WAY)
	lab[6][3] = point.New(3, 6, point.WALL)
	lab[6][4] = point.New(4, 6, point.WAY)
	lab[6][5] = point.New(5, 6, point.WAY)
	lab[6][6] = point.New(6, 6, point.WAY)
	lab[6][7] = point.New(7, 6, point.WALL)
	lab[6][8] = point.New(8, 6, point.WAY)
	lab[6][9] = point.New(9, 6, point.WALL)
	lab[7] = make([]*point.Point, 10)

	lab[7][0] = point.New(0, 7, point.WAY)
	lab[7][1] = point.New(1, 7, point.WALL)
	lab[7][2] = point.New(2, 7, point.WAY)
	lab[7][3] = point.New(3, 7, point.WAY)
	lab[7][4] = point.New(4, 7, point.WAY)
	lab[7][5] = point.New(5, 7, point.WALL)
	lab[7][6] = point.New(6, 7, point.WAY)
	lab[7][7] = point.New(7, 7, point.WAY)
	lab[7][8] = point.New(8, 7, point.WAY)
	lab[7][9] = point.New(9, 7, point.WALL)
	lab[8] = make([]*point.Point, 10)

	lab[8][0] = point.New(0, 8, point.WAY)
	lab[8][1] = point.New(1, 8, point.WALL)
	lab[8][2] = point.New(2, 8, point.WALL)
	lab[8][3] = point.New(3, 8, point.WALL)
	lab[8][4] = point.New(4, 8, point.WALL)
	lab[8][5] = point.New(5, 8, point.WALL)
	lab[8][6] = point.New(6, 8, point.WALL)
	lab[8][7] = point.New(7, 8, point.WALL)
	lab[8][8] = point.New(8, 8, point.WALL)
	lab[8][9] = point.New(9, 8, point.WALL)
	lab[9] = make([]*point.Point, 10)

	lab[9][0] = point.New(0, 9, point.WAY)
	lab[9][1] = point.New(1, 9, point.WAY)
	lab[9][2] = point.New(2, 9, point.WAY)
	lab[9][3] = point.New(3, 9, point.WAY)
	lab[9][4] = point.New(4, 9, point.WAY)
	lab[9][5] = point.New(5, 9, point.WAY)
	lab[9][6] = point.New(6, 9, point.WAY)
	lab[9][7] = point.New(7, 9, point.WAY)
	lab[9][8] = point.New(8, 9, point.FINISH)
	lab[9][9] = point.New(9, 9, point.WALL)

	return lab
}
