package upload

import (
	"bufio"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/fastService/s3"
	"os"
	"time"
)

func News3() *s3.S3 {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Endpoint: aws.String("http://oss-cn-hangzhou.aliyuncs.com"),
			Region:   aws.String("oss")},
		Profile: "aliyun",
	}))

	svc := s3.New(sess)
	resp, _ := svc.ListBuckets(&s3.ListBucketsInput{})
	for _, bucket := range resp.Buckets {
		fmt.Println(*bucket.Name)
	}
	return svc
}

func upload(fastService *s3.S3) {
	fp, _ := os.Open("s3_test.go")

	defer fp.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()

	fastService.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String("hatlonely"),
		Key:    aws.String("test/s3_test.go"),
		Body:   fp,
	})

}

func download(fastService *s3.S3) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()

	out, _ := fastService.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String("hatlonely"),
		Key:    aws.String("test/s3_test.go"),
	})

	defer out.Body.Close()
	scanner := bufio.NewScanner(out.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//目录遍历
func ListObjectsPages(fastService *s3.S3) {
	var objkeys []string

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()

	fastService.ListObjectsPagesWithContext(ctx, &s3.ListObjectsInput{
		Bucket: aws.String("hatlonely"),
		Prefix: aws.String("test/"),
	}, func(output *s3.ListObjectsOutput, b bool) bool {
		for _, content := range output.Contents {
			objkeys = append(objkeys, aws.StringValue(content.Key))
		}
		return true
	})
	fmt.Println(objkeys)
}