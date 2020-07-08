package main

import (
	"context"
	"github.com/aki-yogiri/weather-crawler/requester"
	pb "github.com/aki-yogiri/weather-store/pb/weather"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"log"
	"net/url"
)

type ServerEnv struct {
	Host string
	Port int
}

type ApiEnv struct {
	Uri      string
	Location string
	ApiKey   string
}

func main() {
	var apiEnv ApiEnv
	envconfig.Process("API", &apiEnv)

	endpoint := requester.NewOpenWeatherEndpoint(apiEnv.Uri)
	values := &url.Values{}
	values.Add("appid", apiEnv.ApiKey)
	values.Add("q", apiEnv.Location)

	resp, err := endpoint.Request(values)
	if err != nil {
		log.Fatalln(err)
		return
	}
	resp.Location = apiEnv.Location

	var serverEnv ServerEnv
	envconfig.Process("SERVER", &serverEnv)
	conn, err := grpc.Dial(serverEnv.Host+":"+string(serverEnv.Port), grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()
	client := pb.NewWeatherClient(conn)
	res, err := client.PutWeather(context.TODO(), resp)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(res)
}
