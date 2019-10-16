/*
 * NudmUEAU
 *
 * UDM UE Authentication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package UEAuthentication

import (
	"free5gc/lib/http_wrapper"
	"free5gc/lib/openapi/models"
	"free5gc/src/udm/logger"
	"free5gc/src/udm/udm_handler"
	"free5gc/src/udm/udm_handler/udm_message"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ConfirmAuth - Create a new confirmation event
func ConfirmAuth(c *gin.Context) {
	var authEvent models.AuthEvent
	err := c.ShouldBindJSON(&authEvent)
	if err != nil {
		logger.UeauLog.Errorln(err)
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := http_wrapper.NewRequest(c.Request, authEvent)
	req.Params["supi"] = c.Params.ByName("supi")

	handlerMsg := udm_message.NewHandlerMessage(udm_message.EventConfirmAuth, req)
	udm_handler.SendMessage(handlerMsg)
	rsp := <-handlerMsg.ResponseChan

	HTTPResponse := rsp.HTTPResponse
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}
