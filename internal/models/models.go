package models

// notification methods: email, sms
// notification urgency: high, medium, low
// notification recipient: string indicating who the notification is for
// notification status: string (pending, retrying)
// notification id: string (unique identifier)
// notification type: string indicating what type of notification it is (CRON, internal, testing)
// notification source: string indicating where the notification came from (specific server, database, website)
// notification time: string (HH:MM)
// notification date: string (MM/DD/YYYY)
// notification timezone: string (EST, CST, PST)
// notification subject: string 1 sentence summary
// notification message: string explanation

type Notification struct {
	NotificationMethod string `json:"method"`
	NotificationUrgency string `json:"urgency"`
	NotificationRecipient string `json:"recipient"`
	NotificationStatus string `json:"status"`
	NotificationID string `json:"id"`
	NotificationType string `json:"type"`
	NotificationSource string `json:"source"`
	NotificationTime string `json:"time"`
	NotificationDate string `json:"date"`
	NotificationTimezone string `json:"timezone"`
	NotificationSubject string `json:"subject"`
	NotificationMessage string `json:"message"`
	AccessSecret string `json:"password"`
}

// notification recipient: string indicating who the notification is for
// notification time: string (HH:MM)
// notification date: string (MM/DD/YYYY)
// notification timezone: string (EST, CST, PST)
// notification subject: string 1 sentence summary
// notification message: string explanation

type DailyAnalytics struct {
	NotificationRecipient string `json:"recipient"`
	NotificationTime string `json:"time"`
	NotificationDate string `json:"date"`
	NotificationTimezone string `json:"timezone"`
	NotificationSubject string `json:"subject"`
	NotificationMessage string `json:"message"`
}