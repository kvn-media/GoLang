package Day2



type Rectangle struct {
	width, length int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) SetLength(width int) {
	r.width = width
}

func (r *Rectangle) GetArea() int {
	return r.width * r.length
}	