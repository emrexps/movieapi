package movie

type Movie struct {
	Id          int    `json:"id"`
	MovieName   string `json:"moviename"  binding:"required"`
	ReleaseYear int    `json:"releaseyear"  binding:"required"`
	DirectedBy  string `json:"directedby"  binding:"required"`
	Genre       string `json:"genre"  binding:"required"`
}
