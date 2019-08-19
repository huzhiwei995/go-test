package main

import (
	"golang.org/x/net/context"
	pb "go-grpc/proto"
	"google.golang.org/grpc"
	"net"
	"log"
)

type SearchService struct {}

func (s *SearchService) Search(ctx context.Context,r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + "Server"},nil
}
const PORT = "9001"

func main() {
	server := grpc.NewServer()
	pb.RegisterSearchServiceServer(server,&SearchService{})
	lis, err := net.Listen("tcp",":"+PORT)
	if err != nil {
		log.Fatal("net.Listen err: %v",err)
	}
	server.Serve(lis)
}