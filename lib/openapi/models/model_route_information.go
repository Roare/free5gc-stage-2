/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type RouteInformation struct {
	Ipv4Addr   string `json:"ipv4Addr,omitempty" bson:"ipv4Addr"`
	Ipv6Addr   string `json:"ipv6Addr,omitempty" bson:"ipv6Addr"`
	PortNumber int32  `json:"portNumber" bson:"portNumber"`
}
