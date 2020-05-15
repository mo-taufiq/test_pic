# test_pic

## soal0
    `cd path/soal0`
    `go run main.go`

## soal1
    `cd path/soal1`
    `go run main.go`

## soal2
    `cd path/soal2`
    `go run main.go`

## soal3
    `cd path/soal3`
    `go run main.go`
    `curl http://localhost:8080/find_missing_number\?input\=23242526272830`

## soal4 pull from Docker Hub
    `sudo docker pull motaufiq/go-app:1.0`
    `sudo docker run -it -p 8080:8080 motaufiq/go-app:1.0`

## soal4 build image from Dockerfile manually
    `cd path/soal4`
    `sudo docker build --tag go-app:1.0 .`
    `sudo docker run -it -p 8080:8080 go-app:1.0`