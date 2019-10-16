/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type RuleReport struct {

	// Contains the identifier of the affected PCC rule(s).
	PccRuleIds []string `json:"pccRuleIds" bson:"pccRuleIds"`

	RuleStatus RuleStatus `json:"ruleStatus" bson:"ruleStatus"`

	// Indicates the version of a PCC rule.
	ContVers []int32 `json:"contVers,omitempty" bson:"contVers"`

	FailureCode FailureCode `json:"failureCode" bson:"failureCode"`

	FinUnitAct FinalUnitAction `json:"finUnitAct,omitempty" bson:"finUnitAct"`

	// indicates the RAN or NAS release cause code information.
	RanNasRelCauses []RanNasRelCause `json:"ranNasRelCauses,omitempty" bson:"ranNasRelCauses"`
}
