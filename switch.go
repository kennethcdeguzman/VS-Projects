package main

type httpRequest struct {
	Method string
}

func main() {
	i := httpRequest{Method: "TEST"}
	switch i.Method {
		case "PUT":
			println("Method is PUT")
		case "POST":
			println("Method is POST")
		case "DELETE":
			println("Method is DELETE")
		case "GET":
			println("Method is GET")
		default:
			println("Unhandled value")
	}
}

