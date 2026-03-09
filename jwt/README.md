# jwt

JWT and JWK utilities: key generation, encryption at rest, and JWK for OIDC discovery.

## What it does

- **GenerateKey**: RSA key pair for JWT signing; encrypted private key for storage.
- **EncryptPrivateKey** / **DecryptPrivateKey**: AES-GCM for storing private keys (e.g. in DB).
- **ParsePublicPEM**: Parse a PEM-encoded public key.
- **ConvertToJWK**: Turn public PEM + key ID into RFC 7517 JWK for `/.well-known/jwks.json`.
- Types: **GeneratedKey**, **JWTSigningKey**, **JWK**, **JWKSet**, **AccessTokenKeyStates** (for key lifecycle).

## Installation

```bash
go get github.com/teathedev/pkg/jwt
```

## Usage

```go
import (
	"github.com/google/uuid"
	"github.com/teathedev/pkg/jwt"
)

// 1. Generate a key (e.g. at app startup or key rotation)
masterKey := []byte("32-byte-master-key-for-aes-256!!")
gen, err := jwt.GenerateKey(masterKey)
if err != nil {
	// handle
}

// 2. Persist: gen.ID, gen.PrivateKeyEncrypted, gen.PublicPem, gen.State
// 3. For signing: use gen.ToSigningKey() (in-memory)
signingKey := gen.ToSigningKey()

// 4. For JWK discovery endpoint: convert stored public key + ID to JWK
jwk, err := jwt.ConvertToJWK(gen.PublicPem, gen.ID)
if err != nil {
	// handle
}
// Serve jwk in your JWK set (e.g. {"keys": [jwk]})

// Encrypt/decrypt existing key material (e.g. for storage)
encrypted, _ := jwt.EncryptPrivateKey(privateKeyPEM, masterKey)
decrypted, err := jwt.DecryptPrivateKey(encrypted, masterKey)

// Parse a public key from PEM string (e.g. from DB)
pub, err := jwt.ParsePublicPEM(publicPemString)
```

## Dependencies

- `github.com/google/uuid` (key IDs)

## API overview

| Function / Type                                                   | Description                                          |
| ----------------------------------------------------------------- | ---------------------------------------------------- |
| `GenerateKey(masterKey []byte) (*GeneratedKey, error)`            | New RSA key; private key encrypted with masterKey.   |
| `EncryptPrivateKey(privateKey, masterKey []byte) ([]byte, error)` | AES-GCM encrypt key for storage.                     |
| `DecryptPrivateKey(encrypted, masterKey []byte) ([]byte, error)`  | Decrypt stored key.                                  |
| `ParsePublicPEM(pemStr string) (*rsa.PublicKey, error)`           | Parse PEM to RSA public key.                         |
| `ConvertToJWK(publicPem string, keyID uuid.UUID) (*JWK, error)`   | Build JWK for discovery.                             |
| `GeneratedKey`                                                    | ID, private key, encrypted bytes, public PEM, state. |
| `GeneratedKey.ToSigningKey() *JWTSigningKey`                      | In-memory signing key.                               |
| `JWK`, `JWTSigningKey`, `AccessTokenKeyStates`                    | Types for JWK and key lifecycle.                     |
