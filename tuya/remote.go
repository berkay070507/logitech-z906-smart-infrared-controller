package tuya

import "time"

var lastLoginError = time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

func SendRemoteControl(code string) {
	// Perform request
	res, err := PerformRequest(
		"POST",
		"/v1.0/infrareds/"+IRDeviceID+"/remotes/"+AudioDeviceID+"/raw/command",
		[]byte(`{"raw_key": "`+code+`"}`),
	)

	// Check if response was successful
	if err != nil {
		// In case there is an authentication error, re-login and wait for user to retry
		if res.Code == 1010 {
			// Perform Login request
			GetToken()

			// Check last login time and update it for the next request
			lastLogin1minuteAgo := time.Now().After(lastLoginError.Add(1 * time.Minute))
			lastLoginError = time.Now()

			// Repeat the request if the last login intent was older than 1 minute (avoid infinite loop)
			if lastLogin1minuteAgo {
				SendRemoteControl(code)
			}
		}
		return
	}
}
