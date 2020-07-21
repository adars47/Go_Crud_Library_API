package models


type Books struct{
	Id		int		`json:"id"`
	Name 	string		`json:"name"`
	ISBN	string	`json:"isbn"`
	Author	Author	
}
