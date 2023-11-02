package dto

type RectanglesRequest struct {
	Main  Rectangle   `json:"main"`
	Input []Rectangle `json:"input"`
}

type Rectangle struct {
	X      int     `json:"x"`
	Y      int     `json:"y"`
	Width  int     `json:"width" binding:"gt=0"`
	Height int     `json:"height" binding:"gt=0"`
	Time   TimeDTO `json:"time"`
}
