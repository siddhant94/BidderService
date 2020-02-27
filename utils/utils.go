package utils


import (
	"fmt"
	"math/rand"
	"strconv"
	"github.com/siddhant94/BidderService/models"
	"github.com/siddhant94/BidderService/config"
)

// Map returns a new slice containing the results of applying the function f to each string in the original slice.
func Map(vs []*models.Bidder, f func() *models.Bidder) []*models.Bidder {
	vsm := make([]*models.Bidder, len(vs))
	for i, _ := range vs {
		vsm[i] = f()
	}
	return vsm
}

// PopulateBidder fills the data for input Bidder struct reference.
func PopulateBidder() *models.Bidder {
	const (
		b = 500 // max delay
		a = 10  // min delay
		portStart = 6001
	)
	bidderRef := new(models.Bidder)
	bidderRef.Delay = a + rand.Intn(b-a+1) // a ≤ n ≤ b
	bidderRef.Port = ":" + strconv.Itoa(portStart)
	config := config.GetAppConfig()
	bidderRef.RegistrationURL = config.BidderRegistrationAPI
	fmt.Println("%+v", bidderRef)
	return bidderRef
}