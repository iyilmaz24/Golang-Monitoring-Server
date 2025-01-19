package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iyilmaz24/Golang-Notification-Server/internal/models"
)

func (app *application) verifyPostRequest (w http.ResponseWriter, r *http.Request) (error) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return fmt.Errorf("invalid request method: %s", r.Method);
	}
	return nil;
}

func (app *application) getDailyAnalyticsObject(w http.ResponseWriter, r *http.Request) (models.DailyAnalytics, error) { // marshalls and returns the dailyAnalytics object from the POST request body
	var analyticsObj models.DailyAnalytics;
	err := json.NewDecoder(r.Body).Decode(&analyticsObj);

	if err != nil || analyticsObj.NotificationSubject == "" || analyticsObj.NotificationMessage == "" {
		app.clientError(w, http.StatusBadRequest);
		return models.DailyAnalytics{}, fmt.Errorf("error decoding JSON: %v", err);
	}
	
	return analyticsObj, nil;
}

func (app *application) getNotificationObject(w http.ResponseWriter, r *http.Request) (models.Notification, error) { // marshalls and returns the notification object from the POST request body
	var notificationObj models.Notification;
	err := json.NewDecoder(r.Body).Decode(&notificationObj);

	if err != nil || notificationObj.NotificationSubject == "" || notificationObj.NotificationMessage == "" {
		app.clientError(w, http.StatusBadRequest);
		return models.Notification{}, fmt.Errorf("error decoding JSON: %v", err);
	}
	
	return notificationObj, nil;
}
