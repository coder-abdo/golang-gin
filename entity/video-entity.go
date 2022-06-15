package entity

type Person struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Age       int    `json:"age" binding:"required,gte=1,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}
type Video struct {
	Title       string `json:"title" binding:"min=2,max=20"`
	Description string `json:"description" binding: "max=500"`
	Url         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
