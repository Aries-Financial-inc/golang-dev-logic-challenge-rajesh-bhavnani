package tests

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "encoding/json"
    "bytes"
    "github.com/stretchr/testify/assert"
    "github.com/Aries-Financial-inc/golang-dev-logic-challenge-rajesh-bhavnani/controllers"
    "github.com/Aries-Financial-inc/golang-dev-logic-challenge-rajesh-bhavnani/models"
)

func TestOptionsContractModelValidation(t *testing.T) {
    // Example test case for validating the OptionsContract model
    contract := models.OptionsContract{
        Type: "call",
        StrikePrice: 100.0,
        Bid: 5.0,
        Ask: 6.0,
        ExpirationDate: "2024-12-31",
        LongShort: "long",
    }
    
    assert.Equal(t, "call", contract.Type)
    assert.Equal(t, 100.0, contract.StrikePrice)
    assert.Equal(t, 5.0, contract.Bid)
    assert.Equal(t, 6.0, contract.Ask)
    assert.Equal(t, "2024-12-31", contract.ExpirationDate)
    assert.Equal(t, "long", contract.LongShort)
}

func TestAnalysisEndpoint(t *testing.T) {
    contracts := []models.OptionsContract{
        {
            Type: "call",
            StrikePrice: 100.0,
            Bid: 5.0,
            Ask: 6.0,
            ExpirationDate: "2024-12-31",
            LongShort: "long",
        },
    }
    
    requestBody, err := json.Marshal(contracts)
    assert.NoError(t, err)

    req, err := http.NewRequest("POST", "/analyze", bytes.NewBuffer(requestBody))
    assert.NoError(t, err)
    
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(controllers.AnalysisHandler)
    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    var response controllers.AnalysisResponse
    err = json.NewDecoder(rr.Body).Decode(&response)
    assert.NoError(t, err)

    // Check that the response contains the expected fields
    assert.NotNil(t, response.XYValues)
    assert.NotNil(t, response.MaxProfit)
    assert.NotNil(t, response.MaxLoss)
    assert.NotNil(t, response.BreakEvenPoints)
}

func TestIntegration(t *testing.T) {
    contracts := []models.OptionsContract{
        {
            Type: "call",
            StrikePrice: 100.0,
            Bid: 5.0,
            Ask: 6.0,
            ExpirationDate: "2024-12-31",
            LongShort: "long",
        },
        {
            Type: "put",
            StrikePrice: 90.0,
            Bid: 3.0,
            Ask: 4.0,
            ExpirationDate: "2024-12-31",
            LongShort: "short",
        },
    }

    requestBody, err := json.Marshal(contracts)
    assert.NoError(t, err)

    req, err := http.NewRequest("POST", "/analyze", bytes.NewBuffer(requestBody))
    assert.NoError(t, err)

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(controllers.AnalysisHandler)
    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    var response controllers.AnalysisResponse
    err = json.NewDecoder(rr.Body).Decode(&response)
    assert.NoError(t, err)

    // Check that the response contains the expected fields
    assert.NotNil(t, response.XYValues)
    assert.NotNil(t, response.MaxProfit)
    assert.NotNil(t, response.MaxLoss)
    assert.NotNil(t, response.BreakEvenPoints)
}
