package helpers

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FetchPersistentJobs fetches limited job data from the indexPersistent collection
func FetchPersistentJobs(filter interface{}, serverAddr string, countryCode string) ([]PersistentIndexJob, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
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
	jobObjectIdsNotInFeed := make([]primitive.ObjectID, 0)
	for i := 0; i < len(ids); i++ {
		objectId, _ := primitive.ObjectIDFromHex(ids[i])
		jobObjectIdsNotInFeed = append(jobObjectIdsNotInFeed, objectId)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(serverAddr))
	if err != nil {
		return nil, err
	}
	query := bson.M{"_id": bson.M{"$in": jobObjectIdsNotInFeed}}
	update := bson.M{"$set": bson.M{"deletedFromIndex": true, "deletedTime": time.Now()}, "$unset": bson.M{"description": ""}}
	return client.Database(countryDatabases[countryCode]).Collection("indexPersistent").UpdateMany(ctx, query, update)
}

func UpsertEmployers(jobs []StandardJob, serverAddr string, countryCode string) (int, error) {
	var addedCompanies int
	var processedSlugs []string
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Hour)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(serverAddr))
	if err != nil {
		return addedCompanies, err
	}
	for i := 0; i < len(jobs); i++ {
		job := jobs[i]
		if !contains(processedSlugs, job.Slug) {
			processedSlugs = append(processedSlugs, job.Slug)
			res := client.Database(countryDatabases[countryCode]).Collection("indexPersistent").FindOne(ctx, bson.M{"slug": job.Slug})
			if res.Err() == mongo.ErrNoDocuments {
				object := bson.M{
					"employer": job.Company,
					"slug":     job.Slug}
				_, err := client.Database(countryDatabases[countryCode]).Collection("indexPersistent").InsertOne(ctx, object)
				if err != nil {
					return addedCompanies, err
				} else {
					addedCompanies++
				}
			} else if res.Err() != nil {
				return addedCompanies, res.Err()
			}
		}
	}
	return addedCompanies, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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
