package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nlittlepoole/whenami"
	"github.com/sirupsen/logrus"
	"github.com/toorop/gin-logrus"
	"strconv"
)

// Result has a list of errors and a timezone if one could be computed
type Result struct {
	Errors   []string `json:"errors"`
	Timezone string   `json:"timezone"`
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
		lat, latErr := strconv.ParseFloat(c.Query("latitude"), 64)
		if latErr != nil {
			log.Warn(latErr)
			errors = append(errors, latErr.Error())
		}
		long, longErr := strconv.ParseFloat(c.Query("longitude"), 64)
		if longErr != nil {
			log.Warn(longErr)
			errors = append(errors, longErr.Error())
		}
		if latErr == nil && longEerr == nil {
			var err error
			timezone, err = whenami.WhenAmI(lat, long)
			if err != nil {
				log.Warn(err)
				errors = append(errors, err.Error())
			}
		}
		result := &Result{
			Errors:   errors,
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
