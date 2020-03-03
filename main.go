package main

/* SELF NOTE go mod init, go build, go mod tidy, go mod vendor - GO Modules project initialization*/
import (
	"math/rand"
	"os"
	"time"
	"log"
	"strconv"
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
	var (
		numOfBidders int
		err error
	)
	if len(os.Args) > 1 {
		numOfBidders, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Println("Unable to convert argument as int. Picking up default value from config.")
		}
	}
	if numOfBidders == 0 { // Since zero value of int is 0
		numOfBidders, err = strconv.Atoi(config.GetAppConfig().DefaultBidders)
		if err != nil {
			log.Println("Error: string to int for default value in config. Unable to proceed")
			return
		}
	}

	var bidderList []*models.Bidder
	bidderList = make([]*models.Bidder, numOfBidders)
	
	log.Println(numOfBidders)
	bidderList = utils.Map(bidderList, utils.PopulateBidder)
	utils.BidderService(bidderList)
	log.Println("main finished")
	// select {
	// 	case signal, ok := <-biddersSignalChan:
	// 		if !ok {

	// 		}
	// }
}