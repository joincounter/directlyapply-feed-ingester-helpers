package helpers

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FeedStatisticsHandler controls feed statistics
type FeedStatisticsHandler struct {
	InsertedID interface{}
	ServerAddr string
	StartTime  time.Time
}

// NewFeedStatistics this create a new feed statistics
func NewFeedStatistics(url, feed, country, serverAddr string) FeedStatisticsHandler {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(serverAddr))
	if err != nil {
		fmt.Println(err)
	}

	startTime := time.Now()

	cur, err := client.Database("directlyapplyjobs").Collection("feedStatistics").InsertOne(ctx, bson.M{
		"url":       url,
		"feed":      feed,
		"country":   country,
		"startTime": startTime,
		"success":   false,
	})

	return FeedStatisticsHandler{
		StartTime:  startTime,
		InsertedID: cur.InsertedID,
		ServerAddr: serverAddr,
	}
}

// SetJobsRemovedByFilters this will set jobs removed by filter
func (fs *FeedStatisticsHandler) SetJobsRemovedByFilters(jobsRemovedByFilters int) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fs.ServerAddr))
	if err != nil {
		fmt.Println(err)
	}

	_, err = client.Database("directlyapplyjobs").Collection("feedStatistics").UpdateOne(ctx, fs.idFilter(), bson.M{"$set": bson.M{"jobsRemovedByFilters": jobsRemovedByFilters}})
	if err != nil {
		fmt.Println(err)
	}
}

// SetJobsInFeed this will set jobs in feed
func (fs *FeedStatisticsHandler) SetJobsInFeed(jobsInFeed int) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fs.ServerAddr))
	if err != nil {
		fmt.Println(err)
	}

	_, err = client.Database("directlyapplyjobs").Collection("feedStatistics").UpdateOne(ctx, fs.idFilter(), bson.M{"$set": bson.M{"jobsInFeed": jobsInFeed}})
	if err != nil {
		fmt.Println(err)
	}
}

// EndAndSendFeedStatistics this will end and finalize the feed statistics
func (fs *FeedStatisticsHandler) EndAndSendFeedStatistics() {
	duration := time.Since(fs.StartTime)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fs.ServerAddr))
	if err != nil {
		fmt.Println(err)
	}

	_, err = client.Database("directlyapplyjobs").Collection("feedStatistics").UpdateOne(ctx, fs.idFilter(), bson.M{
		"duration": duration,
		"success":  true,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func (fs *FeedStatisticsHandler) idFilter() interface{} {
	return bson.M{"_id": bson.M{"$eq": fs.InsertedID}}
}
