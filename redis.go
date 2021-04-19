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

	_, err = conn.Do("SADD", append([]string{feedname}, urls...))
	if err != nil {
		return err
	}
	return nil
}
