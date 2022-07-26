package platform

import (
    "context"

    "github.com/gogo/protobuf/types"
    "gorm.io/gorm"

    pb "github.com/dinhtp/lets-run-pbtype/gateway"
)

type Service struct {
    db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
    return &Service{db: db}
}

func (s Service) Get(ctx context.Context, r *pb.OnePlatformRequest) (*pb.Platform, error) {
    panic("implement me")
}

func (s Service) Create(ctx context.Context, r *pb.Platform) (*pb.Platform, error) {
    panic("implement me")
}

func (s Service) Update(ctx context.Context, r *pb.Platform) (*pb.Platform, error) {
    panic("implement me")
}

func (s Service) Delete(ctx context.Context, r *pb.OnePlatformRequest) (*types.Empty, error) {
    panic("implement me")
}

func (s Service) List(ctx context.Context, r *pb.ListPlatformRequest) (*pb.ListPlatformResponse, error) {
    panic("implement me")
}
