package main

/* SELF NOTE go mod init, go build, go mod tidy, go mod vendor - GO Modules project initialization*/
import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"log"
	"github.com/siddhant94/BidderService/config"
	"github.com/siddhant94/BidderService/models"
	"github.com/siddhant94/BidderService/utils"
)
var defaultBidders int

func init() {
	rand.Seed(time.Now().UnixNano()) // Initialize application wide Seed for rand
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	// Bidder Service params - delay-time, portToBind, URL (to register itself with auctioneer)
	// 1) Number of bidders and rest system generated. (JSON, YAML etc) config file, if user-input support for all three.
	var biddersToCreate string
	if len(os.Args) > 1 {
		biddersToCreate = os.Args[1]
	} else {
		biddersToCreate = config.GetAppConfig().DefaultBidders
	}
	fmt.Println(biddersToCreate)

	var bidderList []*models.Bidder
	bidderList = make([]*models.Bidder, 5, 5)
	log.Println(bidderList)
	
	utils.Map(bidderList, utils.PopulateBidder)
}
