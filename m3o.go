package m3o

import (
	"go.m3o.com/address"
	"go.m3o.com/analytics"
	"go.m3o.com/answer"
	"go.m3o.com/app"
	"go.m3o.com/avatar"
	"go.m3o.com/bitcoin"
	"go.m3o.com/cache"
	"go.m3o.com/carbon"
	"go.m3o.com/chat"
	"go.m3o.com/comments"
	"go.m3o.com/contact"
	"go.m3o.com/crypto"
	"go.m3o.com/currency"
	"go.m3o.com/db"
	"go.m3o.com/dns"
	"go.m3o.com/email"
	"go.m3o.com/emoji"
	"go.m3o.com/evchargers"
	"go.m3o.com/event"
	"go.m3o.com/file"
	"go.m3o.com/forex"
	"go.m3o.com/function"
	"go.m3o.com/geocoding"
	"go.m3o.com/gifs"
	"go.m3o.com/google"
	"go.m3o.com/helloworld"
	"go.m3o.com/holidays"
	"go.m3o.com/id"
	"go.m3o.com/image"
	"go.m3o.com/ip"
	"go.m3o.com/joke"
	"go.m3o.com/lists"
	"go.m3o.com/location"
	"go.m3o.com/memegen"
	"go.m3o.com/minecraft"
	"go.m3o.com/movie"
	"go.m3o.com/mq"
	"go.m3o.com/news"
	"go.m3o.com/nft"
	"go.m3o.com/notes"
	"go.m3o.com/otp"
	"go.m3o.com/password"
	"go.m3o.com/ping"
	"go.m3o.com/place"
	"go.m3o.com/postcode"
	"go.m3o.com/prayer"
	"go.m3o.com/price"
	"go.m3o.com/qr"
	"go.m3o.com/quran"
	"go.m3o.com/routing"
	"go.m3o.com/rss"
	"go.m3o.com/search"
	"go.m3o.com/sentiment"
	"go.m3o.com/sms"
	"go.m3o.com/space"
	"go.m3o.com/spam"
	"go.m3o.com/stock"
	"go.m3o.com/stream"
	"go.m3o.com/sunnah"
	"go.m3o.com/thumbnail"
	"go.m3o.com/time"
	"go.m3o.com/translate"
	"go.m3o.com/tunnel"
	"go.m3o.com/twitter"
	"go.m3o.com/url"
	"go.m3o.com/user"
	"go.m3o.com/vehicle"
	"go.m3o.com/weather"
	"go.m3o.com/wordle"
	"go.m3o.com/youtube"
)

