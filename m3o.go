package m3o

import (
	"github.com/m3o/m3o-go/service/address"
	"github.com/m3o/m3o-go/service/answer"
	"github.com/m3o/m3o-go/service/cache"
	"github.com/m3o/m3o-go/service/crypto"
	"github.com/m3o/m3o-go/service/currency"
	"github.com/m3o/m3o-go/service/db"
	"github.com/m3o/m3o-go/service/email"
	"github.com/m3o/m3o-go/service/emoji"
	"github.com/m3o/m3o-go/service/file"
	"github.com/m3o/m3o-go/service/forex"
	"github.com/m3o/m3o-go/service/geocoding"
	"github.com/m3o/m3o-go/service/helloworld"
	"github.com/m3o/m3o-go/service/id"
	"github.com/m3o/m3o-go/service/image"
	"github.com/m3o/m3o-go/service/ip"
	"github.com/m3o/m3o-go/service/location"
	"github.com/m3o/m3o-go/service/otp"
	"github.com/m3o/m3o-go/service/postcode"
	"github.com/m3o/m3o-go/service/routing"
	"github.com/m3o/m3o-go/service/rss"
	"github.com/m3o/m3o-go/service/sentiment"
	"github.com/m3o/m3o-go/service/sms"
	"github.com/m3o/m3o-go/service/stock"
	"github.com/m3o/m3o-go/service/stream"
	"github.com/m3o/m3o-go/service/thumbnail"
	"github.com/m3o/m3o-go/service/time"
	"github.com/m3o/m3o-go/service/url"
	"github.com/m3o/m3o-go/service/user"
	"github.com/m3o/m3o-go/service/weather"
)

func NewClient(token string) *Client {
	return &Client{
		token: token,

		AddressService:    address.NewAddressService(token),
		AnswerService:     answer.NewAnswerService(token),
		CacheService:      cache.NewCacheService(token),
		CryptoService:     crypto.NewCryptoService(token),
		CurrencyService:   currency.NewCurrencyService(token),
		DbService:         db.NewDbService(token),
		EmailService:      email.NewEmailService(token),
		EmojiService:      emoji.NewEmojiService(token),
		FileService:       file.NewFileService(token),
		ForexService:      forex.NewForexService(token),
		GeocodingService:  geocoding.NewGeocodingService(token),
		HelloworldService: helloworld.NewHelloworldService(token),
		IdService:         id.NewIdService(token),
		ImageService:      image.NewImageService(token),
		IpService:         ip.NewIpService(token),
		LocationService:   location.NewLocationService(token),
		OtpService:        otp.NewOtpService(token),
		PostcodeService:   postcode.NewPostcodeService(token),
		RoutingService:    routing.NewRoutingService(token),
		RssService:        rss.NewRssService(token),
		SentimentService:  sentiment.NewSentimentService(token),
		SmsService:        sms.NewSmsService(token),
		StockService:      stock.NewStockService(token),
		StreamService:     stream.NewStreamService(token),
		ThumbnailService:  thumbnail.NewThumbnailService(token),
		TimeService:       time.NewTimeService(token),
		UrlService:        url.NewUrlService(token),
		UserService:       user.NewUserService(token),
		WeatherService:    weather.NewWeatherService(token),
	}
}

type Client struct {
	token string

	AddressService    *address.AddressService
	AnswerService     *answer.AnswerService
	CacheService      *cache.CacheService
	CryptoService     *crypto.CryptoService
	CurrencyService   *currency.CurrencyService
	DbService         *db.DbService
	EmailService      *email.EmailService
	EmojiService      *emoji.EmojiService
	FileService       *file.FileService
	ForexService      *forex.ForexService
	GeocodingService  *geocoding.GeocodingService
	HelloworldService *helloworld.HelloworldService
	IdService         *id.IdService
	ImageService      *image.ImageService
	IpService         *ip.IpService
	LocationService   *location.LocationService
	OtpService        *otp.OtpService
	PostcodeService   *postcode.PostcodeService
	RoutingService    *routing.RoutingService
	RssService        *rss.RssService
	SentimentService  *sentiment.SentimentService
	SmsService        *sms.SmsService
	StockService      *stock.StockService
	StreamService     *stream.StreamService
	ThumbnailService  *thumbnail.ThumbnailService
	TimeService       *time.TimeService
	UrlService        *url.UrlService
	UserService       *user.UserService
	WeatherService    *weather.WeatherService
}
