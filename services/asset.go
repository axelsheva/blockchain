package services

type IAsset interface {
	ITransaction
	CalculateFee() uint64
}
