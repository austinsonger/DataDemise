// Custom error handling specific to security operations, like authentication failures or encryption errors.

package security

import "errors"

// Custom security error definitions
var ErrUnauthorized = errors.New("unauthorized access")
var ErrEncryptionFailed = errors.New("encryption failed")
