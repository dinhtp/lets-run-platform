package platform

import (
    "fmt"
    pb "github.com/dinhtp/lets-run-pbtype/gateway"
    "github.com/dinhtp/lets-run-platform/model"
    "time"
)

const (
    TypeShopify     = "shopify"
    TypeWooCommerce = "woocommerce"
)

func preparePlatformToCreate(r *pb.Platform) *model.Platform {
    return &model.Platform{
        Type:              r.GetType(),
        Name:              r.GetName(),
        ApiUrl:            r.GetApiUrl(),
        AccessToken:       r.GetAccessToken(),
        AccessTokenSecret: r.GetAccessTokenSecret(),
    }
}

func preparePlatformToResponse(o *model.Platform) *pb.Platform {
    return &pb.Platform{
        Id:                fmt.Sprintf("%d", o.ID),
        Type:              o.Type,
        Name:              o.Name,
        ApiUrl:            o.ApiUrl,
        AccessToken:       o.AccessToken,
        AccessTokenSecret: o.AccessTokenSecret,
        CreatedAt:         o.CreatedAt.Format(time.RFC3339),
        UpdatedAt:         o.UpdatedAt.Format(time.RFC3339),
    }
}
