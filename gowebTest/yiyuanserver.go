package main

import "golang.org/x/net/context"

type SearchService struct {

}

func (s *SearchService) Search(ctx context.Context,r *pb.SearchRequest)