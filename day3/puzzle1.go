package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"math"
)

func readFile(filePath string) ([]string, []string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
	}
	r := bufio.NewReader(file)

	line1, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
	}
	line2, err2 := r.ReadString('\n')
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	line1 = strings.TrimSuffix(line1, "\n")
	line2 = strings.TrimSuffix(line2, "\n")
	arrayLine1 := strings.Split(line1, ",")
	arrayLine2 := strings.Split(line2, ",")
	return arrayLine1, arrayLine2
}

func createGrid(size int) ([][]int, int) {
	a := make([][]int, size)
	for i := range a {
	    a[i] = make([]int, size)
	}
	return a, size/2
}

func distance(s string) int {
    for i := range s {
        if i > 0 {
		   num, _ := strconv.Atoi(s[i:])
		   return num
 	   }
    }
    num, _ := strconv.Atoi(s[:0])
    return num
}

func direction(s string) int {
	if strings.Contains(s, "U") {
		return 1
	} else if strings.Contains(s, "D") {
		return 2
	} else if strings.Contains(s, "R") {
		return 3
	} else {
		return 4
	}
}

func drawPath(path []string, grid [][]int, start int, token int) [][]int {
	locX, locY := start, start
	for i:=0; i<len(path); i++ {
		locX, locY, grid = drawLine(direction(path[i]), distance(path[i]), token, locX, locY, grid)
	}
	return grid
}

func drawLine(direction int, length int, token int, locX int, locY int, grid [][]int) (int,int,[][]int) {
	if direction == 1 {
		for i:=length; i>0; i--{
			locY++
			if (grid[locX][locY] == 1) && token == 2 {
				grid[locX][locY] = 3
			} else {
				grid[locX][locY] = token
			}
		}
	} else if direction == 2 {
		for i:=length; i>0; i--{
			locY--
			if (grid[locX][locY] == 1) && token == 2 {
				grid[locX][locY] = 3
			} else {
				grid[locX][locY] = token
			}
		}
	} else if direction == 3 {
		for i:=length; i>0; i--{
			locX--
			if (grid[locX][locY] == 1) && token == 2 {
				grid[locX][locY] = 3
			} else {
				grid[locX][locY] = token
			}
		}
	} else {
		for i:=length; i>0; i--{
			locX++
			if (grid[locX][locY] == 1) && token == 2 {
				grid[locX][locY] = 3
			} else {
				grid[locX][locY] = token
			}
		}
	}
	return locX, locY, grid
}

func getDistance(grid [][]int, centre int, size int) {
	result := 10000
	for i:=0;i<size;i++{
		for j:=0;j<size;j++{
			if grid[i][j] == 3 {
				distance := int(math.Abs(float64(i)-float64(centre)) + math.Abs(float64(j)-float64(centre)))
				if distance < result {
					result = distance
				}

			}
		}
	}
	fmt.Println(result)
}

func main() {
	size := 17500
	line1, line2 := readFile("./input.txt")
	grid, centre := createGrid(size)
	grid = drawPath(line1, grid, centre, 1)
	grid = drawPath(line2, grid, centre, 2)
	getDistance(grid, centre, size)
}
