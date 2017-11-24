//https://golang.org/doc/articles/wiki/
//https://github.com/data-representation/go-ajax/blob/master/webapp.go
package main
import (
	"fmt"
	"net/http"
	"./chatbot" 
)
func main() {
	dir := http.Dir("./static") // directs to static folder
	fileServer := http.FileServer(dir)
	http.Handle("/", fileServer) // handle resource handles everything coming in
	http.HandleFunc("/ask", HandleAsk) // ask resource
	http.ListenAndServe(":8080", nil) //starts web server 
}
func HandleAsk(writer http.ResponseWriter, request *http.Request) { // takes user input from request and sends to ask. then writes result back to ResponsWriter
	userInput := request.URL.Query().Get("userInput")
	answer := chatbot.Ask(userInput) 
	fmt.Fprintf(writer, answer)  
}
