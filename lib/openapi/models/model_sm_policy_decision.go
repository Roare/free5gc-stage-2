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

type SmPolicyDecision struct {

	// A map of Sessionrules with the content being the SessionRule as described in subclause 5.6.2.7.
	SessRules map[string]SessionRule `json:"sessRules,omitempty" bson:"sessRules"`

	// A map of PCC rules with the content being the PCCRule as described in subclause 5.6.2.6.
	PccRules map[string]PccRule `json:"pccRules,omitempty" bson:"pccRules"`

	// If it is included and set to true, it indicates the P-CSCF Restoration is requested.
	PcscfRestIndication bool `json:"pcscfRestIndication,omitempty" bson:"pcscfRestIndication"`

	// Map of QoS data policy decisions.
	QosDecs map[string]QosData `json:"qosDecs,omitempty" bson:"qosDecs"`

	// Map of Charging data policy decisions.
	ChgDecs map[string]ChargingData `json:"chgDecs,omitempty" bson:"chgDecs"`

	ChargingInfo *ChargingInformation `json:"chargingInfo,omitempty" bson:"chargingInfo"`

	// Map of Traffic Control data policy decisions.
	TraffContDecs map[string]TrafficControlData `json:"traffContDecs,omitempty" bson:"traffContDecs"`

	// Map of Usage Monitoring data policy decisions.
	UmDecs map[string]UsageMonitoringData `json:"umDecs,omitempty" bson:"umDecs"`

	// Map of QoS characteristics for non standard 5QIs. This map uses the 5QI values as keys.
	QosChars map[string]QosCharacteristics `json:"qosChars,omitempty" bson:"qosChars"`

	ReflectiveQoSTimer int32 `json:"reflectiveQoSTimer,omitempty" bson:"reflectiveQoSTimer"`

	// A map of condition data with the content being as described in subclause 5.6.2.9.
	Conds map[string]ConditionData `json:"conds,omitempty" bson:"conds"`

	RevalidationTime *time.Time `json:"revalidationTime,omitempty" bson:"revalidationTime"`

	// Indicates the offline charging is applicable to the PDU session or PCC rule.
	Offline bool `json:"offline,omitempty" bson:"offline"`

	// Indicates the online charging is applicable to the PDU session or PCC rule.
	Online bool `json:"online,omitempty" bson:"online"`

	// Defines the policy control request triggers subscribed by the PCF.
	PolicyCtrlReqTriggers []PolicyControlRequestTrigger `json:"policyCtrlReqTriggers,omitempty" bson:"policyCtrlReqTriggers"`

	// Defines the last list of rule control data requested by the PCF.
	LastReqRuleData []RequestedRuleData `json:"lastReqRuleData,omitempty" bson:"lastReqRuleData"`

	LastReqUsageData *RequestedUsageData `json:"lastReqUsageData,omitempty" bson:"lastReqUsageData"`

	// Map of PRA information.
	PraInfos map[string]PresenceInfoRm `json:"praInfos,omitempty" bson:"praInfos"`

	Ipv4Index int32 `json:"ipv4Index,omitempty" bson:"ipv4Index"`

	Ipv6Index int32 `json:"ipv6Index,omitempty" bson:"ipv6Index"`

	QosFlowUsage QosFlowUsage `json:"qosFlowUsage,omitempty" bson:"qosFlowUsage"`

	SuppFeat string `json:"suppFeat,omitempty" bson:"suppFeat"`
}
