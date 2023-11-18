package main

import (
"net/http"
"github.com/gin-gonic/gin"
"errors" 
"fmt"
)

 type book struct{
	ID  string  `json:"id"`
	Title  string `json:"title"`
	Author  string `json:"author"`
	Quantity  int `json:"quantity"`
 }

 var books = []book{
    {ID: "1" , Title: "Atomic habits", Author: "James", Quantity:2},
    {ID: "2" , Title: "Atomic habits", Author: "James", Quantity:2},
    {ID: "3" , Title: "Atomic habits", Author: "James", Quantity:2}, 
 } 


 func getBooks(c *gin.Context) {
   c.IndentedJSON(http.StatusOK,books)
 }

 func createBook(c *gin.Context){
	var newBook book
    if err := c.BindJSON(&newBook); err != nil{
		return
	}

	books = append(books,newBook);
    c.IndentedJSON(http.StatusCreated,newBook)
 }

 //Single book fetching using id
 func bookById(c *gin.Context){
	id := c.Param("id")
	fmt.Printf(id)
	book, e := getBookById(id)

	if e != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	fmt.Println(book)

	c.IndentedJSON(http.StatusOK,book)
 }



 func getBookById(id string)(*book, error){
 for i,v := range books{
	if v.ID == id {
		return &books[i],nil 
	}
  }
  return nil, errors.New("Book is not found")
 }


 func checkoutBook(c *gin.Context){
    id,ok := c.GetQuery("id")
    if ok == false{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Missing id query param"})
		return
	}
   
	book,err := getBookById(id)
    if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book is not found"})
		return
	}

	if book.Quantity <=0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Out of stock"})
		return
	}

	book.Quantity -=1
	c.IndentedJSON(http.StatusOK,book)
}

func bookReturn(c *gin.Context){
	id,ok := c.GetQuery("id")
    if ok == false{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Missing id query param"})
		return
	}
   
	book,err := getBookById(id)
    if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book is not found"})
		return
	}

	if book.Quantity >=10  {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Maximum quantity reached"})
		return
	}

	book.Quantity +=1
	c.IndentedJSON(http.StatusOK,book)
}


 func main (){
    router:= gin.Default()
	router.GET("/books",getBooks)
	router.POST("/books",createBook)
	router.GET("/books/:id",bookById)
	router.PATCH("/checkout",checkoutBook)
	router.PATCH("/return",bookReturn)
    router.Run("localhost:8081")
 }