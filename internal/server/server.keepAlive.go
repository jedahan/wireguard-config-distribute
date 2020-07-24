package server

import (
	"context"
	"fmt"

	"github.com/gongt/wireguard-config-distribute/internal/protocol"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *serverImplement) KeepAlive(context.Context, *emptypb.Empty) (*protocol.KeepAliveStatus, error) {
	fmt.Println("Call to KeepAlive")
	return nil, nil
}
