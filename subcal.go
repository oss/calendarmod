package calendarmod

import (
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

type Subcal struct {
	Summary string `json:"summary" bson:"summary"`
	Id      string `json:"id" bson:"id"`
}

// Subscribe user to dynamic calendar
//
//	@param {string} user - a valid google email address under the same domain of the Service Account client
//	@param {string} calendarID - of the calendar for subscription. Can be retrived from Google Calendar => Calendar settings.
//	@return {bool} if success
func (c *CalendarClient) SubscribeUserToCalendar(user string, calendarID string) bool {
	log.Println("Subscribe User To Calendar...")
	log.Printf("Calendar ID: %s\n", calendarID)
	log.Printf("user: %s\n", user)
	serviceClient := UserInitiateService(c.ctx, c.config, user)
	if serviceClient == nil {
		log.Printf("Fail to initiate service client for the %s\n", user)
		return false
	}
	calendarListEntry := GetCalendarListEntry(calendarID)

	// create CalendarListService for user
	calendarListService := calendar.NewCalendarListService(serviceClient)
	if calendarListService == nil {
		log.Printf("The calendar Service for user: %s is null\n", user)
		return false
	}
	_, err := calendarListService.Insert(calendarListEntry).Do()
	if err != nil {
		log.Println(err)
		log.Println("Fail")
		return false
	}
	log.Println("Success")
	return true
}

// Unsubscribe user from dynamic calendar
//
//	@param {string} user - a valid google email address under the same domain of the Service Account client
//	@param {string} calendarID - of the calendar for subscription. Can be retrived from Google Calendar => Calendar settings.
//	@return {bool} if success
func (c *CalendarClient) UnsubscribeUserFromCalendar(user string, calendarID string) bool {
	log.Println("Unsubscribe User From Calendar...")
	log.Printf("Calendar ID: %s\n", calendarID)
	log.Printf("user: %s\n", user)
	serviceClient := UserInitiateService(c.ctx, c.config, user)
	if serviceClient == nil {
		log.Printf("Fail to initiate service client for the %s\n", user)
		return false
	}

	// create CalendarListService for user
	calendarListService := calendar.NewCalendarListService(serviceClient)
	if calendarListService == nil {
		log.Printf("The calendar Service for user: %s is null\n", user)
		return false
	}

	// check if user has subscribed to the calendar
	_, err := calendarListService.Get(calendarID).Do()
	if err != nil {
		// Check if the error is due to the calendar not being found
		if apiErr, ok := err.(*googleapi.Error); ok && apiErr.Code == 404 {
			// Calendar not found, which means user is not subscribed
			log.Printf("User %s is not subscribed to the calendar. \n", user)
			return true
		}
		log.Println(err)
		log.Println("Fail")
		return false
	}

	// delete
	err = calendarListService.Delete(calendarID).Do()
	if err != nil {
		log.Println(err)
		log.Println("Fail")
		return false
	}
	log.Println("Success")
	return true
}

// Initiate a service access with service account's impersonation of the user
//
// This service is for all functionalities to Google Calendar API
//
//	@return {*calendar.Service} if success
//	@param {string} user - impersanation of the user
func UserInitiateService(ctx context.Context, config *jwt.Config, user string) *calendar.Service {
	config.Subject = user
	//client := config.Client(context.Background())
	client := config.Client(oauth2.NoContext)
	calendarClient, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Println(err)
	}
	return calendarClient
}

// Get CalendarListEntry from Calendar ID.
//
// CalendarListEntry is Calendar representation on the user's calendar list (read-only)
//
//	@return {*calendar.CalendarListEntry}
//	@param {string} canlendarID - calendar ID
func GetCalendarListEntry(calendarid string) *calendar.CalendarListEntry {
	n_Cle := &calendar.CalendarListEntry{Id: calendarid}
	return n_Cle
}

// // subscirbe user to calendar
// //
// //	@return {*calendar.CalendarListEntry} updated userCalendarListEntry
// //	@param {calendar.CalendarListEntry} targetCalendarListEntry - CalendarListEntry of the calendar user is subscribed to
// //	@param {*calendar.Service} calendarClient
// func UserSubscribeToCalendar(calendarClient *calendar.Service, targetCalendarListEntry *calendar.CalendarListEntry) *calendar.CalendarListEntry {
// 	// create CalendarListService for user
// 	calendarListClient := calendar.NewCalendarListService(calendarClient)
// 	//func (r *CalendarListService) Insert(calendarlistentry *CalendarListEntry) *CalendarListInsertCall
// 	if calendarListClient == nil {
// 		return targetCalendarListEntry
// 	}

// 	userCalendarListEntry, err := calendarListClient.Insert(targetCalendarListEntry).Do()
// 	if err != nil {
// 		log.Println(err)
// 		panic(err)
// 	}
// 	return userCalendarListEntry
// }
