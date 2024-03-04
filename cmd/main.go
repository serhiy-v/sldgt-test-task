package main

import (
	"github.com/gorilla/mux"
	"sldt-test-task/internal/server"
	"sldt-test-task/internal/service"
)

func main() {
	router := mux.NewRouter()

	verifySvc := service.NewVerficationService()
	srv := server.NewServer(router, verifySvc)

	srv.RunServer()

	//var luhn int
	//
	//number := 4441114469726432
	//
	//for i := 1; number > 0; i++ {
	//	fmt.Println(number)
	//	cur := number % 10
	//	altern := false
	//	fmt.Println("----")
	//	fmt.Println(cur)
	//
	//	if i%2 == 0 {
	//		cur = cur * 2
	//		if cur > 9 {
	//			cur -= 9
	//		}
	//	}
	//
	//	luhn += cur
	//	fmt.Println("LUHHHN")
	//	fmt.Println(luhn)
	//	number = number / 10
	//	altern = !altern
	//}

}
