package main

import (
	"infrastructure/aws_services"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		bucket, err := aws_services.CreateS3(ctx)

		if(err != nil){
			return err
		}

		policy, err := aws_services.CreateIAMPolicy(ctx, bucket)

		if(err != nil){
			return err
		}

		// // Export the ARN (Amazon Resource Name) of the created policy
		ctx.Export("policyArn", policy.Arn)
		ctx.Export("Bucket Name", bucket.Arn)
		return nil

	})

}
