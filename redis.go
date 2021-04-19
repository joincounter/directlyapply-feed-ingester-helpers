package helpers

import (
	"github.com/gomodule/redigo/redis"
)

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
