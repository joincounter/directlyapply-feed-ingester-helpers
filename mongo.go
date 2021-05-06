package helpers

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FetchPersistentJobs fetches limited job data from the indexPersistent collection
func FetchPersistentJobs(ctx context.Context, filter interface{}, serverAddr string, countryCode string) ([]PersistentIndexJob, error) {
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

// DeletePersistentJobsClient remove drinks from
func DeletePersistentJobsClient(ctx context.Context, ids []string, serverAddr string, countryCode string) (*mongo.UpdateResult, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(serverAddr))
	if err != nil {
		return nil, err
	}
	query := bson.M{"_id": bson.M{"$in": ids}}
	update := bson.M{"$set": bson.M{"deletedFromIndex": true, "deletedTime": time.Now()}, "$unset": bson.M{"description": ""}}
	return client.Database(countryDatabases[countryCode]).Collection("indexPersistent").UpdateMany(ctx, query, update)
}

// DeletePersistentJobs remove drinks from
func DeletePersistentJobs(client mongo.Client, ids []string, serverAddr string, countryCode string) (*mongo.UpdateResult, error) {
	query := bson.M{"_id": bson.M{"$in": ids}}
	update := bson.M{"$set": bson.M{"deletedFromIndex": true, "deletedTime": time.Now()}, "$unset": bson.M{"description": ""}}
	return client.Database(countryDatabases[countryCode]).Collection("indexPersistent").UpdateMany(context.Background(), query, update)
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
