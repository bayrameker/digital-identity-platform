package utils

import (
    "mime/multipart"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
    "fmt"
)

func UploadFileToS3(file multipart.File, fileHeader *multipart.FileHeader, key string) (string, error) {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"),
    })
    if err != nil {
        return "", err
    }

    uploader := s3manager.NewUploader(sess)

    _, err = uploader.Upload(&s3manager.UploadInput{
        Bucket: aws.String("your-s3-bucket"),
        Key:    aws.String(key),
        Body:   file,
        ContentType: aws.String(fileHeader.Header.Get("Content-Type")),
    })
    if err != nil {
        return "", err
    }

    fileURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", "your-s3-bucket", key)
    return fileURL, nil
}
