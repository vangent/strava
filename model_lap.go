/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package strava

import (
	"time"
)

type Lap struct {
	// The unique identifier of this lap
	Id int64 `json:"id,omitempty"`
	Activity *MetaActivity `json:"activity,omitempty"`
	Athlete *MetaAthlete `json:"athlete,omitempty"`
	// The lap's average cadence
	AverageCadence float32 `json:"average_cadence,omitempty"`
	// The lap's average speed
	AverageSpeed float32 `json:"average_speed,omitempty"`
	// The lap's distance, in meters
	Distance float32 `json:"distance,omitempty"`
	// The lap's elapsed time, in seconds
	ElapsedTime int32 `json:"elapsed_time,omitempty"`
	// The start index of this effort in its activity's stream
	StartIndex int32 `json:"start_index,omitempty"`
	// The end index of this effort in its activity's stream
	EndIndex int32 `json:"end_index,omitempty"`
	// The index of this lap in the activity it belongs to
	LapIndex int32 `json:"lap_index,omitempty"`
	// The maximum speed of this lat, in meters per second
	MaxSpeed float32 `json:"max_speed,omitempty"`
	// The lap's moving time, in seconds
	MovingTime int32 `json:"moving_time,omitempty"`
	// The name of the lap
	Name string `json:"name,omitempty"`
	// The athlete's pace zone during this lap
	PaceZone int32 `json:"pace_zone,omitempty"`
	Split int32 `json:"split,omitempty"`
	// The time at which the lap was started.
	StartDate time.Time `json:"start_date,omitempty"`
	// The time at which the lap was started in the local timezone.
	StartDateLocal time.Time `json:"start_date_local,omitempty"`
	// The elevation gain of this lap, in meters
	TotalElevationGain float32 `json:"total_elevation_gain,omitempty"`
}