func New(token string) *Client {
	return &Client{
		token: token,

		Address:    address.NewAddressService(token),
		Analytics:  analytics.NewAnalyticsService(token),
		Answer:     answer.NewAnswerService(token),
		App:        app.NewAppService(token),
		Avatar:     avatar.NewAvatarService(token),
		Bitcoin:    bitcoin.NewBitcoinService(token),
		Cache:      cache.NewCacheService(token),
		Carbon:     carbon.NewCarbonService(token),
		Chat:       chat.NewChatService(token),
		Comments:   comments.NewCommentsService(token),
		Contact:    contact.NewContactService(token),
		Crypto:     crypto.NewCryptoService(token),
		Currency:   currency.NewCurrencyService(token),
		Db:         db.NewDbService(token),
		Dns:        dns.NewDnsService(token),
		Email:      email.NewEmailService(token),
		Emoji:      emoji.NewEmojiService(token),
		Evchargers: evchargers.NewEvchargersService(token),
		Event:      event.NewEventService(token),
		File:       file.NewFileService(token),
		Forex:      forex.NewForexService(token),
		Function:   function.NewFunctionService(token),
		Geocoding:  geocoding.NewGeocodingService(token),
		Gifs:       gifs.NewGifsService(token),
		Google:     google.NewGoogleService(token),
		Helloworld: helloworld.NewHelloworldService(token),
		Holidays:   holidays.NewHolidaysService(token),
		Id:         id.NewIdService(token),
		Image:      image.NewImageService(token),
		Ip:         ip.NewIpService(token),
		Joke:       joke.NewJokeService(token),
		Lists:      lists.NewListsService(token),
		Location:   location.NewLocationService(token),
		Memegen:    memegen.NewMemegenService(token),
		Minecraft:  minecraft.NewMinecraftService(token),
		Movie:      movie.NewMovieService(token),
		Mq:         mq.NewMqService(token),
		News:       news.NewNewsService(token),
		Nft:        nft.NewNftService(token),
		Notes:      notes.NewNotesService(token),
		Otp:        otp.NewOtpService(token),
		Password:   password.NewPasswordService(token),
		Ping:       ping.NewPingService(token),
		Place:      place.NewPlaceService(token),
		Postcode:   postcode.NewPostcodeService(token),
		Prayer:     prayer.NewPrayerService(token),
		Price:      price.NewPriceService(token),
		Qr:         qr.NewQrService(token),
		Quran:      quran.NewQuranService(token),
		Routing:    routing.NewRoutingService(token),
		Rss:        rss.NewRssService(token),
		Search:     search.NewSearchService(token),
		Sentiment:  sentiment.NewSentimentService(token),
		Sms:        sms.NewSmsService(token),
		Space:      space.NewSpaceService(token),
		Spam:       spam.NewSpamService(token),
		Stock:      stock.NewStockService(token),
		Stream:     stream.NewStreamService(token),
		Sunnah:     sunnah.NewSunnahService(token),
		Thumbnail:  thumbnail.NewThumbnailService(token),
		Time:       time.NewTimeService(token),
		Translate:  translate.NewTranslateService(token),
		Tunnel:     tunnel.NewTunnelService(token),
		Twitter:    twitter.NewTwitterService(token),
		Url:        url.NewUrlService(token),
		User:       user.NewUserService(token),
		Vehicle:    vehicle.NewVehicleService(token),
		Weather:    weather.NewWeatherService(token),
		Wordle:     wordle.NewWordleService(token),
		Youtube:    youtube.NewYoutubeService(token),
	}
}

type Client struct {
	token string

	Address    address.Address
	Analytics  analytics.Analytics
	Answer     answer.Answer
	App        app.App
	Avatar     avatar.Avatar
	Bitcoin    bitcoin.Bitcoin
	Cache      cache.Cache
	Carbon     carbon.Carbon
	Chat       chat.Chat
	Comments   comments.Comments
	Contact    contact.Contact
	Crypto     crypto.Crypto
	Currency   currency.Currency
	Db         db.Db
	Dns        dns.Dns
	Email      email.Email
	Emoji      emoji.Emoji
	Evchargers evchargers.Evchargers
	Event      event.Event
	File       file.File
	Forex      forex.Forex
	Function   function.Function
	Geocoding  geocoding.Geocoding
	Gifs       gifs.Gifs
	Google     google.Google
	Helloworld helloworld.Helloworld
	Holidays   holidays.Holidays
	Id         id.Id
	Image      image.Image
	Ip         ip.Ip
	Joke       joke.Joke
	Lists      lists.Lists
	Location   location.Location
	Memegen    memegen.Memegen
	Minecraft  minecraft.Minecraft
	Movie      movie.Movie
	Mq         mq.Mq
	News       news.News
	Nft        nft.Nft
	Notes      notes.Notes
	Otp        otp.Otp
	Password   password.Password
	Ping       ping.Ping
	Place      place.Place
	Postcode   postcode.Postcode
	Prayer     prayer.Prayer
	Price      price.Price
	Qr         qr.Qr
	Quran      quran.Quran
	Routing    routing.Routing
	Rss        rss.Rss
	Search     search.Search
	Sentiment  sentiment.Sentiment
	Sms        sms.Sms
	Space      space.Space
	Spam       spam.Spam
	Stock      stock.Stock
	Stream     stream.Stream
	Sunnah     sunnah.Sunnah
	Thumbnail  thumbnail.Thumbnail
	Time       time.Time
	Translate  translate.Translate
	Tunnel     tunnel.Tunnel
	Twitter    twitter.Twitter
	Url        url.Url
	User       user.User
	Vehicle    vehicle.Vehicle
	Weather    weather.Weather
	Wordle     wordle.Wordle
	Youtube    youtube.Youtube
}
