package versions

import (
    "gorm.io/gorm"
    "time"
)

func Version20211215000000(tx *gorm.DB) error {
    type Platform struct {
        ID                uint   `gorm:"TYPE:BIGINT(20) UNSIGNED AUTO_INCREMENT;NOT NULL;PRIMARY_KEY"`
        Type              string `gorm:"TYPE:VARCHAR(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
        Name              string `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
        ApiUrl            string `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
        AccessToken       string `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;NOT NULL"`
        AccessTokenSecret string `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
        CreatedAt         time.Time
        UpdatedAt         time.Time
        DeletedAt         gorm.DeletedAt `gorm:"index"`
    }

    return tx.AutoMigrate(&Platform{})
}
