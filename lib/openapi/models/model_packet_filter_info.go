/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type PacketFilterInfo struct {

	// An identifier of packet filter.
	PackFiltId string `json:"packFiltId,omitempty" bson:"packFiltId"`

	// The requested order for the PCC rule generated fromt the packet fitler information.
	Precedence int32 `json:"precedence,omitempty" bson:"precedence"`

	// Defines a packet filter for an IP flow.Refer to subclause 5.3.54 of 3GPP TS 29.212 [23] for encoding.
	PackFiltCont string `json:"packFiltCont,omitempty" bson:"packFiltCont"`

	// Contains the Ipv4 Type-of-Service and mask field or the Ipv6 Traffic-Class field and mask field.
	TosTrafficClass string `json:"tosTrafficClass,omitempty" bson:"tosTrafficClass"`

	// The security parameter index of the IPSec packet.
	Spi string `json:"spi,omitempty" bson:"spi"`

	// The Ipv6 flow label header field.
	FlowLabel string `json:"flowLabel,omitempty" bson:"flowLabel"`

	FlowDirection FlowDirection `json:"flowDirection,omitempty" bson:"flowDirection"`
}
