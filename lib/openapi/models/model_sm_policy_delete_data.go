/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type SmPolicyDeleteData struct {
	UserLocationInfo *UserLocation `json:"userLocationInfo,omitempty" bson:"userLocationInfo"`

	UeTimeZone string `json:"ueTimeZone,omitempty" bson:"ueTimeZone"`

	ServingNetwork *NetworkId `json:"servingNetwork,omitempty" bson:"servingNetwork"`

	UserLocationInfoTime *time.Time `json:"userLocationInfoTime,omitempty" bson:"userLocationInfoTime"`

	// Contains the RAN and/or NAS release cause.
	RanNasRelCauses []RanNasRelCause `json:"ranNasRelCauses,omitempty" bson:"ranNasRelCauses"`

	// Contains the usage report
	AccuUsageReports []AccuUsageReport `json:"accuUsageReports,omitempty" bson:"accuUsageReports"`
}
