package models

type PublicKey string

type Account struct {
	PublicKey       PublicKey
	SecondPublicKey PublicKey
	Balance         int64
	Delegate        *Delegate   // fix
	Votes           []PublicKey //
	Referral        *Account    //
	Stakes          string      //
}
