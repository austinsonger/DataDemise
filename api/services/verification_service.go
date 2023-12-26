package services

// VerificationService defines the interface for a service that verifies data destruction
type VerificationService interface {
	VerifyDestruction(cloudProvider, identifier string) (bool, error)
}
