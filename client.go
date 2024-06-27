package calendarmod

import (
	"log"
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/calendar/v3"
)

// Client based on Service Account to connect to Google API
type CalendarClient struct {
	ctx        context.Context
	config     *jwt.Config
	httpClient *http.Client
}

// Initialize authentification client with service account to access Google API
//
//	@return {*CalendarClient}
//
//	@param {bool} useCalendar- true if need to use Google Calendar API, should always be true to use subscription service
func SetUpSVAClient(serviceAccountJSON []byte, useCalendar bool) *CalendarClient {
	// This is a variable needed for all http actions with the google API
	ctx := context.Background()

	var scope []string
	if useCalendar {
		scope = append(scope, calendar.CalendarScope)
	}

	config, err := google.JWTConfigFromJSON(serviceAccountJSON, scope...)
	if err != nil {
		log.Printf("Could not create config for service account=> {%s}", err)
		return nil
	}

	client := config.Client(oauth2.NoContext)

	// initilize authentification client
	calendarclient := &CalendarClient{ctx: ctx, config: config, httpClient: client}

	log.Println("Client sets up")
	return calendarclient
}
