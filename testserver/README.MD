CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -installsuffix cgo -ldflags '-s' -o server  

docker build -t mirkar/goserver:latest $PWD


docker run -d --rm --name goserver  \
  -p 127.0.0.1:8087:8087 mirkar/goserver:latest
