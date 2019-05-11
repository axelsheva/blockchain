package types

type Transaction uint8

const (
	ReferralType  Transaction = 0
	SendType                  = 10
	SignatureType             = 20
	DelegateType              = 30
)
