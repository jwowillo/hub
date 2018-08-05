package main

import "time"

// Action to perform
type Action func()

// Schedule the Action to be performed every time.Duration hours.
//
// The entire duration is waited for before performing the action for the first
// time.
func Schedule(action Action, interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval * time.Hour)
			action()
		}
	}()
}
