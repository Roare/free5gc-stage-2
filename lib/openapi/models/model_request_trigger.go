/*
 * Npcf_AMPolicyControl
 *
 * Access and Mobility Policy Control Service API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// RequestTrigger : Possible values are - LOC_CH: Location change (tracking area). The tracking area of the UE has changed. - PRA_CH: Change of UE presence in PRA. The UE is entering/leaving a Presence Reporting Area. - SERV_AREA_CH: Service Area Restriction change. The UDM notifies the AMF that the subscribed service area restriction information has changed. - RFSP_CH: RFSP index change. The UDM notifies the AMF that the subscribed RFSP index has changed.
type RequestTrigger string

// List of RequestTrigger
const (
	LOC_CHRequestTrigger           RequestTrigger = "LOC_CH"
	PRA_CHRequestTrigger           RequestTrigger = "PRA_CH"
	SERV_AREA_CHRequestTrigger     RequestTrigger = "SERV_AREA_CH"
	RFSP_CHRequestTrigger          RequestTrigger = "RFSP_CH"
	ALLOWED_NSSAI_CHRequestTrigger RequestTrigger = "ALLOWED_NSSAI_CH"
)
