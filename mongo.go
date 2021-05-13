package helpers

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FetchPersistentJobs fetches limited job data from the indexPersistent collection
func FetchPersistentJobs(filter interface{}, serverAddr string, countryCode string) ([]PersistentIndexJob, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	appcastJobs := make([]PersistentIndexJob, 0)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(serverAddr))
	if err != nil {
		return appcastJobs, err
	}

	cur, err := client.Database(countryDatabases[countryCode]).Collection("indexPersistent").Find(ctx, filter, &options.FindOptions{Projection: bson.M{"url": 1}})
	if err != nil {
		return appcastJobs, err
	}
	defer cur.Close(ctx)
	for cur.Next(context.Background()) {
		var job PersistentIndexJob
		cur.Decode(&job)
		appcastJobs = append(appcastJobs, job)
	}

	return appcastJobs, nil
}

// FetchPersistentJobsFromURLs fetches limited job data from the indexPersistent collection
func FetchPersistentJobsFromURLs(urls []string, serverAddr string, countryCode string) ([]PersistentIndexJob, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	appcastJobs := make([]PersistentIndexJob, 0)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(serverAddr))
	if err != nil {
		return appcastJobs, err
	}

	cur, err := client.Database(countryDatabases[countryCode]).Collection("indexPersistent").Find(ctx, bson.M{"url": bson.M{"$in": urls}}, &options.FindOptions{Projection: bson.M{"url": 1}})
	if err != nil {
		return appcastJobs, err
	}
	defer cur.Close(ctx)
	for cur.Next(context.Background()) {
		var job PersistentIndexJob
		cur.Decode(&job)
		appcastJobs = append(appcastJobs, job)
	}

	return appcastJobs, nil
}

// DeletePersistentJobs remove drinks from
func DeletePersistentJobs(ids []string, serverAddr string, countryCode string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(serverAddr))
	if err != nil {
		return nil, err
	}
	query := bson.M{"_id": bson.M{"$in": ids}}
	update := bson.M{"$set": bson.M{"deletedFromIndex": true, "deletedTime": time.Now()}, "$unset": bson.M{"description": ""}}
	return client.Database(countryDatabases[countryCode]).Collection("indexPersistent").UpdateMany(ctx, query, update)
}

// PersistentIndexJob minimum representaion of a job
type PersistentIndexJob struct {
	ID  string `bson:"_id" db:"external_id"`
	URL string `bson:"url" db:"url"`
}

var countryDatabases = map[string]string{
	"UK": "counterjobs",
	"US": "directlyapplyjobs",
	"CA": "directlyapplyjobs",
}
