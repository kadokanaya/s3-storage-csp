package s3

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
)

type Config struct {
	Session *session.Session
	Bucket     string
}

// 環境変数上のキーを用いてSessionConfigを作成
func newS3Config() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	creds := credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY"), os.Getenv("PRIVATE_KEY"), "")
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String(os.Getenv("REGION")),
	}))
	return &Config{
		Session: sess,
		Bucket:     os.Getenv("BUCKET"),
	}, nil
}


func UploadContentS3(key string, file []byte) error {
	c, err := newS3Config()
	if err != nil {
		return err
	}
	f := bytes.NewReader(file)
	objectKey := key

	// Uploaderを作成し、ローカルファイルをアップロード
	uploader := s3manager.NewUploader(c.Session)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(c.Bucket),
		Key:    aws.String(objectKey),
		Body:   f,
	})
	if err != nil {
		return err
	}
	return nil
}

// 署名付きurlを取得 (key:string) -> string, error
func GetPreSignedURL(key string) (string, error) {
	c, err := newS3Config()
	if err != nil {
		return "", err
	}
	svc := s3.New(c.Session)

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(c.Bucket),
		Key:    aws.String(key),
	})
	u, err := req.Presign(3 * time.Minute) // 有効期限3分
	if err != nil {
		return "", err
	}
	return u, nil
}

// ファイルをダウンロード(key:string) -> *s3.GetObjectOutput, error
func DownLoadContentS3(key string) (*s3.GetObjectOutput, error) {
	c, err := newS3Config()
	if err != nil {
		return nil, err
	}
	objectKey := key

	// S3 clientを作成
	svc := s3.New(c.Session)

	// S3からファイルをダウンロードせずに読み込む
	obj, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(c.Bucket),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				err := "bucket does not exist"
				return nil, errors.New(err)
			case s3.ErrCodeNoSuchKey:
				err := fmt.Sprintf("object with key %s does not exist in bucket", objectKey)
				return nil, errors.New(err)
			default:
				return nil, aerr
			}
		}
	}
	return obj, nil
}
