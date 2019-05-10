package services

type IAsset interface {
	ITransaction
	CalculateFee() int64
}
