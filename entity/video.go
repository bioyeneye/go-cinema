package entity

type Video struct {
	Id 			int		`json:"id" binding:"required"`
	Title		string	`json:"title" binding:"required"`
	Description string	`json:"description" binding:"required,min=20"`
	Url 		string	`json:"url" binding:"required"`
}
