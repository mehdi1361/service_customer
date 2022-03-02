package server

import (
	"fmt"
	empty "github.com/golang/protobuf/ptypes/empty"
	"golang.org/x/net/context"
	models "service_customer/models"
	service "service_customer/service/proto"
)

func (Server) JobAll(ctx context.Context, e *empty.Empty) (*service.JobListResponse, error) {

	var data []*service.JobItemResponse

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	jobs := []models.Job{}
	db.Find(&jobs)

	for _, v := range jobs {
		data = append(data, &service.JobItemResponse{
			Id: int32(v.ID),
			Title: v.Title,
			JobId: v.JobId,
		},
		)
	}

	result := service.JobListResponse{
		Jobitemresponse: data,
	}

	return &result, nil
}

func (Server) GetJobById(ctx context.Context, e *service.JobItemRequest) (*service.JobItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	job := models.Job{}
	db.Find(&job, "id=?", e.Id)
	result := service.JobItemResponse{
		Id:   int32(job.ID),
		Title: job.Title,
		JobId: job.JobId,
	}
	return &result, nil
}

func (Server) CreateJob(ctx context.Context, e *service.CreateJobItemRequest) (*service.JobItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	job := models.Job{
		Title: e.Title,
		JobId: e.JobId,
	}
	db.Create(&job)
	result := service.JobItemResponse{
		Id:   int32(job.ID),
		Title: job.Title,
		JobId: job.JobId,
	}
	return &result, nil
}

func (Server) DeleteJob(ctx context.Context, e *service.JobItemRequest) (*service.DeleteJobItemResponse, error) {

	db, err := models.Connect()
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	db.Delete(&models.Job{}, e.Id)

	result := service.DeleteJobItemResponse{
		Message: "delete successfull",
	}
	return &result, nil
}
