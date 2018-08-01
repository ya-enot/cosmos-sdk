package types

import crypto "github.com/ya-enot/go-crypto"

type StdSignature struct {
	crypto.PubKey // optional
	crypto.Signature
	Sequence int64
}
