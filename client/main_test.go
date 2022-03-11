package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	pb "github.com/n6o/grpc-go-sandbox/gen-client"
)

func Test_Mock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	pb.NewMockEchoClient(ctrl)
}
