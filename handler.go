package main

import (
  "github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net"
  "github.com/eawsy/aws-lambda-go-net/service/lambda/runtime/net/apigatewayproxy"
  "net/http"
  "os"
)

// Handle is the exported handler called by AWS Lambda.
var Handle apigatewayproxy.Handler

func init() {
  ln := net.Listen()

  // Amazon API Gateway binary media types are supported out of the box.
  // If you don't send or receive binary data, you can safely set it to nil.
  Handle = apigatewayproxy.New(ln, []string{"image/png"}).Handle

  // Any Go framework complying with the Go http.Handler interface can be used.
  // This includes, but is not limited to, Vanilla Go, Gin, Echo, Gorrila, Goa, etc.
  go http.Serve(ln, http.HandlerFunc(handle))
}

func handle(w http.ResponseWriter, r *http.Request) {
  key := os.Getenv("KEY_ID")
  value := os.Getenv("KEY_SECRET")
  w.Write([]byte("Hello, World!" + string(key) + string(value)))
}

