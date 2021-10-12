package common

type DomainEventType int

const (
	DOMAIN_EVENT_UNKNOWN DomainEventType = iota
	DOMAIN_EVENT_USER_REGISTRATION
	DOMAIN_EVENT_LOYALTY_CARD_CREATE
)
