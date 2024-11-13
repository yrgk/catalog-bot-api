package s3

import (
	"catalog-bot-api/internal/config"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)


func UploadToS3(sess *session.Session, bucket, key, filePath string) (string, error) {
	s3Client := s3.New(sess)

	// Загружаем файл
	_, err := s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		// Добавьте настройки источника файла и тип данных по необходимости
	})
	if err != nil {
	 	return "", err
	}

	// Формируем URL файла
	fileURL := fmt.Sprintf("%s/%s/%s", config.Config.S3ApiUrl, config.Config.BucketName, key)
	return fileURL, nil
}