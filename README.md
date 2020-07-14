# weather-crawler

weather-crawlerは[OpenWeather](https://openweathermap.org/)から取得した気象データを
[weather-store](https://github.com/aki-yogiri/weather-store)に追加するサービスです。


# Build Image

Docker でのビルドを想定しています。

```
$ git clone https://github.com/aki-yogiri/weather-crawler.git
$ cd weather-crawler
$ sudo docker build -t weather-crawler:v1.0.3 .
```

# Deploy on Kubernetes

```
$ kubectl apply -f <path>/<to>/<weather-crawler>/kubernetes/weather-csv.yaml
```


# Configuration

weather-crawlerは以下の環境変数を利用します。

| variable | default | |
|----------|---------|-|
| API_KEY | none | [OpenWeather](https://openweathermap.org/)に登録し、API Keyを発行 |
| API_URI | https://api.openweathermap.org/data/2.5/weather | |
| API_LOCATION | Tokyo,jp | `都市名,国`の形式で記述 |
| SERVER_HOST | weather-store | weather-storeサービスのホスト名 |
| SERVER_PORT | 80 | weather-storeサービスのホスト名 |
