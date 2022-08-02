package model

import (
    "gorm.io/gorm"
)

type Platform struct {
    gorm.Model
    Type              string
    Name              string
    ApiUrl            string
    AccessToken       string
    AccessTokenSecret string
}
