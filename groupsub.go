package calendarmod

type GroupSubResult struct {
	successuserlist []string
	failuserlist    []string
}

// Subscribe group of user to group of calendar
//
//	@param {string[]} userlist - list of valid Google emails of targeted users, in email format
//	@param {string} calendarid - id of the calendar for the users to subscribe
//
//	@return {GroupSubResult} - result of group calendar subscription, with a list of success users and a list of failed user cases
func (c *CalendarClient) SubscribeGroupToCalendar(calendarid string, userlist []string) GroupSubResult {

	var gr GroupSubResult

	for _, u := range userlist {
		res := c.SubscribeUserToCalendar(u, calendarid)
		if res {
			gr.successuserlist = append(gr.successuserlist, u)
		} else {
			gr.failuserlist = append(gr.failuserlist, u)
		}
	}

	return gr
}
