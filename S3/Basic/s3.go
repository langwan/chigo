package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
)

type Config struct {
	Key    string `yaml:"key"`
	Secret string `yaml:"secret"`
	//https://oss-cn-hangzhou.aliyuncs.com
	Endpoint string `yaml:"endpoint"`
}

type S3FileSystem struct {
	client *s3.Client
}

func NewFileSystem(key, secret, endpoint, session string) (*S3FileSystem, error) {
	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "oss",
			URL:           endpoint,
			SigningRegion: region,
		}, nil
	})
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(resolver), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(key, secret, session)))
	if err != nil {
		return nil, err
	}
	s3Client := s3.NewFromConfig(cfg)
	if err != nil {
		return nil, err
	}
	return &S3FileSystem{
		client: s3Client,
	}, nil
}

func (fs S3FileSystem) PutFile(ctx context.Context, bucket, key string, body io.Reader, contentType string, len int64) error {
	_, err := fs.client.PutObject(ctx, &s3.PutObjectInput{Bucket: aws.String(bucket), Key: aws.String(key), Body: body, ContentType: aws.String(contentType), ContentLength: len})
	return err
}

func (fs S3FileSystem) RemoveFile(ctx context.Context, bucket, key string) error {
	_, err := fs.client.DeleteObject(ctx, &s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(key)})
	return err
}

func (fs S3FileSystem) List(ctx context.Context, bucket, key string) ([]string, error) {
	var items []string
	output, err := fs.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{Bucket: aws.String(bucket), Prefix: aws.String(key), Delimiter: aws.String("/")})
	if err != nil {
		return nil, err
	}
	for _, folder := range output.CommonPrefixes {
		items = append(items, *folder.Prefix)
	}
	for _, file := range output.Contents {
		items = append(items, *file.Key)
	}
	return items, nil
}
