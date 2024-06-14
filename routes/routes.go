package routes

import (
	"net/http"
    "github.com/Aries-Financial-inc/golang-dev-logic-challenge-rajesh-bhavnani/controllers"

	"github.com/gin-gonic/gin"
)

// OptionsContract structure for the request body
type OptionsContract struct {
    Type          string  `json:"type"`          // call or put
    StrikePrice   float64 `json:"strike_price"`  // strike price of the option
    Bid           float64 `json:"bid"`           // bid price
    Ask           float64 `json:"ask"`           // ask price
    ExpirationDate string `json:"expiration_date"` // expiration date of the option
    LongShort     string  `json:"long_short"`    // long or short position
}

// AnalysisResult structure for the response body
type AnalysisResult struct {
	GraphData       []GraphPoint `json:"graph_data"`
	MaxProfit       float64      `json:"max_profit"`
	MaxLoss         float64      `json:"max_loss"`
	BreakEvenPoints []float64    `json:"break_even_points"`
}

// GraphPoint structure for X & Y values of the risk & reward graph
type GraphPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/analyze", func(c *gin.Context) {
        var contracts []OptionsContract

        if err := c.ShouldBindJSON(&contracts); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Convert to controllers.OptionsContract
        var controllerContracts []controllers.OptionsContract
        for _, contract := range contracts {
            controllerContracts = append(controllerContracts, controllers.OptionsContract{
                Type: contract.Type,
                StrikePrice: contract.StrikePrice,
                Bid: contract.Bid,
                Ask: contract.Ask,
                ExpirationDate: contract.ExpirationDate,
                LongShort: contract.LongShort,
            })
        }

        xyValues := controllers.CalculateXYValues(controllerContracts)
        maxProfit := controllers.CalculateMaxProfit(controllerContracts)
        maxLoss := controllers.CalculateMaxLoss(controllerContracts)
        breakEvenPoints := controllers.CalculateBreakEvenPoints(controllerContracts)

        var graphData []GraphPoint
        for _, xy := range xyValues {
            graphData = append(graphData, GraphPoint{X: xy.X, Y: xy.Y})
        }

        result := AnalysisResult{
            GraphData:       graphData,
            MaxProfit:       maxProfit,
            MaxLoss:         maxLoss,
            BreakEvenPoints: breakEvenPoints,
        }

        c.JSON(http.StatusOK, result)
    })

	return router
}
