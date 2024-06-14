package controllers

import (
    "encoding/json"
    "net/http"
    "time"
    "strconv"
)

// OptionsContract represents the data structure of an options contract
type OptionsContract struct {
    Type          string  `json:"type"`          // call or put
    StrikePrice   float64 `json:"strike_price"`  // strike price of the option
    Bid           float64 `json:"bid"`           // bid price
    Ask           float64 `json:"ask"`           // ask price
    ExpirationDate string `json:"expiration_date"` // expiration date of the option
    LongShort     string  `json:"long_short"`    // long or short position
}

// AnalysisResponse represents the data structure of the analysis result
type AnalysisResponse struct {
    XYValues        []XYValue `json:"xy_values"`
    MaxProfit       float64   `json:"max_profit"`
    MaxLoss         float64   `json:"max_loss"`
    BreakEvenPoints []float64 `json:"break_even_points"`
}

// XYValue represents a pair of X and Y values
type XYValue struct {
    X float64 `json:"x"`
    Y float64 `json:"y"`
}

func AnalysisHandler(w http.ResponseWriter, r *http.Request) {
    var contracts []OptionsContract
    if err := json.NewDecoder(r.Body).Decode(&contracts); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    xyValues := calculateXYValues(contracts)
    maxProfit := calculateMaxProfit(contracts)
    maxLoss := calculateMaxLoss(contracts)
    breakEvenPoints := calculateBreakEvenPoints(contracts)

    response := AnalysisResponse{
        XYValues:        xyValues,
        MaxProfit:       maxProfit,
        MaxLoss:         maxLoss,
        BreakEvenPoints: breakEvenPoints,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func calculateXYValues(contracts []OptionsContract) []XYValue {
    // Implement the logic for calculating X and Y values
    var xyValues []XYValue
    // Example logic (placeholder)
    for _, contract := range contracts {
        // Logic to calculate values based on contract details
        xyValues = append(xyValues, XYValue{X: contract.StrikePrice, Y: 0}) // Placeholder calculation
    }
    return xyValues
}

func calculateMaxProfit(contracts []OptionsContract) float64 {
    maxProfit := 0.0
    for _, contract := range contracts {
        if contract.LongShort == "long" {
            if contract.Type == "call" {
                // For long calls, profit is unlimited
                maxProfit = math.Inf(1)
            } else if contract.Type == "put" {
                // For long puts, profit is limited to strike price minus premium paid
                premium := contract.Ask
                profit := contract.StrikePrice - premium
                if profit > maxProfit {
                    maxProfit = profit
                }
            }
        } else if contract.LongShort == "short" {
            if contract.Type == "call" {
                // For short calls, profit is limited to the premium received
                premium := contract.Bid
                if premium > maxProfit {
                    maxProfit = premium
                }
            } else if contract.Type == "put" {
                // For short puts, profit is limited to the premium received
                premium := contract.Bid
                if premium > maxProfit {
                    maxProfit = premium
                }
            }
        }
    }
    return maxProfit
}

func calculateMaxLoss(contracts []OptionsContract) float64 {
    maxLoss := 0.0
    for _, contract := range contracts {
        if contract.LongShort == "long" {
            if contract.Type == "call" {
                // For long calls, loss is limited to premium paid
                premium := contract.Ask
                if premium > maxLoss {
                    maxLoss = premium
                }
            } else if contract.Type == "put" {
                // For long puts, loss is limited to premium paid
                premium := contract.Ask
                if premium > maxLoss {
                    maxLoss = premium
                }
            }
        } else if contract.LongShort == "short" {
            if contract.Type == "call" {
                // For short calls, loss is unlimited
                maxLoss = math.Inf(1)
            } else if contract.Type == "put" {
                // For short puts, loss is limited to strike price minus premium received
                premium := contract.Bid
                loss := contract.StrikePrice - premium
                if loss > maxLoss {
                    maxLoss = loss
                }
            }
        }
    }
    return maxLoss
}

func calculateBreakEvenPoints(contracts []OptionsContract) []float64 {
    var breakEvenPoints []float64
    for _, contract := range contracts {
        if contract.LongShort == "long" {
            if contract.Type == "call" {
                // For long calls, break-even is strike price plus premium paid
                premium := contract.Ask
                breakEven := contract.StrikePrice + premium
                breakEvenPoints = append(breakEvenPoints, breakEven)
            } else if contract.Type == "put" {
                // For long puts, break-even is strike price minus premium paid
                premium := contract.Ask
                breakEven := contract.StrikePrice - premium
                breakEvenPoints = append(breakEvenPoints, breakEven)
            }
        } else if contract.LongShort == "short" {
            if contract.Type == "call" {
                // For short calls, break-even is strike price plus premium received
                premium := contract.Bid
                breakEven := contract.StrikePrice + premium
                breakEvenPoints = append(breakEvenPoints, breakEven)
            } else if contract.Type == "put" {
                // For short puts, break-even is strike price minus premium received
                premium := contract.Bid
                breakEven := contract.StrikePrice - premium
                breakEvenPoints = append(breakEvenPoints, breakEven)
            }
        }
    }
    return breakEvenPoints
}
