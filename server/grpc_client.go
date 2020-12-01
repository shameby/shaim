package server

import (
	"log"
	"context"
	"errors"

	"shaim/util/redisC"
	"shaim/conf"
	"shaim/util/grpcP"
	pb "shaim/proto"
)

func grpcBroadcasting(nsp, room, msg string) error {
	rpcServerList, err := getServerList()
	if err != nil {
		return err
	}
	for _, addr := range rpcServerList {
		conn, err := grpcP.GetConn(addr)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		client := pb.NewIoClient(conn)
		res, err := client.BroadCast(context.Background(), &pb.BcRequest{
			Nsp:  nsp,
			Room: room,
			Msg:  msg,
		})
		if err != nil {
			log.Println(err.Error())
			continue
		}
		if res.Err != "" {
			log.Println(res.Err)
			continue
		}
	}
	return nil
}

func grpcSay(nsp, room, toName, msg string) error {
	host, err := getServerByName(room, toName)
	if host == "" || err != nil {
		return err
	}
	conn, err := grpcP.GetConn(host)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	client := pb.NewIoClient(conn)
	res, err := client.Say(context.Background(), &pb.SayRequest{
		Nsp:    nsp,
		ToUser: toName,
		Msg:    msg,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}
	if res.Err != "" {
		log.Println(res.Err)
		return errors.New(res.Err)
	}
	return nil
}

func getServerList() (list []string, err error) {
	list, err = redisC.Conn.Set.SMembers(0, conf.RedisCheckListKey).Strings()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func getServerByName(room, username string) (host string, err error) {
	host, err = redisC.Conn.Hash.HGet(0, conf.RedisCheckHashKey+room, username).String()
	if err != nil {
		return "", errors.New("该用户不存在")
	}
	return host, nil
}
