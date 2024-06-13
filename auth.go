package calendarmod

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/drive/v2"
	"google.golang.org/api/gmail/v1"
)

// Client based on Service Account to connect to Google API
type AuthentificationClient struct {
	ctxPtr    *context.Context
	configPtr **jwt.Config
	clientPtr **http.Client
}

// global authentification client
// Access through Context() and Config
var auth *AuthentificationClient

// Return global context
func (auth *AuthentificationClient) Context() context.Context {
	ctx := *auth.ctxPtr
	return ctx
}

// Return global config
func (auth *AuthentificationClient) Config() *jwt.Config {
	config := *auth.configPtr
	return config
}

// Return service account client
func (auth *AuthentificationClient) Client() *http.Client {
	client := *auth.clientPtr
	return client
}

// Initialize authentification client with service account
//
//	@return {*AuthentificationClient}
//	Create a global var auth *calendarservice.AuthentificationClient to access auth context and config
//
//	@param {bool} useDrive- true if need to use Google Drive API
//	@param {bool} useGmail- true if need to use Google Gmail API
//	@param {bool} useCalendar- true if need to use Google Calendar API
func SetUpSVAAuth(serviceAccountJSON []byte, useDrive bool, useGmail bool, useCalendar bool) *AuthentificationClient {
	// This is a variable needed for all http actions with the google API
	ctx := context.Background()

	var scope []string
	if useDrive {
		scope = append(scope, drive.DriveScope)
	}
	if useGmail {
		scope = append(scope, gmail.GmailReadonlyScope)
	}

	if useCalendar {
		scope = append(scope, calendar.CalendarScope)
	}

	config, err := google.JWTConfigFromJSON(serviceAccountJSON, scope...)
	if err != nil {
		log.Fatalf("Could not create config for service account=> {%s}", err)
	}

	client := config.Client(oauth2.NoContext)

	// initilize authentification client
	auth = &AuthentificationClient{ctxPtr: &ctx, configPtr: &config, clientPtr: &client}

	fmt.Println("Authentification sets up")
	return auth
}