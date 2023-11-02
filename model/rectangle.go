package model

import "time"

type Rectangle struct {
	X      int
	Y      int
	Width  int
	Height int
	Time   time.Time
}

func (r *Rectangle) Intersects(other Rectangle) bool {
	return r.X <= other.X+other.Width &&
		r.X+r.Width >= other.X &&
		r.Y <= other.Y+other.Height &&
		r.Y+r.Height >= other.Y
}
