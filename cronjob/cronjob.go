package cronjob

import (
	"context"
	"log"

	pb "github.com/Somvaded/cronjob-task/proto"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CronJob struct {
	cronjob *cron.Cron
}
func CreateCronJob() *CronJob {
	return &CronJob{
		cronjob: cron.New(),
	}
}
func (cj *CronJob)SetupCronjob() {
	_ , err :=cj.cronjob.AddFunc("@every 10s", func() {
		conn ,err := grpc.NewClient("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Failed to connect to gRPC server: %v", err)
			return
		}
		defer conn.Close()
		client := pb.NewReportServiceClient(conn)

		userid :=[]string{"sovajit1", "abhijit","surojeet"}

		for _ ,user := range userid{
			resp , err := client.GenerateReport(context.Background(), &pb.GenerateReportRequest{
				UserId: user,
			})
			if err != nil {
				log.Printf("Failed to generate report for user %s: %v\n", user, err)
			} else{
				log.Printf("Generated report for user %s received with report ID: %s\n", user, resp.ReportId)
			}
		}
	})
	if err != nil {
		log.Fatalf("Failed to add cron job: %v", err)
	}
}

func(cj *CronJob) Start(){
	cj.cronjob.Start()
	log.Println("Cron job started")
}

func(cj *CronJob) Stop() {
	cj.cronjob.Stop()
	log.Println("Cron job stopped")
}