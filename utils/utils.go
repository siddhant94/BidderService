package utils

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
	"strconv"
	"sync"

	"github.com/siddhant94/BidderService/config"
	"github.com/siddhant94/BidderService/models"
)

var portSequence func() int
func init() {
	portSequence = getPortNumber()
}

// Map returns a new slice containing the results of applying the function f to each string in the original slice.
func Map(vs []*models.Bidder, f func() *models.Bidder) []*models.Bidder {
	vsm := make([]*models.Bidder, len(vs))
	for i := range vs {
		vsm[i] = f()
	}
	return vsm
}

func getPortNumber() func() int {
	i := 6000 // So that it would start from 6001, covert to const or config
	return func() int {
		i++
		return i
	}
}

// PopulateBidder fills the data for input Bidder struct reference.
func PopulateBidder() *models.Bidder {
	const (
		b         = 500 // max delay
		a         = 10  // min delay
	)
	// portSequence := getPortNumber()
	bidderRef := new(models.Bidder)
	bidderRef.Delay = a + rand.Intn(b-a+1) // a ≤ n ≤ b
	bidderRef.Port = ":" + strconv.Itoa(portSequence())
	config := config.GetAppConfig()
	bidderRef.RegistrationURL = config.BidderRegistrationAPI
	return bidderRef
}

var wg sync.WaitGroup

// BidderService takes list of bidders, starts server and signals to biddersSignalChan
func BidderService(biddersList []*models.Bidder) {
	fmt.Println("In bidder service")
	// We use channel for signalling for 1) staring of bidder server, 2) Error encountered
	// Wait for biddersSignalChan to fill upto buffer. 1) Can use wait groups 2) Blocking select
	biddersSignalChan := make(chan error, len(biddersList))
	for _, val := range biddersList {
		serverConf := getBidderServer(val.Port)
		wg.Add(1)
		startServer(serverConf, biddersSignalChan)
	}
	wg.Wait()
	close(biddersSignalChan)
	// GOTCHA: Don't range over channel without closing as then channel remains open-ended i.e. range would not know when to end
	var failed []error
	for elem := range biddersSignalChan {
		fmt.Printf("%+v\n", elem)
		if elem != nil {
			failed = append(failed, elem)
		}
	}
	if len(failed) > 0 {
		log.Printf("Bidder server creation failed for : %d\n", len(failed))
	}
}

func startServer(serverConf *http.Server, biddersSignalChan chan<- error) {
	defer wg.Done()
	fmt.Println("In start server")
	go func() {
		if err := serverConf.ListenAndServe(); err != nil {
			log.Println(err)
			biddersSignalChan <- err
		}
	}()
	biddersSignalChan <- nil
	return
}

func getBidderServer(port string) *http.Server {
	router := http.NewServeMux()
	router.HandleFunc("/", bidderServerHandler)

	server := &http.Server{
		Addr:         "127.0.0.1" + port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return server
}
// Message is response for default bidder server handler
type Message struct {
    Name string
    Body string
    Time int64
}

func bidderServerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// m := Message{"Alice", "Hello", 1294706395881547000}
	b := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	w.Write(b)
}