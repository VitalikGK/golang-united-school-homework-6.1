package golang_united_school_homework

import (
	"errors"
	"fmt"
	"reflect"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		//shapes:         make([]Shape, shapesCapacity, shapesCapacity),
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
		return errors.New("превышен размер коробки")
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
	if i > b.shapesCapacity-1 || i < 0 {
		// err := errors.New("Превышен размер коробки")
		// fmt.Println("the expected length is 1, but actual %w", err)
		return nil, errors.New("вышли за пределы")
	} else {
		//b.shapes = append(b.shapes, shape)
		fmt.Println(" i = ", i)
		fmt.Println(" b.shapes = ", b.shapes[i])
		return b.shapes[i], nil
	}

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	//	panic("implement me")
	if i > b.shapesCapacity-1 || i < 0 {
		// err := errors.New("Превышен размер коробки")
		// fmt.Println("the expected length is 1, but actual %w", err)
		return nil, errors.New("вышли за пределы")
	} else {
		//b.shapes = append(b.shapes, shape)
		fmt.Println(" i = ", i)
		fmt.Println(" b.shapes = ", b.shapes[i])
		bb := b.shapes[i]
		copy(b.shapes[i:], b.shapes[i+1:])
		b.shapes = b.shapes[:len(b.shapes)-1]
		return bb, nil
	}

}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	//	panic("implement me")
	if i > b.shapesCapacity-1 || i < 0 {
		// err := errors.New("Превышен размер коробки")
		// fmt.Println("the expected length is 1, but actual %w", err)
		return nil, errors.New("вышли за пределы")
	} else {
		//b.shapes = append(b.shapes, shape)
		fmt.Println(" i = ", i)
		fmt.Println(" b.shapes = ", b.shapes[i])
		bb := b.shapes[i]
		b.shapes[i] = shape
		//b.shapes = b.shapes[:len(b.shapes)-1]
		return bb, nil
	}

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	//panic("implement me")
	var sumPerimeter float64 = 0
	for i := 0; i < len(b.shapes); i++ {
		fmt.Println("Shapes = ", b.shapes[i])
		fmt.Println("perim = ", b.shapes[i].CalcPerimeter())
		sumPerimeter = sumPerimeter + b.shapes[i].CalcPerimeter()

	}
	return sumPerimeter
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	//	panic("implement me")
	var sumArea float64 = 0
	for i := 0; i < len(b.shapes); i++ {
		sumArea = sumArea + b.shapes[i].CalcArea()

	}
	return sumArea

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	//	panic("implement me")
	circle := &Circle{}
	c := 0
	var err error
	for i := 0; i < len(b.shapes); i++ {
		if reflect.TypeOf(b.shapes[i]) == reflect.TypeOf(circle) {
			copy(b.shapes[i:], b.shapes[i+1:])
			b.shapes = b.shapes[:len(b.shapes)-1]
			c = c + 1
			i = i - 1
		}
	}
	if c == 0 {
		err = errors.New("нет кругов")
	} else {
		err = nil
	}
	return err

}
