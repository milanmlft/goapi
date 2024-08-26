package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "456EDF",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
	},
}

var mockFlavours = Flavours{
	{Name: "chocolate", Quantity: 42},
	{Name: "vanilla", Quantity: 10},
	{Name: "raspberry", Quantity: 23},
}

// Database methods
// GetUserLoginDetails returns a pointer to a LoginDetails object
func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	clientData := LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetFlavours() *Flavours {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	clientData := mockFlavours

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
