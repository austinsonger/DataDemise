package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// AWSService struct holds the AWS clients and related configurations
type AWSService struct {
	S3Client *s3.S3
}

// NewAWSService creates a new instance of AWSService with initialized AWS clients
func NewAWSService() *AWSService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("your-region"), // specify the AWS region
	}))

	return &AWSService{
		S3Client: s3.New(sess),
	}
}

// DeleteObject deletes an object from an S3 bucket and returns a boolean indicating success
func (svc *AWSService) DeleteObject(bucket, key string) (bool, error) {
	_, err := svc.S3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return false, err
	}

	// Optionally, add logic to wait for the deletion to be confirmed.
	// This can be done using the WaitUntilObjectNotExists function with the required input.

	return true, nil
}

// VerifyDeletion checks if an object exists in an S3 bucket to confirm deletion
func (svc *AWSService) VerifyDeletion(bucket, key string) (bool, error) {
	_, err := svc.S3Client.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		// Handle the specific error to determine if it's because the object doesn't exist
		return true, nil
	}
	return false, nil
}

// Generate additional functions as needed for other AWS services or operations.
