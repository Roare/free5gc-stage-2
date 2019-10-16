/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type UeInitiatedResourceRequest struct {
	PackFiltOp PacketFilterOperation `json:"packFiltOp" bson:"packFiltOp"`

	PackFiltInfo []PacketFilterInfo `json:"packFiltInfo" bson:"packFiltInfo"`

	ReqQos *RequestedQos `json:"reqQos,omitempty" bson:"reqQos"`
}
