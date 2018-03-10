package main

import (
    "strconv"
    "github.com/sirupsen/logrus"
    "github.com/toorop/gin-logrus"
    "github.com/gin-gonic/gin"
    "github.com/nlittlepoole/whenami"
)

type Result struct {
	Errors []string `json:"errors"`
	Timezone string `json:"timezone"`
}

func main() {
	log := logrus.New()
	r := gin.New()
	r.Use(ginlogrus.Logger(log), gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})
	r.GET("/whenami/", func(c *gin.Context) {
		var timezone string
		errors := make([]string, 0)
		lat, lat_err := strconv.ParseFloat(c.Query("latitude"), 64)
		if lat_err != nil{
			log.Warn(lat_err)
			errors = append(errors, lat_err.Error())
		}
		long, long_err := strconv.ParseFloat(c.Query("longitude"), 64)
		if long_err != nil{
			log.Warn(long_err)
			errors = append(errors, long_err.Error())
		}
		if lat_err == nil && long_err == nil {
			var err error
			timezone, err = whenami.WhenAmI(lat, long)
			if err != nil{
				log.Warn(err)
				errors = append(errors, err.Error())
			}
		}
		result := &Result{
			Errors: errors,
			Timezone: timezone,
		}
		if len(errors) > 0 {
			c.JSON(500, result)
		} else {
			c.JSON(200, result)
		}
	})
	r.Run(":80")
}