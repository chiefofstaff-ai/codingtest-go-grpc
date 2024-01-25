package server

import "cosgrpctest/protos"

type CosServer struct {
	protos.UnimplementedCosServiceServer
}
