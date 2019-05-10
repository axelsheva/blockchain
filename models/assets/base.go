package assets

type Base struct{}

func (*Base) GetAmount() int64 {
	return 0
}

func (*Base) CalculateFee() int64 {
	return 0
}
