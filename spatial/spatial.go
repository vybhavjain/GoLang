package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
)

//The data is stored in a single cell of field
type Cell struct {
	strategy string  //Represent C or D corresponding to the prisoner in cell
	score    float64 //Represent score of the cell based on prisoners relationship with neighbouring cell
}

//The gameboard is a 2D slice of Cell objects
type GameBoard [][]Cell

//Drawing the GameBoard (from drawing.go)
func DrawGameBoards(boards []GameBoard, cellWidth int) []image.Image {
	numGenerations := len(boards)
	imageList := make([]image.Image, numGenerations)
	for i := range boards {
		imageList[i] = DrawGameBoard(boards[i], cellWidth)
	}
	return imageList
}

func DrawGameBoard(board GameBoard, cellWidth int) image.Image {
	height := len(board) * cellWidth
	width := len(board[0]) * cellWidth
	c := CreateNewPalettedCanvas(width, height, nil)

	// declare colors
	blue := MakeColor(0, 0, 255)
	red := MakeColor(255, 0, 0)
	white := MakeColor(255, 255, 255)

	// draw the grid lines in white
	c.SetStrokeColor(white)
	DrawGridLines(c, cellWidth)

	// fill in colored squares
	for i := range board {
		for j := range board[i] {
			if board[i][j].strategy == "C" {
				c.SetFillColor(blue)
			} else if board[i][j].strategy == "D" {
				c.SetFillColor(red)
			} else {
				panic("Error: Out of range value in board when drawing board.")
			}
			x := j * cellWidth
			y := i * cellWidth
			c.ClearRect(x, y, x+cellWidth, y+cellWidth)
			c.Fill()
		}
	}

	return c.img
}

func DrawGridLines(pic Canvas, cellWidth int) {
	w, h := pic.Width(), pic.Height()
	// first, draw vertical lines
	for i := 1; i < w/cellWidth; i++ {
		y := i * cellWidth
		pic.MoveTo(0.0, float64(y))
		pic.LineTo(float64(w), float64(y))
	}
	// next, draw horizontal lines
	for j := 1; j < h/cellWidth; j++ {
		x := j * cellWidth
		pic.MoveTo(float64(x), 0.0)
		pic.LineTo(float64(x), float64(h))
	}
	pic.Stroke()
}

