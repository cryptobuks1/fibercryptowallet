package util

import (
	"github.com/fibercrypto/fibercryptowallet/src/core"
)

func NewGenericAddress(addr string) GenericAddress {
	return GenericAddress{
		Address: addr,
	}
}

// GenericAddress transient editable crypto address
type GenericAddress struct {
	Address string
}

func (ga *GenericAddress) IsBip32() bool {
	return false
}

func (ga *GenericAddress) String() string {
	return ga.Address
}

func (ga *GenericAddress) GetCryptoAccount() core.CryptoAccount {
	return nil
}

// Type assertions
var (
	_ core.Address = &GenericAddress{}
)