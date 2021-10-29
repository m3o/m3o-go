package m3o

import (
	"github.com/m3o/m3o-go/address"
	"github.com/m3o/m3o-go/answer"
	"github.com/m3o/m3o-go/cache"
	"github.com/m3o/m3o-go/crypto"
	"github.com/m3o/m3o-go/currency"
	"github.com/m3o/m3o-go/db"
	"github.com/m3o/m3o-go/email"
	"github.com/m3o/m3o-go/emoji"
	"github.com/m3o/m3o-go/evchargers"
	"github.com/m3o/m3o-go/file"
	"github.com/m3o/m3o-go/forex"
	"github.com/m3o/m3o-go/function"
	"github.com/m3o/m3o-go/geocoding"
	"github.com/m3o/m3o-go/gifs"
	"github.com/m3o/m3o-go/google"
	"github.com/m3o/m3o-go/helloworld"
	"github.com/m3o/m3o-go/holidays"
	"github.com/m3o/m3o-go/id"
	"github.com/m3o/m3o-go/image"
	"github.com/m3o/m3o-go/ip"
	"github.com/m3o/m3o-go/location"
	"github.com/m3o/m3o-go/notes"
	"github.com/m3o/m3o-go/otp"
	"github.com/m3o/m3o-go/postcode"
	"github.com/m3o/m3o-go/prayer"
	"github.com/m3o/m3o-go/qr"
	"github.com/m3o/m3o-go/quran"
	"github.com/m3o/m3o-go/routing"
	"github.com/m3o/m3o-go/rss"
	"github.com/m3o/m3o-go/sentiment"
	"github.com/m3o/m3o-go/sms"
	"github.com/m3o/m3o-go/stock"
	"github.com/m3o/m3o-go/stream"
	"github.com/m3o/m3o-go/sunnah"
	"github.com/m3o/m3o-go/thumbnail"
	"github.com/m3o/m3o-go/time"
	"github.com/m3o/m3o-go/twitter"
	"github.com/m3o/m3o-go/url"
	"github.com/m3o/m3o-go/user"
	"github.com/m3o/m3o-go/vehicle"
	"github.com/m3o/m3o-go/weather"
	"github.com/m3o/m3o-go/youtube"
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
		EvchargersService: evchargers.NewEvchargersService(token),
		FileService:       file.NewFileService(token),
		ForexService:      forex.NewForexService(token),
		FunctionService:   function.NewFunctionService(token),
		GeocodingService:  geocoding.NewGeocodingService(token),
		GifsService:       gifs.NewGifsService(token),
		GoogleService:     google.NewGoogleService(token),
		HelloworldService: helloworld.NewHelloworldService(token),
		HolidaysService:   holidays.NewHolidaysService(token),
		IdService:         id.NewIdService(token),
		ImageService:      image.NewImageService(token),
		IpService:         ip.NewIpService(token),
		LocationService:   location.NewLocationService(token),
		NotesService:      notes.NewNotesService(token),
		OtpService:        otp.NewOtpService(token),
		PostcodeService:   postcode.NewPostcodeService(token),
		PrayerService:     prayer.NewPrayerService(token),
		QrService:         qr.NewQrService(token),
		QuranService:      quran.NewQuranService(token),
		RoutingService:    routing.NewRoutingService(token),
		RssService:        rss.NewRssService(token),
		SentimentService:  sentiment.NewSentimentService(token),
		SmsService:        sms.NewSmsService(token),
		StockService:      stock.NewStockService(token),
		StreamService:     stream.NewStreamService(token),
		SunnahService:     sunnah.NewSunnahService(token),
		ThumbnailService:  thumbnail.NewThumbnailService(token),
		TimeService:       time.NewTimeService(token),
		TwitterService:    twitter.NewTwitterService(token),
		UrlService:        url.NewUrlService(token),
		UserService:       user.NewUserService(token),
		VehicleService:    vehicle.NewVehicleService(token),
		WeatherService:    weather.NewWeatherService(token),
		YoutubeService:    youtube.NewYoutubeService(token),
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
	EvchargersService *evchargers.EvchargersService
	FileService       *file.FileService
	ForexService      *forex.ForexService
	FunctionService   *function.FunctionService
	GeocodingService  *geocoding.GeocodingService
	GifsService       *gifs.GifsService
	GoogleService     *google.GoogleService
	HelloworldService *helloworld.HelloworldService
	HolidaysService   *holidays.HolidaysService
	IdService         *id.IdService
	ImageService      *image.ImageService
	IpService         *ip.IpService
	LocationService   *location.LocationService
	NotesService      *notes.NotesService
	OtpService        *otp.OtpService
	PostcodeService   *postcode.PostcodeService
	PrayerService     *prayer.PrayerService
	QrService         *qr.QrService
	QuranService      *quran.QuranService
	RoutingService    *routing.RoutingService
	RssService        *rss.RssService
	SentimentService  *sentiment.SentimentService
	SmsService        *sms.SmsService
	StockService      *stock.StockService
	StreamService     *stream.StreamService
	SunnahService     *sunnah.SunnahService
	ThumbnailService  *thumbnail.ThumbnailService
	TimeService       *time.TimeService
	TwitterService    *twitter.TwitterService
	UrlService        *url.UrlService
	UserService       *user.UserService
	VehicleService    *vehicle.VehicleService
	WeatherService    *weather.WeatherService
	YoutubeService    *youtube.YoutubeService
}
