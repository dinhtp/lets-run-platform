package platform

import (
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    pb "github.com/dinhtp/lets-run-pbtype/gateway"
)

func validateOne(r *pb.OnePlatformRequest) error {
    if r.GetId() == "" {
        return status.Error(codes.InvalidArgument, "platform url is required")
    }

    return nil
}

func validateCreate(r *pb.Platform) error {
    if r.GetName() == "" {
        return status.Error(codes.InvalidArgument, "platform name is required")
    }

    if r.GetApiUrl() == "" {
        return status.Error(codes.InvalidArgument, "platform url is required")
    }

    if r.GetAccessToken() == "" {
        return status.Error(codes.InvalidArgument, "access token is required")
    }

    if r.GetType() != TypeShopify && r.GetType() != TypeWooCommerce {
        return status.Errorf(codes.InvalidArgument, "platform type '%s' is not supported", r.GetType())
    }

    if r.GetType() == TypeWooCommerce && r.GetAccessTokenSecret() == "" {
        return status.Error(codes.InvalidArgument, "access token secret is required")
    }

    return nil
}
