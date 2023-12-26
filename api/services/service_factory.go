package services

import "fmt"

// GetVerificationService returns the appropriate verification service for a given cloud provider
func GetVerificationService(cloudProvider string) (VerificationService, error) {
	switch cloudProvider {
	case "aws":
		return NewAWSVerificationService( /* params */ ), nil
	case "gcp":
		return NewGCPVerificationService( /* params */ ), nil
	case "azure":
		return NewAzureVerificationService( /* params */ ), nil
	default:
		return nil, fmt.Errorf("unsupported cloud provider: %s", cloudProvider)
	}
}
