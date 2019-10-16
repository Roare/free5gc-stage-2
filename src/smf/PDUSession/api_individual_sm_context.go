/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package PDUSession

import (
	"github.com/gin-gonic/gin"
	"free5gc/lib/http_wrapper"
	"free5gc/lib/openapi"
	"free5gc/lib/openapi/models"
	"free5gc/src/smf/smf_handler/smf_message"
	"log"
	"net/http"
	"strings"
)

// ReleaseSmContext - Release SM Context
func ReleaseSmContext(c *gin.Context) {
	var request models.ReleaseSmContextRequest
	request.JsonData = new(models.SmContextReleaseData)

	s := strings.Split(c.GetHeader("Content-Type"), ";")
	var err error
	switch s[0] {
	case "application/json":
		err = c.ShouldBindJSON(request.JsonData)
	case "multipart/related":
		err = c.ShouldBindWith(&request, openapi.MultipartRelatedBinding{})
	}
	if err != nil {
		log.Print(err)
		return
	}

	req := http_wrapper.NewRequest(c.Request, request)
	req.Params["smContextRef"] = c.Params.ByName("smContextRef")

	message := smf_message.NewHandlerMessage(smf_message.PDUSessionSMContextRelease, req)
	smf_message.SendMessage(message)

	_ = <-message.ResponseChan

	c.Status(http.StatusNoContent)

}

// RetrieveSmContext - Retrieve SM Context
func RetrieveSmContext(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateSmContext - Update SM Context
func UpdateSmContext(c *gin.Context) {
	var request models.UpdateSmContextRequest
	request.JsonData = new(models.SmContextUpdateData)

	s := strings.Split(c.GetHeader("Content-Type"), ";")
	var err error
	switch s[0] {
	case "application/json":
		err = c.ShouldBindJSON(request.JsonData)
	case "multipart/related":
		err = c.ShouldBindWith(&request, openapi.MultipartRelatedBinding{})
	}
	if err != nil {
		log.Print(err)
		return
	}

	req := http_wrapper.NewRequest(c.Request, request)
	req.Params["smContextRef"] = c.Params.ByName("smContextRef")

	message := smf_message.NewHandlerMessage(smf_message.PDUSessionSMContextUpdate, req)
	smf_message.SendMessage(message)

	rsp := <-message.ResponseChan

	HTTPResponse := rsp.HTTPResponse
	if HTTPResponse.Status < 300 {
		c.Render(HTTPResponse.Status, openapi.MultipartRelatedRender{Data: HTTPResponse.Body})
	} else {
		c.JSON(HTTPResponse.Status, HTTPResponse.Body)
	}
}