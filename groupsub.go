package calendarmod

import (
	"encoding/csv"
	"log"
	"os"
)

// Subscribe group of user to group of calendar
//
//	@param {string[]} userlist - list of valid Google emails of targeted users, in email format
//	@param {string} calendarid - id of the calendar for the users to subscribe
//	@paran {bool} success_user_file - (optional) whether to generate a file that stores a list of users that successfully subscribe to calendars
//	@paran {string} success_user_name - (optional) custume name for success_user_file, must be ".csv". Default is "success_user_calendarid.csv"
//	@paran {bool} fail_user_file - (optional) whether to generate the file that stores a list of users that failed to subscribe to calendars
//	@paran {string} fail_user_name - (optional) custume name for fail_user_file, must be ".csv". Default is "fail_user_calendarid.csv"
//
//	@return {bool} if completed: true means process completed, false means process terminated due to error.
func (c *CalendarClient) SubscribeGroupToCalendar(calendarid string, userlist []string,
	success_user_file bool, success_user_name string, fail_user_file bool, fail_user_name string) bool {

	var successuserlist []string
	var failuserlist []string

	for _, u := range userlist {
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
		if success_user_name == "" {
			success_user_name = "success_user_" + calendarid + ".csv"
		}

		result := CreateOutputFile(success_user_name, successuserlist)
		if !result {
			log.Println("Error occured while trying to create success user file")
		}
	}

	// create successful user file
	if fail_user_file {
		// check and set name to default
		if fail_user_name == "" {
			fail_user_name = "fail_user_" + calendarid + ".csv"
		}

		result := CreateOutputFile(fail_user_name, failuserlist)
		if !result {
			log.Println("Error occured while trying to create fail user file")
		}
	}

	return true
}

func CreateOutputFile(filename string, userlist []string) bool {
	output_folder := "outputs"

	// Get the current working directory
	originalDir, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting current directory: %v\n", err)
		return false
	}

	// Check if the folder already exists
	if _, err := os.Stat(output_folder); os.IsNotExist(err) {
		// Folder does not exist, so create it
		err := os.Mkdir(output_folder, 0755)
		if err != nil {
			log.Printf("Error creating folder: %v\n", err)
			return false
		}
		log.Printf("Folder '%s' created successfully.\n", output_folder)
	} else {
		log.Printf("Folder '%s' already exists.\n", output_folder)
	}

	// Change the current working directory to the new folder
	err = os.Chdir(output_folder)
	if err != nil {
		log.Printf("Error changing directory: %v\n", err)
		return false
	}
	// Create the output file
	usercsv, err := os.Create(filename)
	if err != nil {
		log.Printf("Unable to create output file: %v", err)
	}
	defer usercsv.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(usercsv)
	defer writer.Flush()

	// Write header row
	header := []string{"NetID"}
	if err := writer.Write(header); err != nil {
		log.Printf("Error writing header: %v", err)
	}

	// Write data rows
	for _, netID := range userlist {
		record := []string{netID}
		if err := writer.Write(record); err != nil {
			log.Printf("Error writing record: %v", err)
		}
	}

	log.Printf("Successfully generate File: %s", filename)

	// Move back to the original directory
	err = os.Chdir(originalDir)
	if err != nil {
		log.Printf("Error changing back to original directory: %v\n", err)
		return false
	}

	return true
}
