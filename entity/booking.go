package entity

import "time"

type Booking struct {
	Id                 string        `json:"id,omitempty"`
	EntityId           string        `json:"entityId,omitempty"`
	IsEntityActive     bool          `json:"isEntityActive"`
	EventEntityId      string        `json:"eventEntityId" binding:"required"`
	EventId            string        `json:"eventId"`
	ClientId           string        `json:"clientId" binding:"required"`
	ClientMobileNumber *string       `json:"clientMobileNumber"`
	ClientEmail        *string       `json:"clientEmail"`
	ClientFirstName    *string       `json:"clientFirstName"`
	ClientLastName     *string       `json:"clientLastName"`
	StartDateTimeLocal string        `json:"startDateTimeLocal" binding:"required"`
	EndDateTimeLocal   string        `json:"endDateTimeLocal" binding:"required"`
	Timezone           string        `json:"timezone" binding:"required"`
	Status             string        `json:"status"`
	CreatedAtUTC       time.Time     `json:"createdAtUTC"`
	Repetitions        *[]Repetition `json:"repetitions"`
}

type Repetition struct {
	StartDateTimeLocal string `json:"startDateTimeLocal"`
	EndDateTimeLocal   string `json:"endDateTimeLocal"`
}
