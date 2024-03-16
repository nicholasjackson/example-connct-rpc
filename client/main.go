package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	v1 "github.com/nicholasjackson/demo-connect-rpc/gen/greet/v1"
	"github.com/nicholasjackson/demo-connect-rpc/gen/greet/v1/greetv1connect"
)

func main() {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)

	gr, err := client.Greet(context.Background(), &connect.Request[v1.GreetRequest]{
		Msg: &v1.GreetRequest{Name: "YouTube"},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(gr.Msg.Greeting)
}
