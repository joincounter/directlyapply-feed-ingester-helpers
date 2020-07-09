package helpers

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FetchPersistentJobs fetches limited job data from the indexPersistent collection
func FetchPersistentJobs(filter interface{}, serverAddr string) (*[]PersistentIndexJob, error) {
	appcastJobs := make([]PersistentIndexJob, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(serverAddr))
	if err != nil {
		return nil, err
	}

	cur, err := client.Database("directlyapplyjobs").Collection("indexPersistent").Find(ctx, filter, &options.FindOptions{Projection: bson.M{"url": 1}})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(context.Background()) {
		var job PersistentIndexJob
		cur.Decode(&job)
		appcastJobs = append(appcastJobs, job)
	}

	return &appcastJobs, nil
}

// PersistentIndexJob minimum representaion of a job
type PersistentIndexJob struct {
	ID  string `bson:"_id"`
	URL string `bson:"url"`
}
