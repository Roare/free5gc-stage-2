/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type RouteToLocation struct {
	Dnai        string            `json:"dnai" bson:"dnai"`
	RouteInfo   *RouteInformation `json:"routeInfo,omitempty" bson:"routeInfo"`
	RouteProfId string            `json:"routeProfId,omitempty" bson:"routeProfId"`
}