package helpers

import (
	"github.com/gomodule/redigo/redis"
)

func RenameRedisKey(connectionString, oldKeyName, newKeyName string) error {
	conn, err := redis.DialURL(connectionString)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("RENAME", oldKeyName, newKeyName)

	return err
}

func ExcludeURLSubset(connectionString, primaryFeedname, secondaryFeedname string) ([]string, error) {
	urls := make([]string, 0)
	conn, err := redis.DialURL(connectionString)
	if err != nil {
		return urls, err
	}
	defer conn.Close()

	data, err := conn.Do("SDIFF", primaryFeedname, secondaryFeedname)
	if err != nil {
		return urls, err
	}

	asArry := data.([]interface{})
	for _, d := range asArry {
		xxx := d.([]uint8)
		urls = append(urls, string(xxx))
	}

	return urls, nil
}

func SaveURLsToRedis(connectionString string, feedname string, urls []string) error {
	conn, err := redis.DialURL(connectionString)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Do("DEL", feedname)
	if err != nil {
		return err
	}

	for _, url := range urls {
		_, err = conn.Do("SADD", feedname, url)
		if err != nil {
			return err
		}
	}
	return nil
}
