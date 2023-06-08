FROM golang:latest
WORKDIR app
ARG LOG_DIR=/app/logs
RUN mkdir -p ${LOG_DIR}
ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install
EXPOSE 8089
VOLUME [${LOG_DIR}]
CMD ["API-server" ,"start"]
#  docker build -t pritam99/api-server:0.0.3 .
# docker push pritam99/api-server:0.0.3
# docker run -d -p 8089:8089 -v ~/logs/go-docker:/app/logs pritam99/api-server:0.0.3