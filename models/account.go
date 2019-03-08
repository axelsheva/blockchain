package models

type PublicKey string

type Account struct {
	PublicKey       PublicKey
	SecondPublicKey PublicKey
	Balance         uint64
	Delegate        *Delegate   // fix
	Votes           []PublicKey //
	Referrals       []PublicKey //
	Stakes          string      //
}
