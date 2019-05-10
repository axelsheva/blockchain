package assets

type Delegate struct {
	Base
	Username string `json:"username"`
}

func (*Delegate) CalculateFee() uint64 {
	return 1000000000
}
