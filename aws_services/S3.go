package aws_services

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateS3(ctx *pulumi.Context) (*s3.Bucket, error) {
	bucket, err := s3.NewBucket(ctx, "test-s3-bucket", nil)
	fmt.Println("Inside Create S3 Function")
	fmt.Println(bucket.Bucket.ToStringOutput())
	return bucket, err
}


