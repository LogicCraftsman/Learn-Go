package main

import (
	"context"
	"log"
	"net"

	invoicer "github.com/LogicCraftsman/Learn-Go/proto"
	"google.golang.org/grpc"
)

type InvoicerServerImpl struct {
	invoicer.UnimplementedInvoicerServer
}

func (s InvoicerServerImpl) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf: []byte("test"),
		Docx: []byte("test"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("Cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &InvoicerServerImpl{}

	invoicer.RegisterInvoicerServer(serverRegistrar, service)
	err = serverRegistrar.Serve(lis)

	if err != nil {
		log.Fatalf("Impossible to serve: %s", err)
	}
}