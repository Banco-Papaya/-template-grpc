package main

import (
	"context"
	"log"
	pb "template-grpc/internal/infra/ploto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("error al conectarse al grpc")
	}

	// Code removed for brevity

	client := pb.NewUserCrudClient(conn)

	// Note how we are calling the GetBookList method on the server
	// This is available to us through the auto-generated code
	bookList, err := client.Insert(context.Background(), &pb.User{
		Name: "pepapgi",
	})
	if err != nil {
		panic("error al insertar el usuario todo poderos")
	}

	log.Printf("book list: %v", bookList)
}
