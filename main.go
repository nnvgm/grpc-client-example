package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nnvgm/grpc-common-example/proto/math"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type httpRequestSum struct {
	A int32
	B int32
}

var (
	c   math.MathClient
	ctx = context.Background()
)

func init() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", os.Getenv("GRPC_HOST"), os.Getenv("GRPC_PORT")), grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Start GRPC Client")

	c = math.NewMathClient(conn)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/sum", sum).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r))
}

func sum(w http.ResponseWriter, r *http.Request) {
	var body httpRequestSum

	json.NewDecoder(r.Body).Decode(&body)

	ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	sum, err := c.Sum(ctxTimeout, &math.SumRequest{
		A: body.A,
		B: body.B,
	})

	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(sum.String()))
}
