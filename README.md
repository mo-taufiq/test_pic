# test_pic

## soal0
```console
$ cd <another_path>/test_pic/soal0
$ go run main.go
```

## soal1
```console
$ cd <another_path>/test_pic/soal1
$ go run main.go
```

## soal2
```console
$ cd <another_path>/test_pic/soal2
$ go run main.go
```

## soal3
```console
$ cd <another_path>/test_pic/soal3
$ go run main.go
$ curl http://localhost:8080/find_missing_number\?input\=23242526272830
```

## soal4 pull from Docker Hub
```console
$ sudo docker pull motaufiq/go-app:1.0
$ sudo docker run -it -p 8080:8080 motaufiq/go-app:1.0
$ curl http://localhost:8080/find_missing_number\?input\=23242526272830
```

## soal4 build image from Dockerfile manually
```console
$ cd <another_path>/test_pic/soal4
$ sudo docker build --tag go-app:1.0 .
$ sudo docker run -it -p 8080:8080 go-app:1.0
$ curl http://localhost:8080/find_missing_number\?input\=23242526272830
```
    
