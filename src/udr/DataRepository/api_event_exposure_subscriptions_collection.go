/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DataRepository

import (
	"github.com/gin-gonic/gin"
	"free5gc/lib/http_wrapper"
	"free5gc/lib/openapi/models"
	"free5gc/src/udr/logger"
	"free5gc/src/udr/udr_handler/udr_message"
)

// CreateEeSubscriptions - Create individual EE subscription
func CreateEeSubscriptions(c *gin.Context) {
	var eeSubscription models.EeSubscription
	if err := c.ShouldBindJSON(&eeSubscription); err != nil {
		logger.DataRepoLog.Panic(err.Error())
	}

	req := http_wrapper.NewRequest(c.Request, eeSubscription)
	req.Params["ueId"] = c.Params.ByName("ueId")

	handlerMsg := udr_message.NewHandlerMessage(udr_message.EventCreateEeSubscriptions, req)
	udr_message.SendMessage(handlerMsg)

	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}

// Queryeesubscriptions - Retrieves the ee subscriptions of a UE
func Queryeesubscriptions(c *gin.Context) {
	req := http_wrapper.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")

	handlerMsg := udr_message.NewHandlerMessage(udr_message.EventQueryeesubscriptions, req)
	udr_message.SendMessage(handlerMsg)

	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}
