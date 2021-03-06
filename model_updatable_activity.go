/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package strava

type UpdatableActivity struct {
	// Whether this activity is a commute
	Commute bool `json:"commute"`
	// Whether this activity was recorded on a training machine
	Trainer bool `json:"trainer"`
	// The description of the activity
	Description string `json:"description"`
	// The name of the activity
	Name string `json:"name"`
	// The type of the activity
	Type_ *ActivityType `json:"type"`
	// Identifier for the gear associated with the activity. ‘none’ clears gear from activity
	GearId string `json:"gear_id,omitempty"`
	// The workout type for the activity
	WorkoutType int `json:"workout_type"`
}
