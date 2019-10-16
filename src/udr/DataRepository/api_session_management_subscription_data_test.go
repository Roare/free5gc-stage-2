/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package DataRepository_test

import (
	"context"
	"free5gc/src/udr/logger"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson"
	"free5gc/lib/Nudr_DataRepository"
	"free5gc/lib/openapi/models"
)

// QuerySmData - Retrieves the Session Management subscription data of a UE
func TestQuerySmData(t *testing.T) {
	runTestServer(t)

	connectMongoDB(t)

	// Drop old data
	collection := Client.Database("free5gc").Collection("subscriptionData.provisionedData.smData")
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789", "servingPlmnId": "20893"})

	// Set client and set url
	client := setTestClient(t)

	// Set test data
	ueId := "imsi-0123456789"
	servingPlmnId := "20893"
	{
		testData := models.SessionManagementSubscriptionData{
			SharedDnnConfigurationsIds: "1",
		}
		insertTestData := toBsonM(testData)
		insertTestData["ueId"] = ueId
		insertTestData["servingPlmnId"] = servingPlmnId
		collection.InsertOne(context.TODO(), insertTestData)
	}
	{
		testData := models.SessionManagementSubscriptionData{
			SharedDnnConfigurationsIds: "2",
		}
		insertTestData := toBsonM(testData)
		insertTestData["ueId"] = ueId
		insertTestData["servingPlmnId"] = servingPlmnId
		collection.InsertOne(context.TODO(), insertTestData)
	}

	testData := []models.SessionManagementSubscriptionData{
		{
			SharedDnnConfigurationsIds: "1",
		},
		{
			SharedDnnConfigurationsIds: "2",
		},
	}

	{
		// Check test data (Use RESTful GET)
		var querySmDataParamOpts Nudr_DataRepository.QuerySmDataParamOpts
		sessionManagementSubscriptionDatas, res, err := client.SessionManagementSubscriptionDataApi.QuerySmData(context.TODO(), ueId, servingPlmnId, &querySmDataParamOpts)
		if err != nil {
			logger.AppLog.Panic(err)
		}

		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if cmp.Equal(testData, sessionManagementSubscriptionDatas, Opt) != true {
			t.Errorf("handler returned unexpected body: got %v want %v",
				sessionManagementSubscriptionDatas, testData)
		}
	}

	// Clean test data
	collection.DeleteOne(context.TODO(), bson.M{"ueId": "imsi-0123456789", "servingPlmnId": "20893"})

	// TEST END
}
