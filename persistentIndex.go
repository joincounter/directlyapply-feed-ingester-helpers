package helpers

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FetchPersistentJobs fetches limited job data from the indexPersistent collection
func FetchPersistentJobs(ctx context.Context, filter interface{}, serverAddr string) (*[]PersistentIndexJob, error) {
	appcastJobs := make([]PersistentIndexJob, 0)

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
	ID         string    `bson:"_id" db:"external_id"`
	URL        string    `bson:"url" db:"url"`
	Title      string    `db:"title"`
	CompanyID  uuid.UUID `db:"company_id"`
	LocationID uuid.UUID `db:"location_id"`
	CPA        float32   `db:"cpa"`
	CPC        float32   `db:"cpc"`
}
