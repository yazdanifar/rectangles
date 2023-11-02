package dto

type RectanglesRequest struct {
	Main  Rectangle   `json:"main"`
	Input []Rectangle `json:"input"`
}

type Rectangle struct {
	X      int     `json:"x"`
	Y      int     `json:"y"`
	Width  int     `json:"width"`
	Height int     `json:"height"`
	Time   TimeDTO `json:"time"`
}
