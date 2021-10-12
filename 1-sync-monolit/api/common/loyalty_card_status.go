package common

type LoyaltyCardStatus int

const (
	LOYALTY_CARD_CREATED LoyaltyCardStatus = iota
	LOYALTY_CARD_EXPIRED
	LOYALTY_CARD_STOPED
)
