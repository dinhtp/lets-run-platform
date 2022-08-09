package platform

import (
    "context"
    "strconv"

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
    if err := validateOne(r); err != nil {
        return nil, err
    }

    id, _ := strconv.Atoi(r.GetId())

    result, err := NewRepository(s.db).FindOne(id)
    if err != nil {
        return nil, err
    }

    return preparePlatformToResponse(result), nil
}

func (s Service) Create(ctx context.Context, r *pb.Platform) (*pb.Platform, error) {
    if err := validateCreate(r); err != nil {
        return nil, err
    }

    result, err := NewRepository(s.db).Create(preparePlatformToCreate(r))
    if err != nil {
        return nil, err
    }

    return preparePlatformToResponse(result), nil
}

func (s Service) Update(ctx context.Context, r *pb.Platform) (*pb.Platform, error) {
    // TODO: implement logic
    return &pb.Platform{}, nil
}

func (s Service) Delete(ctx context.Context, r *pb.OnePlatformRequest) (*types.Empty, error) {
    // TODO: implement logic
    return &types.Empty{}, nil
}

func (s Service) List(ctx context.Context, r *pb.ListPlatformRequest) (*pb.ListPlatformResponse, error) {
    // TODO: implement logic
    return &pb.ListPlatformResponse{}, nil
}
