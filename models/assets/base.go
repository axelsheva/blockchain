package assets

type Base struct{}

func (*Base) GetAmount() uint64 {
	return 0
}

func (*Base) CalculateFee() uint64 {
	return 0
}
