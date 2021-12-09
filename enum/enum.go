package enum

type ProductType int
type EnvMode int

const (
	Development EnvMode = iota
	Production
)

func (env EnvMode) String() string {
	switch env {
	case Development:
		return "DEV"
	case Production:
		return "PROD"
	default:
		return "DEV"
	}
}

const (
	Packet ProductType = iota
	Raw
)

func (pt ProductType) String() string {
	switch pt {
	case Packet:
		return "P"
	case Raw:
		return "R"
	default:
		return "P"
	}
}
