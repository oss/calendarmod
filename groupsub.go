package calendarmod

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
)

// Subscribe group of user to group of calendar
//
//	@param {*os.File} usercsv - csv file with a list of users
//	@param {string} calendarid - id of the calendar for the users to subscribe
//	@paran {bool} SUCCESSUSER_FILE - (optional) whether to generate a file that stores a list of users that successfully subscribe to calendars
//	@paran {*string} SUCCESSUSER_NAME - (optional) name of the SUCCESSUSER_FILE, default is "success_user_calendarid"
//	@paran {*string} SUCCESSUSER_PATH- (optional)path to store SUCCESSUSER_FILE, default is current directory
//	@paran {*string} FAILUSER_FAIL- (optional) whether to generate the file that stores a list of users that failed to subscribe to calendars
//	@paran {*string} FAILUSER_NAME- (optional) name of the FAILUSER_FILE , default is "fail_user_calendarid"
//	@paran {*string} FAILUSERUSER_PATH- (optional)path to store FAILUSER_PATH,  default is current directory
func (c *Client) SubscribeGroupToCalendar(usercsv *os.File, calendarid string,
	SUCCESSUSER_FILE bool, SUCCESSUSER_NAME string, SUCCESSUSER_PATH string,
	FAILUSER_FILE bool, FAILUSER_NAME string, FAILUSER_PATH string) bool {
	// Create a new CSV reader to read usercsv
	ur := csv.NewReader(usercsv)
	// Skip the header row
	_, err := ur.Read()
	if err != nil {
		log.Printf("Unable to read header row netid, please add a header row: %v", err)
		return false
	}

	// Create user list
	var users []string

	// Read the users
	for {
		record, err := ur.Read()
		if err != nil {
			break
		}

		// Parse id
		id := record[0]
		// add userid to the list of users
		users = append(users, id)
	}

	var successuserlist []string
	var failuserlist []string

	for _, u := range users {
		res := c.SubscribeUserToCalendar(u, calendarid)
		if res {
			successuserlist = append(successuserlist, u)
		} else {
			failuserlist = append(failuserlist, u)
		}
	}

	// create successful user file
	if SUCCESSUSER_FILE {

		// check and set name to default
		if SUCCESSUSER_NAME == "" {
			SUCCESSUSER_NAME = "success_user_" + calendarid + ".csv"
		} else {
			SUCCESSUSER_NAME = SUCCESSUSER_NAME + ".csv"
		}

		// Combine directory path and file name
		fullPath := filepath.Join(SUCCESSUSER_PATH, SUCCESSUSER_NAME)

		// Create the output file
		successusercsv, err := os.Create(fullPath)
		if err != nil {
			log.Printf("Unable to create output file: %v", err)
		}
		defer successusercsv.Close()

		// Create a new CSV writer
		writer := csv.NewWriter(successusercsv)
		defer writer.Flush()

		// Write header row
		header := []string{"NetID"}
		if err := writer.Write(header); err != nil {
			log.Printf("Error writing header: %v", err)
		}

		// Write data rows
		for _, netID := range successuserlist {
			record := []string{netID}
			if err := writer.Write(record); err != nil {
				log.Printf("Error writing record: %v", err)
			}
		}
	}

	// create successful user file
	if FAILUSER_FILE {

		// check and set name to default
		if FAILUSER_NAME == "" {
			FAILUSER_NAME = "success_user_" + calendarid + ".csv"
		} else {
			FAILUSER_NAME = FAILUSER_NAME + ".csv"
		}

		// Combine directory path and file name
		fullPath := filepath.Join(FAILUSER_PATH, FAILUSER_NAME)

		// Create the output file
		failusercsv, err := os.Create(fullPath)
		if err != nil {
			log.Printf("Unable to create output file: %v", err)
		}
		defer failusercsv.Close()

		// Create a new CSV writer
		writer := csv.NewWriter(failusercsv)
		defer writer.Flush()

		// Write header row
		header := []string{"NetID"}
		if err := writer.Write(header); err != nil {
			log.Printf("Error writing header: %v", err)
		}

		// Write data rows
		for _, netID := range failuserlist {
			record := []string{netID}
			if err := writer.Write(record); err != nil {
				log.Printf("Error writing record: %v", err)
			}
		}
	}

	return true
}
