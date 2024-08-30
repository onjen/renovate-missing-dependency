package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gocloud.dev/blob/gcsblob"
	"gocloud.dev/gcp"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	ctx := context.Background()
	gcpclient, err := gcp.NewHTTPClient(
		gcp.DefaultTransport(),
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	bucketReader, err := gcsblob.OpenBucket(ctx, gcpclient, "gcp-public-data-arco-era5", nil)
	if err != nil {
		log.Fatalf("Failed to create bucket reader: %v", err)
	}
	defer bucketReader.Close()

	it := bucketReader.List(nil)
	for {
		obj, err := it.Next(ctx)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate bucket objets: %v", err)
		}
		fmt.Println(obj.Key)
	}

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))

}
