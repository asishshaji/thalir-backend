package enum

type ProductType int

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
