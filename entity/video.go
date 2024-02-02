package entity

type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       string `json:"age" binding:"gte=1, lte=90"`
	Email     string `json:"required,email"`
}

type Video struct {
	Title       string `json:"title" binding:"min=4,max=10"`
	Description string `json:"description" binding:"max=20"`
	URL         string `json:"url" binding:"required,url"`
	//Author      Person `json:"author" binding:"required"`
}
