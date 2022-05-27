package golang_united_school_homework

import (
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapes:         make([]Shape, shapesCapacity, shapesCapacity),
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	//	panic("implement me")
	// var err error = nil
	// d := len(b.shapes)
	// fmt.Println("d = ", d)
	if len(b.shapes) >= b.shapesCapacity {
		// err := errors.New("Превышен размер коробки")
		// fmt.Println("the expected length is 1, but actual %w", err)
		return errors.New("Превышен размер коробки")
	} else {
		b.shapes = append(b.shapes, shape)
		return nil
	}

	//	return err
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	//	panic("implement me")
	if i > b.shapesCapacity {
		// err := errors.New("Превышен размер коробки")
		// fmt.Println("the expected length is 1, but actual %w", err)
		return nil, errors.New("Вышли за пределы")
	} else {
		//b.shapes = append(b.shapes, shape)
		return b.shapes[i-1], nil
	}

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	panic("implement me")

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	panic("implement me")

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	panic("implement me")

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	panic("implement me")

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	panic("implement me")

}
