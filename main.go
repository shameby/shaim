package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"shaim/server"
	"shaim/conf"
	"shaim/check"
	"net"
	"log"
	"google.golang.org/grpc"
	pb "shaim/proto"

	_ "net/http/pprof"
)

func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("origin")
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

func main() {
	go check.Heartbeat()
	router := gin.New()

	go server.S.Serve()
	defer server.S.Close()

	go func() {
		lis, err := net.Listen("tcp", ":"+ conf.C.RpcPort)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		pb.RegisterIoServer(grpcServer, &server.IoGServer{})
		grpcServer.Serve(lis)
	}()

	go func() {
		http.ListenAndServe("0.0.0.0:8088", nil)
	}()
	router.Use(GinMiddleware())
	router.GET("/socket.io/*any", gin.WrapH(server.S))
	router.POST("/socket.io/*any", gin.WrapH(server.S))
	router.StaticFS("/public", http.Dir("../asset"))

	router.Run(":" + conf.C.IoPort)
}
