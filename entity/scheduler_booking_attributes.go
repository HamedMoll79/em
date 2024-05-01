package entity

type SchedulerBookingAttributes struct {
	EventTitle    string `json:"event_title"`
	EventEntityId string `json:"event_entity_id" binding:"required"`
	EventId       string `json:"event_id"`
	// EventId desc: because of old data this field cannot be required
	// and booking attributes excel report will be empty for old data,
	// but from now this should be required and must be sent by frontend
	StartDateTimeLocal string `json:"start_date_time_local" binding:"required"`
	EndDateTimeLocal   string `json:"end_date_time_local" binding:"required"`
	Timezone           string `json:"timezone" binding:"required"`
}

type BookingAttributes SchedulerBookingAttributes