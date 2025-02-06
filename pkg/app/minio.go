package app

import (
	"context"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/uidea/artwork-backend/global"
)

func CheckBuckets(minioClient *minio.Client) error {
	appEnv := os.Getenv("APP_ENV")
	var bucketName string
	if appEnv == "production" {
		bucketName = os.Getenv("MINIO_BUCKETNAME")
	} else {
		bucketName = "testbucket"
	}
	ctx := context.Background()
	err := minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
			return nil
		} else {
			return err
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
	return nil
}

func GetMinioPresignedURL(bucketName string, objectName string) (u *url.URL, err error) {
	ctx := context.Background()
	expiry := time.Duration(100) * time.Minute
	reqParams := make(url.Values)
	presignedURL, err := global.MinioClient.PresignedGetObject(ctx, bucketName, objectName, expiry, reqParams)
	if err != nil {
		log.Println("❌ 無法產生 Presigned URL:", err)
		return nil, err
	}
	return presignedURL, nil
}