//ReadBoardFromFile takes a filename as a string and reads in the data provided in this file (from reading_writing.go)
func ReadBoardFromFile(filename string) []string {
	//Returning data
	data := make([]string, 0)

	//Open the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Scan through the file to append lines to data
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

//InitializeBoard takes a number of rows and columns as inputs and returns a gameboard with appropriate number of rows and colums
func InitializeBoard(numRows, numCols int, board_initialData []string) GameBoard {
	// make a 2-D slice (default values = false)
	var board_initial GameBoard
	board_initial = make(GameBoard, numRows)
	// now we need to make the rows too
	for r := range board_initial {
		board_initial[r] = make([]Cell, numCols)
	}

	return InitialBoardStrategy(board_initialData, board_initial)
}

func InitialBoardStrategy(board_initialData []string, board_initial GameBoard) GameBoard {

	if len(board_initialData) == 0 {
		for m := range board_initial {
			for n := range board_initial[m] {
				board_initial[m][n].strategy = string(board_initialData[m][n])
			}

		}
	} else {
		for m := range board_initial {
			for n := range board_initial[m] {
				board_initial[m][n].strategy = "C"
			}
		}
	}
	return board_initial
}

/*func CountRows(board GameBoard) int {
	return len(board)
}

func CountCols(board GameBoard) int {
	// assume that we have a rectangular board
	if CountRows(board) == 0 {
		panic("Error: empty board given to CountCols")
	}
	// give # of elements in 0-th row
	return len(board[0])
}
*/
func UpdateBoard(currBoard GameBoard, b float64) GameBoard {
	// first, create new board corresponding to the next generation.
	// let's have all cells have state 0 to begin.
	//numRows := CountRows(currBoard)
	//numCols := CountCols(currBoard)
	Board_1 := InitializeBoard(len(currBoard), len(currBoard[0]), []string{})
	Board_2 := InitializeBoard(len(currBoard), len(currBoard[0]), []string{})

	//now, update values of newBoard
	//range through all cells of currBoard and update each one into newBoard.
	for r := range Board_1 {
		for c := range Board_1[r] {
			Board_1[r][c].score = Score(currBoard, r, c, b)
			Board_1[r][c].strategy = currBoard[r][c].strategy
			Board_2[r][c].score = Board_1[r][c].score
		}
	}
	for r := range Board_1 {
		for c := range Board_1[r] {
			Board_2[r][c].strategy = UpdateStrategy(Board_1, r, c, b)
		}

	}

	return Board_2
}

func UpdateBoards(currBoard GameBoard, numGens int, b float64) []GameBoard {
	numBoards := make([]GameBoard, numGens+1)
	numBoards[0] = currBoard

	for i := 1; i < numGens+1; i++ {
		numBoards[i] = UpdateBoard(numBoards[i-1], b)
	}
	return numBoards

}

func Score(currBoard GameBoard, row, col int, b float64) float64 {
	var score_n float64 = 0
	if currBoard[row][col].strategy == "C" {
		for x := row - 1; x <= row+1; x++ {
			for y := col - 1; y <= col+1; y++ {
				if x == row && y == col {
					continue
				}
				if ValidateCells(currBoard, x, y) {
					if currBoard[x][y].strategy == "C" {
						score_n += 1
					}
				}
			}
		}
	}
	if currBoard[row][col].strategy == "D" {
		for x := row - 1; x <= row+1; x++ {
			for y := col - 1; y <= col+1; y++ {
				if x == row && y == col {
					continue
				}
				if ValidateCells(currBoard, x, y) {
					if currBoard[x][y].strategy == "C" {
						score_n += b
					}
				}
			}
		}
	}
	return score_n
}

func ValidateCells(currBoard GameBoard, r, c int) bool {
	numRow := len(currBoard)
	numCol := len(currBoard[0])
	if r > numRow-1 || r < 0 || c > numCol-1 || c < 0 {
		return false
	}
	return true
}

func UpdateStrategy(currBoard GameBoard, row, col int, b float64) string {
	strategy_max := currBoard[row][col].strategy
	score_max := currBoard[row][col].score
	for x := row - 1; x <= row+1; x++ {
		for y := col - 1; y <= col+1; y++ {
			if x >= 0 || y >= 0 || x <= row-1 || y < col-1 {
				if currBoard[x][y].score > score_max {
					score_max = currBoard[x][y].score
					strategy_max = currBoard[x][y].strategy
				}
			}
		}
	}
	return strategy_max
}

func ReadRowCol(data []string) (int, int, []string) {
	Row, errR := strconv.Atoi(strings.Split(data[0], " ")[0])
	if errR != nil {
		fmt.Println("Could not transfer row from string to integer")
		os.Exit(2)
	}

	Col, errC := strconv.Atoi(strings.Split(data[0], " ")[1])
	if errC != nil {
		fmt.Println("Could not transfer column from string to integer")
		os.Exit(2)
	}
	return Row, Col, data[1:]
}

func ImageToPNG(image image.Image, filename string) {
	w, err := os.Create(filename + ".png")

	if err != nil {
		fmt.Println("Sorry, cannot create a png file")
		os.Exit(3)
	}

	defer w.Close()

	png.Encode(w, image)
}

//ImagesToGIF() takes a slice of images and uses them to generate an animated GIF
// with the name "filename.out.gif" where filename is an input parameter.
func ImagesToGIF(imglist []image.Image, filename string) {

	// get ready to write images to files
	w, err := os.Create(filename + ".gif")

	if err != nil {
		fmt.Println("Sorry: couldn't create the file!")
		os.Exit(1)
	}

	defer w.Close()
	var g gif.GIF
	g.Delay = make([]int, len(imglist))
	g.Image = make([]*image.Paletted, len(imglist))
	g.LoopCount = 10

	for i := range imglist {
		g.Image[i] = ImageToPaletted(imglist[i])
		g.Delay[i] = 1
	}

	gif.EncodeAll(w, &g)
}

func ImageToPaletted(img image.Image) *image.Paletted {
	pm, ok := img.(*image.Paletted)
	if !ok {
		b := img.Bounds()
		pm = image.NewPaletted(b, palette.WebSafe)
		var prevC color.Color = nil
		var idx uint8
		var ok bool
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				c := img.At(x, y)
				if c != prevC {
					if idx, ok = mapOfColorIndices[c]; !ok {
						idx = uint8(pm.Palette.Index(c))
						mapOfColorIndices[c] = idx
					}
					prevC = c
				}
				i := pm.PixOffset(x, y)
				pm.Pix[i] = idx
			}
		}
	}
	return pm
}

var mapOfColorIndices map[color.Color]uint8

func init() {
	mapOfColorIndices = make(map[color.Color]uint8)
}

func main() {

	file := os.Args[1]

	b, err1 := strconv.ParseFloat(os.Args[2], 64)
	if err1 != nil {
		panic("Error: Problem converting string b to float")
	}

	numGens, err2 := strconv.Atoi(os.Args[3])
	if err2 != nil {
		panic("Error: Problem converting number of generations to an integer.")
	}

	output := "Prisoners"
	data := ReadBoardFromFile(file)
	cellWidth := 20
	r, c, read_board := ReadRowCol(data)
	currBoard := InitializeBoard(r, c, read_board)

	fmt.Println("Files read successfully")

	update_board := UpdateBoards(currBoard, numGens, b)
	fmt.Println("Spatial Games implemented successfully!")

	// we need a slice of image objects
	imglist := DrawGameBoards(update_board, cellWidth)
	fmt.Println("Boards drawn to images! Now, convert to animated GIF.")

	//Convert images to png
	ImageToPNG(imglist[len(imglist)-1], output)
	fmt.Println("PNG produced.")

	// convert images to a GIF
	ImagesToGIF(imglist, output)
	fmt.Println("GIF produced.")
}
