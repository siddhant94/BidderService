package main

/* SELF NOTE go mod init, go build, go mod tidy, go mod vendor - GO Modules project initialization*/
import (
	"fmt"
	"math/rand"
	"os"
	"github.com/siddhant94/auction-bidding-simulation/bidderService/models"
	"github.com/siddhant94/auction-bidding-simulation/bidderService/utils"
)

const auctioneerURL = "http://127.0.0.1:6000"
const bidderRegistrationAPI = "/register-bidder"

func main() {
	// Bidder Service params - delay-time, portToBind, URL (to register itself with auctioneer)
	// 1) Number of bidders and rest system generated. (JSON, YAML etc) config file, if user-input support for all three.
	biddersToCreate := os.Args[1]
	fmt.Println(bidders)
	rand.Seed(time.Now().UnixNano()) // To run only once at initialization

	var bidderList []*models.Bidder
	bidderList = make([]*models.Bidder, 5, 5)
	
	utils.BidderSliceMap(bidderList, utils.CreateBidder)
	
	= new(models.Bidder)
	n := a + rand.Intn(b-a+1) // a <= n <= b


}
