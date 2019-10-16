/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type TrafficControlData struct {

	// Univocally identifies the traffic control policy data within a PDU session.
	TcId string `json:"tcId" bson:"tcId"`

	FlowStatus FlowStatus `json:"flowStatus,omitempty" bson:"flowStatus"`

	RedirectInfo *RedirectInformation `json:"redirectInfo,omitempty" bson:"redirectInfo"`

	// Indicates whether applicat'on's start or stop notification is to be muted.
	MuteNotif bool `json:"muteNotif,omitempty" bson:"muteNotif"`

	// Reference to a pre-configured traffic steering policy for downlink traffic at the SMF.
	TrafficSteeringPolIdDl string `json:"trafficSteeringPolIdDl,omitempty" bson:"trafficSteeringPolIdDl"`

	// Reference to a pre-configured traffic steering policy for uplink traffic at the SMF.
	TrafficSteeringPolIdUl string `json:"trafficSteeringPolIdUl,omitempty" bson:"trafficSteeringPolIdUl"`

	// A list of location which the traffic shall be routed to for the AF request
	RouteToLocs []RouteToLocation `json:"routeToLocs,omitempty" bson:"routeToLocs"`

	UpPathChgEvent *UpPathChgEvent `json:"upPathChgEvent,omitempty" bson:"upPathChgEvent"`
}
