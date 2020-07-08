module github.com/aki-yogiri/weather-crawler

go 1.14

require (
	github.com/aki-yogiri/weather-crawler/requester v0.0.0-00010101000000-000000000000
	github.com/aki-yogiri/weather-store/pb/weather v0.0.0-20200708103724-0479ff83ee23
	github.com/kelseyhightower/envconfig v1.4.0
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/aki-yogiri/weather-crawler/requester v0.0.0-00010101000000-000000000000 => ./requester
