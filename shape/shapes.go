package main

type shape interface{
	getArea() float64
}

type triangle struct{
	base float64
	height float64
}

func (t triangle) getArea() float64{
	return 0.5 * t.base * t.height
}

type square struct{
	sideLen float64
}

func (s square) getArea() float64{
	return s.sideLen * s.sideLen
}
