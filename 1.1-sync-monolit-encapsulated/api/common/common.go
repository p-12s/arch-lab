package common

// User - пользователь
type User struct {
	Id        int    `json:"-" db:"id"`
	Username  string `json:"username" db:"username"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	Phone     string `json:"phone" db:"phone"`
	Address   string `json:"address" db:"address"`
}

// LoyaltyCard - карта лояльности
type LoyaltyCard struct {
	Id     int               `json:"-" db:"id"`
	UserId int               `json:"user_id" db:"user_id"`
	Code   string            `json:"code" db:"code"`
	Status LoyaltyCardStatus `json:"status" db:"status"`
}

// Notification - уведомление
type Notification struct {
	Id                      int `json:"-" db:"id"`
	UserId                  int `json:"user_id" db:"user_id"`
	NotificationProgramType int `json:"notification_program_type" db:"notification_program_type"`
	DomainEventType         int `json:"domain_event_type" db:"domain_event_type"`
	Status                  int `json:"status" db:"status"`
}

type NotificationSendReq struct {
	Email      string `json:"email" db:"-"`
	TemplateId int    `json:"template_id" db:"-"`
	Context    struct {
		FirstName         string `json:"first_name" db:"-"`
		LastName          string `json:"last_name" db:"-"`
		LoyaltyCardNumber string `json:"loyalty_card_number" db:"-"`
	}
}
