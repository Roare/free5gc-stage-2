/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package SubscriberDataManagement_test

import (
	"testing"
)

// SubscribeToSharedData - subscribe to notifications for shared data
func TestSubscribeToSharedData(t *testing.T) {

	/*	go func() { // udm server
				router := gin.Default()
				Nudm_SDM_Server.AddService(router)

				udmLogPath := path_util.Gofree5gcPath("free5gc/udmsslkey.log")
				udmPemPath := path_util.Gofree5gcPath("free5gc/support/TLS/udm.pem")
				udmKeyPath := path_util.Gofree5gcPath("free5gc/support/TLS/udm.key")

				server, err := http2_util.NewServer(":29503", udmLogPath, router)
				if err == nil && server != nil {
					logger.InitLog.Infoln(server.ListenAndServeTLS(udmPemPath, udmKeyPath))
					assert.True(t, err == nil)
				}
			}()

			udm_util.testInitUdmConfig()
		go udm_handler.Handle()

			go func() { // fake udr server
				router := gin.Default()

				router.POST("/nudr-dr/v1/:shared-data-subscriptions", func(c *gin.Context) {
					ueId := c.Param("ueId")
					fmt.Println("FFFFFF")
					fmt.Println("==========SubscribeToSharedData - subscribe to notifications for shared data==========")
					fmt.Println("ueId: ", ueId)

					var SdmSubscription models.SdmSubscription
					if err := c.ShouldBindJSON(&SdmSubscription); err != nil {
						fmt.Println("fake udr server error")
						c.JSON(http.StatusInternalServerError, gin.H{})
						return
					}
					fmt.Println("SdmSubscription - ", SdmSubscription.NfInstanceId)
					c.JSON(http.StatusCreated, gin.H{})
				})

				udrLogPath := path_util.Gofree5gcPath("free5gc/udrsslkey.log")
				udrPemPath := path_util.Gofree5gcPath("free5gc/support/TLS/udr.pem")
				udrKeyPath := path_util.Gofree5gcPath("free5gc/support/TLS/udr.key")

				server, err := http2_util.NewServer(":29504", udrLogPath, router)
				if err == nil && server != nil {
					logger.InitLog.Infoln(server.ListenAndServeTLS(udrPemPath, udrKeyPath))
					assert.True(t, err == nil)
				}
			}()

			udm_context.Init()
			cfg := Nudm_SDM_Client.NewConfiguration()
			cfg.SetBasePath("https://localhost:29503")
			clientAPI := Nudm_SDM_Client.NewAPIClient(cfg)

			var sdmSubscription models.SdmSubscription
			sdmSubscription.NfInstanceId = "Test_NfinstanceId"

			_, resp, err := clientAPI.SubscriptionCreationForSharedDataApi.SubscribeToSharedData(context.TODO(), sdmSubscription)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("resp: ", resp)
			}
	*/
}
