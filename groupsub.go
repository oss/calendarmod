package calendarmod

import (
	"encoding/csv"
	"log"
	"os"
)

// Subscribe group of user to group of calendar
//
//	@param {string} userlist_path - path to the csv file with a list of users
//	@param {string} calendarid - id of the calendar for the users to subscribe
//	@paran {bool} success_user_file - (optional) whether to generate a file that stores a list of users that successfully subscribe to calendars
//	@paran {string} success_user_path - (optional) path that points at success_user_file, default is "success_user_calendarid.csv" in current directory
//	@paran {bool} fail_user_file - (optional) whether to generate the file that stores a list of users that failed to subscribe to calendars
//	@paran {string} fail_user_path - (optional)path to store FAILUSER_PATH,  default is current directory
//
//	@return {bool} if completed: true means process completed, false means process terminated due to error.
func (c *CalendarClient) SubscribeGroupToCalendar(calendarid string, userlist_path string,
	success_user_file bool, success_user_path string, fail_user_file bool, fail_user_path string) bool {
	usercsv, err := os.Open(userlist_path)
	if err != nil {
		log.Printf("Unable to open file: %v", err)
	}
	defer usercsv.Close()
	// Create a new CSV reader to read usercsv
	ur := csv.NewReader(usercsv)

	// Skip the header row
	_, err = ur.Read()
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
	if success_user_file {

		// check and set name to default
		if success_user_path == "" {
			success_user_path = "success_user_" + calendarid + ".csv"
		}

		// Create the output file
		successusercsv, err := os.Create(success_user_path)
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
	if fail_user_file {

		// check and set name to default
		if fail_user_path == "" {
			fail_user_path = "fail_user_" + calendarid + ".csv"
		}

		// Create the output file
		failusercsv, err := os.Create(fail_user_path)
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
