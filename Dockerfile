FROM golang
WORKDIR /build
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM busybox

WORKDIR /dist
COPY --from=0 /build/main /dist/main

ENV HTTP_HOST ":80"
CMD ["/dist/main"]