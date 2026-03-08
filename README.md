# rk_sensor

rk sensor is a simulated sensor that can data on diffrent protocols

## Sensor Types
* Temprature

## Protocal supported
* HTTP, HTTPS

## Run
```
go run cmd/rk_sensor/main.go
```
or with docker  

```
docker build -t rk_sensor -f .docker/Dockerfile .
docker run --env sender_url="http://localhost:8080/endpoint" --env sensor_type="temprature" rk_sensor
```

