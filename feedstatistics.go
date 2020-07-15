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
	fs.assignData(bson.M{"jobsRemovedByFilters": jobsRemovedByFilters})
}

// SetJobsInFeed this will set jobs in feed
func (fs *FeedStatisticsHandler) SetJobsInFeed(jobsInFeed int) {
	fs.assignData(bson.M{"jobsInFeed": jobsInFeed})
}

// SetDuplicatedJobs will write the number of duplicated jobs to the FeedStatisticsHandler
func (fs *FeedStatisticsHandler) SetDuplicatedJobs(duplicates int) {
	fs.assignData(bson.M{"duplicates": duplicates})
}

// SetSentToScraper will write the number of jobs to the scraper microservice
func (fs *FeedStatisticsHandler) SetSentToScraper(sentToScraper int) {
	fs.assignData(bson.M{"sentToScraper": sentToScraper})
}

// SetAlreadyOnSite will write the number of jobs already on the site
func (fs *FeedStatisticsHandler) SetAlreadyOnSite(alreadyOnSite int) {
	fs.assignData(bson.M{"alreadyOnSite": alreadyOnSite})
}

// SetRemovedFromElastic will write the number of jobs removed from elastic
func (fs *FeedStatisticsHandler) SetRemovedFromElastic(removedFromElastic int) {
	fs.assignData(bson.M{"removedFromElastic": removedFromElastic})
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

	_, err = client.Database("directlyapplyjobs").Collection("feedStatistics").UpdateOne(ctx, fs.idFilter(), bson.M{"$set": bson.M{
		"duration": duration,
		"success":  true,
	}})
	if err != nil {
		fmt.Println(err)
	}
}

func (fs *FeedStatisticsHandler) assignData(data interface{}) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fs.ServerAddr))
	if err != nil {
		fmt.Println(err)
	}

	_, err = client.Database("directlyapplyjobs").Collection("feedStatistics").UpdateOne(ctx, fs.idFilter(), bson.M{"$set": data})
	if err != nil {
		fmt.Println(err)
	}
}

func (fs *FeedStatisticsHandler) idFilter() interface{} {
	return bson.M{"_id": bson.M{"$eq": fs.InsertedID}}
}
