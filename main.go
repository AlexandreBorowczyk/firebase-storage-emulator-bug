package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go/v4"
)

func main() {
	ctx := context.Background()

	googleCloudProjectID := "smartflightplanning"
	firebaseConfig := &firebase.Config{
		ProjectID:     googleCloudProjectID,
		StorageBucket: fmt.Sprintf("%s.firebasestorage.app", googleCloudProjectID),
	}
	app, err := firebase.NewApp(
		ctx,
		firebaseConfig,
	)
	client, err := app.Storage(ctx)
	if err != nil {
		fmt.Printf("could not create firebase app: %v\n", err)
		return
	}

	// check if bucket exists
	bucket, err := client.DefaultBucket()
	if err != nil {
		fmt.Printf("could not get default bucket: %v\n", err)
		return
	}
	filename := "gfs.t06z.pgrb2.0p25.f008"
	blob, err := os.ReadFile(filepath.Join("./", filename))
	if err != nil {
		fmt.Printf("could not read file: %v\n", err)
		return
	}

	wc := bucket.Object(filepath.Join("weather", filename)).NewWriter(ctx)
	wc.ChunkSize = 0
	if _, err := wc.Write(blob); err != nil {
		wc.Close()
		fmt.Printf("could not write file: %v\n", err)
		return
	}
	wc.Close()
}
