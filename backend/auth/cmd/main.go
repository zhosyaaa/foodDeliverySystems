package main

import (
	"github/zhosyaaa/foodDeliverySystems-auth-service/pkg/api/server"
	"github/zhosyaaa/foodDeliverySystems-auth-service/pkg/config"
	"github/zhosyaaa/foodDeliverySystems-auth-service/pkg/db"
	"github/zhosyaaa/foodDeliverySystems-auth-service/pkg/pb"
	"github/zhosyaaa/foodDeliverySystems-auth-service/pkg/utils"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config, err : %v", err)
	}

	db, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to init db, err : %v", err)
	}

	jwt := utils.JwtWrapper{
		SecretKey:       cfg.JwtSecretKey,
		Issuer:          "nabiels.com",
		ExpirationHours: time.Minute * time.Duration(cfg.JwtExpirationMinute),
	}

	addr := cfg.Host + ":" + cfg.Port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to create tcp connection, err : %v", err)
	}

	log.Printf("Auth Service started on : %v\n", addr)
	s := server.Service{
		DB:  db,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve, err : %v", err)
	}
}
