# calendarmod

## Table of Contents
1. [Overview](#overview)
    1. [Version Update](#version-update)
2. [Getting started](#getting-started)
3. [Usage](#usage)
    1. [Subscribe User to a Calendar in Go Project](#subscribe-user-to-a-calendar-in-go-project)
    2. [Subscribe Group to a Calendar in Go Project](#subscribe-group-to-a-calendar-in-go-project)
4. [Functions](#functions)
    1. [SetUpSVAAuth](#func-setupsvaauth)
    2. [SubscribeUserToCalendar](#func-subscribeusertocalendar)
5. [Google Cloud API](#google-cloud-api)
    1. [Set up Google Cloud API Access](#set-up-google-cloud-api-access)
6. [Common Questions](#common-questions)
7. [Troubleshoot](#troubleshoot)
8. [Support](#support)
<br>
<br>

---

## Overview
**calendarmod** is a Go Package designed to facilitate Google Calendar services for Rutgers University. This module includes the following core functionality:

1. Authenticate Client for Google API using a service account
2. Subscirbe user to a dynamic Google calendar
3. Subscirbe a group of users to a dynamic Google calendar


### Version Update
Make sure you are using one of the stable version. Update to the newest minor version. 
- **v0.1.13**: Recommanded stable version 


<br>
<br>

---
## Getting started

#### Prerequisite
- Obtain a Google Service Account Key json file and enable the necessary Google API permissions for the Google Service Account.
    - For detailed instructions, refer to the section [Set up Google Cloud API Access](#set-up-google-cloud-api-access).
- Set up a Go project to utilize this package.

#### Installation
1. move to the directory of your Go Project in the terminal
2. run the command 
    ```
    go get github.com/oss/calendarmod
    ```
<br>
<br>

---

## Usage

#### Subscribe User to a Calendar in Go Project
1. Import the package
    ```
    import (
	    "github.com/oss/calendarmod"
    )
    ```

2. Create a authentification client based on the service account
    ```
    client := calendarmod.SetUpSVAClient(serviceAccountJSON, true)
    ```

3. Subscribe the user to the calendars
    ```
    success := client.SubscribeUserToCalendar(calendarid, userEmail)
    ```

#### Subscribe Group to a Calendar in Go Project
1. Import the package
    ```
    import (
	    "github.com/oss/calendarmod"
    )
    ```

2. Create a authentification client based on the service account
    ```
    client := calendarmod.SetUpSVAClient(serviceAccountJSON, true)
    ```

3. Subscribe the user to the calendars
    ```
    success := client.SubscribeGroupToCalendar(calendarid, userlist)
    ```
<br>
<br>

---
## Functions

#### func SetUpSVAClient
```
func SetUpSVAClient(serviceAccountJSON []byte, useCalendar bool) *CalendarClient
```
SetUpSVAClient initialize authentification client with service account and returns CalendarClient with context and config. 
- *useCalendar* should always to be set to true to use Calendar services. 
- To authentificate, make sure to set up all necessary permissions for the Google Service account.

Exp:
```
// set up Service Account authentification client
serviceAccountJSON, err := os.ReadFile("SERVICE_ACCOUNT_PATH")
if err != nil {
	log.Fatalf("Could not read service account credentials file, %s => {%s}", sap, err)
}

auth := calendarmod.SetUpSVAAuth(serviceAccountJSON, true)
```

#### func SubscribeUserToCalendar
```
func (c *CalendarClient) SubscribeUserToCalendar(user string, calendarid string) bool 
```
SubscribeUserToCalendar subscribes user to a dynamic Google Calendar. Call the function with the Calendar Client created from *func SetUpSVAClient*
- *user* must be a a valid google email address under the same domain of the Service Account client. 
- *calendarid* can be retrived from Google Calendar => Calendar settings. 

Exp:
```
calendarid:= "c_d3e80545746779e9e3957248314356fe4d9e1dcc27c2259b8c029ad5ee6f9cdf@group.calendar.google.com"
userEmail:= "user@gmail.com"
success := calendarmod.SubscribeUserToCalendar(calendarid, userEmail)
```

#### func SubscribeGroupToCalendar
```
func (c *CalendarClient) SubscribeGroupToCalendar(userlist_path string, calendarid string,
	success_user_file bool, success_user_name string, fail_user_file bool, fail_user_name string) bool
```
SubscribeUserToCalendar subscribes a groups of users to a dynamic Google Calendar. Call the function with the Calendar Client created from *func SetUpSVAClient*. <br>
This function produces two lists of user cases documenting the outcomes of user calendar subscription attempts. 
1. A list of successful subscription cases
2. A list of failed subscription attempts

The return boolean value indicates completion: true means process completed, false means process terminated due to error. This doesn't reflect the state of user subscription
- *calendarid* can be retrived from Google Calendar => Calendar settings. 
- *userlist* is list of valid Google emails of targeted users, should be in email format 

Exp:
```
userlist:= ["a@gamil.com", "b@gamil.com", "c@gamil.com"]
calendarid:= "c_d3e80545746779e9e3957248314356fe4d9e1dcc27c2259b8c029ad5ee6f9cdf@group.calendar.google.com"
result := calendarClient.SubscribeGroupToCalendar(userlist, calendarid)
```

<br>
<br>

---
## Google Cloud API

### Set up Google Cloud API Access
1. Log in to Google Cloud Console [Google Cloud Console](https://console.cloud.google.com/)
2. Select the corresponding Google Cloud Project
    1. If the project isn't created yet, Click New Project to create one. 
    2. Project existed but doesn't show up. 
        1. Search it in the search bar
        2. Couldn't find it. Refer to [*Google Cloud API Common Question*](#about-google-cloud-api)
3. Open the console left side menu => APIs & Services => Enable APIs & Services
4. Enable required API
    * Google Calender API (scope: https://www.googleapis.com/auth/calendar)s
5. Open the console left side menu => IAM & Admin => Service accounts
6. Select the corresponding Service Account
    * If the Service Account isn't created yet, Click + create a service account to create one.
    * For service account: Choose the role Project > owner.
7. [Delegating domain-wide authority to the service account](https://developers.google.com/identity/protocols/OAuth2ServiceAccount#delegatingauthority)
    - For OAuth scopes, enter: ttps://www.googleapis.com/auth/calendar
8. Click into the service account => KEYS => Add KEY => Create new key
9. Create private key for "calendar service test", choose "JSON" for key type
    - Save the JSON file in the projecr directory for authentification to the Service Account. 

<br>
<br>

---

## Common Questions
### About Google Cloud API

- **Q: Admin has created the project but I couldn't find it in my console.** 
    
    A: Check your access role to the project. Make sure you are a owner. 
    If need to change access role, refer to [*Manage project members or change project ownership*](####Manage-project-members-or-change-project-ownership:)

<br>
<br>

---

## Troubleshoot

- Error: Subscription: Authentification passed, couldn't fetch token

    ```
    Authentification sets up
    Subscribe User To Calendar...
    Calendar ID: c_b14ec2724ff6a0cac6558c74bebb0e8b986b4084cf059c68bb2d5b9b070d71ed@group.calendar.google.com
    user: acstst21@em-gmail.rutgers.edu
    Post "https://www.googleapis.com/calendar/v3/users/me/calendarList?alt=json&prettyPrint=false": oauth2: cannot fetch token: 401 Unauthorized
    Response: {
    "error": "unauthorized_client",
    "error_description": "Client is unauthorized to retrieve access tokens using this method, or client not authorized for any of the scopes requested."
    }
    panic: Post "https://www.googleapis.com/calendar/v3/users/me/calendarList?alt=json&prettyPrint=false": oauth2: cannot fetch token: 401 Unauthorized
    Response: {
    "error": "unauthorized_client",
    "error_description": "Client is unauthorized to retrieve access tokens using this method, or client not authorized for any of the scopes requested."
    }
    ```

    Solution: <br>
    1. Make sure to enable all APIs included in the scope. <br> 
        Check access to "calendarservices.SetUpSVAAuth()", make sure all required scope has enabled  <br> 
        To enable API: Open the console left side menu => APIs & Services => Enable APIs & Services
    

<br>
<br>

---
## Support
For assistance, please contact us at sk2779@oit.rutgers.edu.

Last updated by Seoli Kim on June 13, 2024.

