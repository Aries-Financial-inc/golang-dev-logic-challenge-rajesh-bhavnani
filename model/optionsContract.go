package model

type OptionsContract struct {
    Type          string  `json:"type"`          // call or put
    StrikePrice   float64 `json:"strike_price"`  // strike price of the option
    Bid           float64 `json:"bid"`           // bid price
    Ask           float64 `json:"ask"`           // ask price
    ExpirationDate string `json:"expiration_date"` // expiration date of the option
    LongShort     string  `json:"long_short"`    // long or short position
}
