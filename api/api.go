package api

type IceCreamParams struct {
	Username string
}

// Ice Cream API response
type IceCreamResponse struct {
	Flavours []string
	Code     int
}

// Error response
type Error struct {
	Message string
	Code    int
}
