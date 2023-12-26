package services

import (
    "context"
    "cloud.google.com/go/storage"
    "google.golang.org/api/iterator"
    "google.golang.org/api/option"
)

// GCPService struct holds the GCP clients and related configurations
type GCPService struct {
    StorageClient *storage.Client
}

// NewGCPService creates a new instance of GCPService with initialized GCP clients
func NewGCPService(ctx context.Context, credsFilePath string) (*GCPService, error) {
    client, err := storage.NewClient(ctx, option.WithCredentialsFile(credsFilePath))
    if err != nil {
        return nil, err
    }

    return &GCPService{
        StorageClient: client,
    }, nil
}

// DeleteObject deletes an object from a GCS bucket and returns a boolean indicating success
func (svc *GCPService) DeleteObject(ctx context.Context, bucket, object string) (bool, error) {
    err := svc.StorageClient.Bucket(bucket).Object(object).Delete(ctx)
    if err != nil {
        return false, err
    }
    return true, nil
}

// VerifyDeletion checks if an object exists in a GCS bucket to confirm deletion
func (svc *GCPService) VerifyDeletion(ctx context.Context, bucket, object string) (bool, error) {
    _, err := svc.StorageClient.Bucket(bucket).Object(object).Attrs(ctx)
    if err == storage.ErrObjectNotExist {
        return true, nil
    }
    return false, err
}

// Add additional functions for other GCP services or more detailed operations as needed.
