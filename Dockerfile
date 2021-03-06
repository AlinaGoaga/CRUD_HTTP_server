FROM golang:latest 
RUN go get -u github.com/gorilla/mux
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
CMD ["/app/main"]