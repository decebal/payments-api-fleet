// +build integration

package api_test

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("ListPayments", func() {
	// var (
	// 	router        *gin.Engine
	// 	requestResult *httptest.ResponseRecorder
	// )
	//
	// router = gin.Default()
	// router = api.InitRoutes(router)
	// go router.Run(port)
	//
	// Describe("Testing Route Endpoints", func() {
	// 	Context(baseURL+"/api/v1/healthcheck", func() {
	//
	// 		requestResult = sendRequest(router, "GET", baseURL+"/api/v1/healthcheck", nil)
	//
	// 		It("should be HTTP Status OK", func() {
	// 			Expect(requestResult.Code).To(Equal(http.StatusOK))
	// 		})
	//
	// 		It("should return JSON containing 'status': 'ok'", func() {
	// 			bodyStr := requestResult.Body.String()
	// 			Expect(bodyStr).To(MatchRegexp(".*status[^a-zA-Z]+ok*"))
	// 		})
	// 	})
	// })
})
