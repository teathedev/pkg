// Package jwt contains JWT and JWK related types and functions.
// types.go holds JWK/JWT key types so this package has no external app dependencies.

package jwt

import (
	"crypto/rsa"

	"github.com/google/uuid"
)

// JWKKeyType (kty) - RFC 7517 Section 4.1
type JWKKeyType string

const (
	JWKKeyTypeRSA JWKKeyType = "RSA"
	JWKKeyTypeEC  JWKKeyType = "EC"
)

// JWKKeyUse (use) - RFC 7517 Section 4.2
type JWKKeyUse string

const (
	JWKKeyUseSignature  JWKKeyUse = "sig"
	JWKKeyUseEncryption JWKKeyUse = "enc"
)

// JWKAlgorithm (alg) - RFC 7518 JWA
type JWKAlgorithm string

const JWKAlgorithmRS256 JWKAlgorithm = "RS256"

// JWK is the public key representation for OIDC discovery and verification (RFC 7517).
type JWK struct {
	KTY       JWKKeyType   `json:"kty"`
	Use       JWKKeyUse    `json:"use"`
	Algorithm JWKAlgorithm `json:"alg"`
	KeyID     uuid.UUID    `json:"kid"`
	N         string       `json:"n,omitempty"`
	E         string       `json:"e,omitempty"`
}

// JWTSigningKey holds the private key for in-memory JWT signing.
type JWTSigningKey struct {
	KeyID      uuid.UUID
	PrivateKey *rsa.PrivateKey
	Algorithm  JWKAlgorithm
}

// AccessTokenKeyStates is the lifecycle state for persisted keys (e.g. ent model).
type AccessTokenKeyStates int8

const (
	AccessTokenKeyStatesActive   AccessTokenKeyStates = 0
	AccessTokenKeyStatesPrevious AccessTokenKeyStates = 1
	AccessTokenKeyStatesRetired  AccessTokenKeyStates = 2
)
