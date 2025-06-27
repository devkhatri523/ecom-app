package enum

type PaymentMethod int

const (
	UNKONW = iota
	PAYPAL
	CREDIT_CARD
	VISA
	MASTER_CARD
	BITCOIN
)

func (p PaymentMethod) String() string {
	return [...]string{"UNKWON", "PAYPAL", "CREDIT_CARD", "VISA", "MASTER_CARD", "BITCOIN"}[p]
}
