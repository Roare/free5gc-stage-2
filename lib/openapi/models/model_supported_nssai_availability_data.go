/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SupportedNssaiAvailabilityData struct {
	Tai *Tai `json:"tai" bson:"tai" yaml:"tai"`

	SupportedSnssaiList []Snssai `json:"supportedSnssaiList" bson:"supportedSnssaiList" yaml:"supportedSnssaiList"`
}