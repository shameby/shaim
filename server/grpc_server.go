package server

import (
	"context"

	pb "shaim/proto"
)

type IoGServer struct{}

func (igs *IoGServer) BroadCast(ctx context.Context, in *pb.BcRequest) (*pb.BcReply, error) {
	suc := S.BroadcastToRoom(in.Nsp, in.Room, "msg", in.Msg)
	if !suc {
		return &pb.BcReply{Suc: false, Err: "board casting err"}, nil
	}
	return &pb.BcReply{Suc: true, Err: ""}, nil
}

func (igs *IoGServer) Say(ctx context.Context, in *pb.SayRequest) (*pb.SayReply, error) {
	if sId, exist := nameId[in.ToUser]; !exist {
		return &pb.SayReply{Suc: false, Err: in.ToUser + " 用户不存在"}, nil
	} else {
		if toConn, exist := conM[sId]; !exist {
			return &pb.SayReply{Suc: false, Err: in.ToUser + " 用户不存在"}, nil
		} else {
			toConn.Emit("msg", in.Msg)
		}
	}
	return &pb.SayReply{Suc: true, Err: ""}, nil
}
