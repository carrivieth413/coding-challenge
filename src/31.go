package main

import (
    "fmt"
    "strings"

    "github.com/gomodule/redigo/redis"
)

func main() {
    // Connect to Redis server
    redisClient := redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379",
        Password: "",      // Empty string means `redis://`
        AuthKey:   "",
    })

    // Get a random value from the database
    var randomValue string

    if err := redisClient.Do("MGET", []string{""}); err != nil {
        fmt.Println(err)
    } else {
        randomValue = strings.TrimSpace(redisClient.Do("SMEMBERS", ""))

        if len(randomValue) == 0 {
            fmt.Println("No members found in the database")
        } else {
            // Split the value into an array
            values := strings.Split(randomValue, ", ")

            // Check if there's at least one member
            if len(values) > 0 {
                fmt.Printf("Found a member with value: %s\n", values[0])
            } else {
                fmt.Println("No members found in the database")
            }
        }
    }

    redisClient.Close()
}
