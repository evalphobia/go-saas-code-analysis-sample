package main

import (
	_ "crypto/sha256"

	_ "golang.org/x/crypto/argon2"
	_ "golang.org/x/crypto/blowfish"
	_ "golang.org/x/crypto/chacha20poly1305"
	_ "golang.org/x/crypto/scrypt"
	_ "golang.org/x/net/http2"

	_ "github.com/hashicorp/vault/api"
)
