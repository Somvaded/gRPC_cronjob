package server

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	pb "github.com/Somvaded/cronjob-task/proto"
)


type ReportService struct {
	pb.UnimplementedReportServiceServer
	Reports map[string]string
	Mutex   sync.Mutex
}

func NewReportService() *ReportService{
	return &ReportService{
		Reports: make(map[string]string),
	}
}

func (rs *ReportService) GenerateReport(ctx context.Context , req *pb.GenerateReportRequest) (*pb.GenerateReportResponse,error){
	rs.Mutex.Lock()
	defer rs.Mutex.Unlock()

	if req.UserId == "" {
		log.Println("Received request with empty user_id")
		return &pb.GenerateReportResponse{
			Error: "user_id cannot be empty",
		}, nil
	}

	reportID := fmt.Sprintf("report-%s-%d",req.UserId,time.Now().Unix())

	rs.Reports[reportID]  = fmt.Sprintf("Report for user %s generated at %s", req.UserId, time.Now().Format(time.RFC3339))

	log.Printf("Generated a report for user_id %s with report_id %s", req.UserId, reportID)

	return &pb.GenerateReportResponse{
		ReportId: reportID,
	}, nil
}

func(rs *ReportService) HealthCheck(ctx context.Context , req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{
		Status: "good health",
	},nil
}