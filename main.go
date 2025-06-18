package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Somvaded/cronjob-task/cronjob"
	pb "github.com/Somvaded/cronjob-task/proto"
	"github.com/Somvaded/cronjob-task/server"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting gRPC service...")

	lis , err := net.Listen("tcp",":1234")
	if err != nil {
		log.Fatalf("Failed to listen on port 1234: %v", err)
	}
	grpcServer := grpc.NewServer()

	reportService := server.NewReportService()

	pb.RegisterReportServiceServer(grpcServer , reportService)

	cronjobFunc := cronjob.CreateCronJob()
	cronjobFunc.SetupCronjob()
	cronjobFunc.Start()

	go func(){
		sig := make(chan os.Signal,1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		log.Println("Received shutdown signal, stopping cron job and gRPC server...")
		cronjobFunc.Stop()
		log.Println("gRPC server stopped gracefully")
		grpcServer.GracefulStop()
	}()

	log.Println("gRPC server is running on port 1234")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}