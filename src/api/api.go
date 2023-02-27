package api

import (
	"context"
	"fmt"

	"storage/models"
	pb "storage/proto"
	"storage/repository"
)

type Server struct {
	pb.UnimplementedStorageServer
	memcached       repository.Repository
	internalStorage repository.Repository
}

func NewServer(memcached repository.Repository, internalStorage repository.Repository) *Server {
	return &Server{
		memcached:       memcached,
		internalStorage: internalStorage,
	}
}

func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	resp := &pb.GetResponse{}
	var (
		item  models.StorageItem
		found bool
		err   error
	)

	switch req.Destination {
	case "inmemory":
		item, found, err = s.internalStorage.Get(req.Key)
	case "memcached":
		item, found, err = s.memcached.Get(req.Key)
	default:
		return resp, fmt.Errorf("Not valid storage")
	}
	if err != nil {
		return resp, err
	}
	if !found {
		return resp, fmt.Errorf("Value doesnt found")
	}

	resp.Key = item.Key
	resp.Expire = item.Expiration
	resp.Value = item.Data

	return resp, nil
}

func (s *Server) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	resp := &pb.SetResponse{}
	var (
		item models.StorageItem
		err  error
	)

	switch req.Destination {
	case "inmemory":
		item, err = s.internalStorage.Set(req.Key, req.Value, req.Expire)
	case "memcached":
		item, err = s.memcached.Set(req.Key, req.Value, req.Expire)
	default:
		return resp, fmt.Errorf("Not valid storage")
	}
	if err != nil {
		return resp, err
	}

	resp.Expire = item.Expiration
	resp.Value = item.Data
	resp.Key = item.Key

	return resp, nil
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	switch req.Destination {
	case "inmemory":
		s.internalStorage.Delete(req.Key)
	case "memcached":
		s.memcached.Delete(req.Key)
	default:
		return &pb.DeleteResponse{}, fmt.Errorf("Not valid storage")
	}

	return &pb.DeleteResponse{}, nil
}
