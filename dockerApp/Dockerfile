FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY main.go /go/src/github.com/edwardIshaq/dockerApp/
WORKDIR /go/src/github.com/edwardIshaq/dockerApp
RUN go get ./
RUN go build -o dockerApp  

CMD [ "dockerApp"]