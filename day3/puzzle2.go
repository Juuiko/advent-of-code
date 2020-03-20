package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
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

func drawPath(path []string, grid [][]int, start int, token int, path2 []string) [][]int {
	locX, locY := start, start
	for i:=0; i<len(path); i++ {
		locX, locY, grid = drawLine(direction(path[i]), distance(path[i]), token, locX, locY, grid, path2)
	}
	return grid
}

func drawLine(direction int, length int, token int, locX int, locY int, grid [][]int, line1 []string) (int,int,[][]int) {
	if direction == 1 {
		for i:=length; i>0; i--{
			locY++
			if (grid[locX][locY] == 1) && token == 2 {
				grid[locX][locY] = 3
			} else {
				grid[locX][locY] = token
			}
		}
		return locX, locY, grid
	} else if direction == 2 {
		for i:=length; i>0; i--{
			locY--
			if (grid[locX][locY] == 1) && token == 2 {
				grid[locX][locY] = 3
			} else {
				grid[locX][locY] = token
			}
		}
		return locX, locY, grid
	} else if direction == 3 {
		for i:=length; i>0; i--{
			locX--
			if (grid[locX][locY] == 1) && token == 2 {
				grid[locX][locY] = 3
			} else {
				grid[locX][locY] = token
			}
		}
		return locX, locY, grid
	} else {
		for i:=length; i>0; i--{
			locX++
			if (grid[locX][locY] == 1) && token == 2 {
				grid[locX][locY] = 3
			} else {
				grid[locX][locY] = token
			}
		}
		return locX, locY, grid
	}
}

func traceCross(line2 []string, locXF int, locYF int, distanceL1 int, start int) {
     distanceL2 := 0
     locX, locY := start, start
     for i:=0; i<len(line2); i++ {
          direction := direction(line2[i])
          length := distance(line2[i])
          if direction == 1 {
               for i:=length; i>0; i--{
                    distanceL2++
                    locY++
                    if locX == locXF && locY == locYF {
                         fmt.Println(distanceL1+distanceL2)
                         return
                    }
               }
          } else if direction == 2 {
               for i:=length; i>0; i--{
                    distanceL2++
                    locY--
                    if locX == locXF && locY == locYF {
                         fmt.Println(distanceL1+distanceL2)
                         return
                    }
               }
          } else if direction == 3 {
               for i:=length; i>0; i--{
                    distanceL2++
                    locX--
                    if locX == locXF && locY == locYF {
                         fmt.Println(distanceL1+distanceL2)
                         return
                    }
               }
          } else {
               for i:=length; i>0; i--{
                    distanceL2++
                    locX++
                    if locX == locXF && locY == locYF {
                         fmt.Println(distanceL1+distanceL2)
                         return
                    }
               }
          }
     }
}

func traceLine(direction int, length int, locX int, locY int, start int, grid [][]int, line2 []string, distanceL1 int) (int,int,int) {
     if direction == 1 {
          for i:=length; i>0; i--{
               distanceL1++
               locY++
               if (grid[locX][locY] == 3) {
                    traceCross(line2, locX, locY, distanceL1, start)
               }
          }
     } else if direction == 2 {
          for i:=length; i>0; i--{
               distanceL1++
               locY--
               if (grid[locX][locY] == 3) {
                    traceCross(line2, locX, locY, distanceL1, start)
               }
          }
     } else if direction == 3 {
          for i:=length; i>0; i--{
               distanceL1++
               locX--
               if (grid[locX][locY] == 3) {
                    traceCross(line2, locX, locY, distanceL1, start)
               }
          }
     } else {
          for i:=length; i>0; i--{
               distanceL1++
               locX++
               if (grid[locX][locY] == 3) {
                    traceCross(line2, locX, locY, distanceL1, start)
               }
          }
     }
     return locX, locY, distanceL1
}

func findDistance(line1 []string, line2 []string, grid [][]int, start int) {
     locX, locY := start, start
     distanceL1 := 0
     for i:=0; i<len(line1); i++ {
          locX, locY, distanceL1 = traceLine(direction(line1[i]), distance(line1[i]), locX, locY, start, grid, line2, distanceL1)
     }
}

func main() {
	size := 17500
	line1, line2 := readFile("./input.txt")
	grid, centre := createGrid(size)
	grid = drawPath(line1, grid, centre, 1, line2)
	grid = drawPath(line2, grid, centre, 2, line1)
     findDistance(line1, line2, grid, centre)
}
