package config


type Configuration struct {
	AuctioneerURL  string
	BidderRegistrationAPI string
	DefaultBidders string
}

var AppConfig = new(Configuration)

func init() {
	AppConfig.AuctioneerURL = "http://127.0.0.1:6000"
	AppConfig.BidderRegistrationAPI = "/register-bidder"
	AppConfig.DefaultBidders = "5"
}

func GetAppConfig() *Configuration{
	return AppConfig
}