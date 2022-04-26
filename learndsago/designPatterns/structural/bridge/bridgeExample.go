package bridge

import "fmt"

type IDrawShape interface {
	drawShape(x [5]float64, y [5]float64)
}

type DrawShape struct{}

func (drawShape DrawShape) drawShape(x [5]float64, y [5]float64) {
	fmt.Println("Drawing shape...")
}

type IContour interface {
	DrawContour(x [5]float64, y [5]float64)
	ResizeByFactor(factor int)
}

type DrawContour struct {
	X      [5]float64
	Y      [5]float64
	Shape  DrawShape
	Factor int
}

func (contour DrawContour) DrawContour(x [5]float64, y [5]float64) {
	fmt.Println("Drawing Contour")
	contour.Shape.drawShape(contour.X, contour.Y)
}

func (contour DrawContour) ResizeByFactor(factor int) {
	contour.Factor = factor
}

func Example() {
	var x = [5]float64{1, 2, 3, 4, 5}
	var y = [5]float64{1, 2, 3, 4, 5}
	var contour IContour = DrawContour{x, y, DrawShape{}, 2}
	contour.DrawContour(x, y)
	contour.ResizeByFactor(2)
}
