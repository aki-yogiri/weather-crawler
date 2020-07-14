package requester

import (
	"encoding/json"
	"errors"
	pb "github.com/aki-yogiri/weather-store/pb/weather"
	"github.com/golang/protobuf/ptypes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Endpoint interface {
	Request(v *url.Values) (*pb.WeatherMessage, error)
}

/*
  以下のツールを使い構造体を自動生成し、不要なパラメータは削除
  https://mholt.github.io/json-to-go/

  APIのパラメータ一覧は以下の通り
  https://openweathermap.org/current
*/
type OpenWeatherResponse struct {
	Weather []struct {
		ID   int    `json:"id"`
		Main string `json:"main"`
	} `json:"weather"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt       int `json:"dt"`
	Timezone int `json:"timezone"`
}

type OpenWeatherEndpoint struct {
	URI string
}

func NewOpenWeatherEndpoint(uri string) *OpenWeatherEndpoint {
	e := &OpenWeatherEndpoint{}
	e.URI = uri
	return e
}

func (e *OpenWeatherEndpoint) Request(v *url.Values) (*pb.WeatherMessage, error) {
	resp, err := http.Get(e.URI + "?" + v.Encode())

	if err != nil {
		log.Print("Error: %v", err)
		return nil, errors.New("Request Failed")
	}

	defer resp.Body.Close()

	owr, err := mapOpenWeatherResponse(resp)
	if err != nil {
		return nil, err
	}

	message, err := convertWeatherMessage(owr)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func mapOpenWeatherResponse(response *http.Response) (*OpenWeatherResponse, error) {
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Printf("Error: %v", err)
		return nil, errors.New("Could not read response body")
	}

	owr := &OpenWeatherResponse{}
	if err := json.Unmarshal(body, owr); err != nil {
		log.Printf("Error: %v", err)
		return nil, errors.New("JSON Unmarshal Error")
	}

	return owr, nil
}

func convertWeatherMessage(owr *OpenWeatherResponse) (*pb.WeatherMessage, error) {
	var err error
	wm := &pb.WeatherMessage{}
	wm.Weather = owr.Weather[0].Main
	wm.Temperature = owr.Main.Temp
	wm.Clouds = uint32(owr.Clouds.All)
	wm.Wind = owr.Wind.Speed
	wm.WindDeg = uint32(owr.Wind.Deg)
	wm.Timestamp, err = ptypes.TimestampProto(time.Unix(int64(owr.Dt), 0))

	if err != nil {
		return nil, errors.New("Could not convert datetime: " + string(owr.Dt))
	}

	return wm, nil
}
