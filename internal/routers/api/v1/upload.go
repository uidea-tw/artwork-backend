package v1

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/uidea/artwork-backend/global"
	"github.com/uidea/artwork-backend/internal/service"
	"github.com/uidea/artwork-backend/pkg/app"
	"github.com/uidea/artwork-backend/pkg/errcode"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) Get(c *gin.Context) {
	response := app.NewResponse(c)
	appEnv := os.Getenv("APP_ENV")
	var bucketName string
	if appEnv == "production" {
		bucketName = os.Getenv("MINIO_BUCKETNAME")
	} else {
		bucketName = "testbucket"
	}
	id := c.Param("id")

	svc := service.New(c.Request.Context())
	file, err := svc.GetUploadFile(id)

	if err != nil {
		global.Logger.Errorf(c, "❌ 無法取得檔案資訊:: %v", err)
		response.ToErrorResponse(errcode.ErrorGetFileFail.WithDetails(err.Error()))
		return
	}

	presignedURL, err := app.GetMinioPresignedURL(bucketName, file.StorageName)
	if err != nil {
		log.Println("❌ 無法產生 Presigned URL:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法取得圖片"})
		return
	}

	response.ToResponse(gin.H{
		"url": presignedURL.String(),
	})
}

func (u Upload) Create(c *gin.Context) {
	response := app.NewResponse(c)

	file, header, err := c.Request.FormFile("file")

	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	defer file.Close()

	// fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), header.Filename)

	fileUUID := uuid.New().String()

	appEnv := os.Getenv("APP_ENV")
	var bucketName string
	if appEnv == "production" {
		bucketName = os.Getenv("MINIO_BUCKETNAME")
	} else {
		bucketName = "testbucket"
	}

	// 設定 Content-Type
	contentType := header.Header.Get("Content-Type")
	newFileName := fileUUID + app.FileTypeFormat(contentType)
	uploadInfo, err := global.MinioClient.PutObject(
		context.Background(),
		bucketName,  // Bucket 名稱
		newFileName, // 檔案名稱
		file,        // 檔案內容
		header.Size, // 檔案大小
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		log.Println("❌ 上傳失敗:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "上傳失敗"})
		return
	}

	uploadFile := service.UploadFile{
		ID:          fileUUID,
		Filename:    header.Filename,
		StorageName: newFileName,
		Bucket:      "testbucket",
		ContentType: contentType,
		Size:        header.Size,
	}

	svc := service.New(c.Request.Context())
	err = svc.CreateUploadFile(uploadFile)
	if err != nil {
		global.Logger.Errorf(c, "❌ 儲存檔案資訊失敗:: %v", err)
		response.ToErrorResponse(errcode.ErrorStoreRecord.WithDetails(err.Error()))
		return
	}

	log.Println("✅ 上傳成功:", uploadInfo)

	response.ToResponse(gin.H{
		"id": fileUUID,
	})
}

func (u Upload) Delete(c *gin.Context) {
	response := app.NewResponse(c)
	id := c.Param("id")

	svc := service.New(c.Request.Context())

	file, err := svc.GetUploadFile(id)

	if err != nil {
		global.Logger.Errorf(c, "❌ 無法取得檔案資訊:: %v", err)
		response.ToErrorResponse(errcode.ErrorGetFileFail.WithDetails(err.Error()))
		return
	}

	err = global.MinioClient.RemoveObject(context.Background(), "testbucket", file.StorageName, minio.RemoveObjectOptions{})
	if err != nil {
		response.ToErrorResponse(errcode.ErrorDeleteFileFail.WithDetails(err.Error()))
		return
	}

	err = svc.DeleteUploadFile(id)
	if err != nil {
		global.Logger.Errorf(c, "❌ 刪除檔案資訊失敗:: %v", err)
		response.ToErrorResponse(errcode.ErrorStoreRecord.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"message": "刪除成功",
	})
}
