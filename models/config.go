package models

type CoreConfig struct {
	PublicHost  string
	RPCPort     int32
	Version     string
	ForgeSecret string
}

type Const struct {
	ActiveDelegatesCount uint16
	SlotInterval         uint8
}
