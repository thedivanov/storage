package main

import (
	"net"

	"storage/api"
	pb "storage/proto"
	"storage/repository/internalstorage"
	"storage/repository/memcached"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		log.WithError(err).Error("parseConfig error")
		return
	}

	memcach, err := memcached.NewMemcahced(cfg.MemcachedAddr)
	if err != nil {
		log.WithError(err).Error("Memcahced init error")
		return
	}

	internalstorage := internalstorage.NewInternalStorage()

	listner, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.WithError(err).Error("TCP server init error")
		return
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterStorageServer(grpcServer, api.NewServer(memcach, internalstorage))

	log.Println("Server is starting on address:", cfg.GRPCAddr)
	if err := grpcServer.Serve(listner); err != nil {
		log.WithError(err)
		return
	}
}
