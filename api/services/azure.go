package services

import (
	"context"
	"net/url"
	"strings"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

// AzureService struct holds the Azure clients and related configurations
type AzureService struct {
	ServiceURL azblob.ServiceURL
}

// NewAzureService creates a new instance of AzureService with initialized Azure clients
func NewAzureService(accountName, accountKey string) (*AzureService, error) {
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		return nil, err
	}

	pipeline := azblob.NewPipeline(credential, azblob.PipelineOptions{})
	u, _ := url.Parse("https://" + accountName + ".blob.core.windows.net")
	serviceURL := azblob.NewServiceURL(*u, pipeline)

	return &AzureService{
		ServiceURL: serviceURL,
	}, nil
}

// DeleteBlob deletes a blob from an Azure storage container and returns a boolean indicating success
func (svc *AzureService) DeleteBlob(ctx context.Context, containerName, blobName string) (bool, error) {
	containerURL := svc.ServiceURL.NewContainerURL(containerName)
	blobURL := containerURL.NewBlockBlobURL(blobName)

	_, err := blobURL.Delete(ctx, azblob.DeleteSnapshotsOptionInclude, azblob.BlobAccessConditions{})
	if err != nil {
		return false, err
	}
	return true, nil
}

// VerifyDeletion checks if a blob exists in an Azure storage container to confirm deletion
func (svc *AzureService) VerifyDeletion(ctx context.Context, containerName, blobName string) (bool, error) {
	containerURL := svc.ServiceURL.NewContainerURL(containerName)
	blobURL := containerURL.NewBlockBlobURL(blobName)

	_, err := blobURL.GetProperties(ctx, azblob.BlobAccessConditions{}, azblob.ClientProvidedKeyOptions{})
	if err != nil {
		if strings.Contains(err.Error(), "BlobNotFound") {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

// Add additional functions for other Azure services or more detailed operations as needed.
