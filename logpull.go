package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"os"
	"github.com/gin-gonic/gin"
	"runtime"
	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"bufio"
	"bytes"
	"io"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/pull/*log_path", func(c *gin.Context) {
		path := c.Param("log_path")
		log.Debugf("path: %s", path)
		
		pathSplit := strings.Split(path, "&")
		justPath := pathSplit[0]
		log.Debugf("justPath: %s", justPath)

		offset := c.DefaultQuery("offset", "-1")
		offsetInt, err := strconv.Atoi(offset)
		if err != nil {
			log.Errorf("failed to convert offset string %s to int: %s", offset, err)
			c.AbortWithStatus(400)
		}

		head := c.DefaultQuery("head", "-1")
		headInt, err := strconv.Atoi(head)
		if err != nil {
			log.Errorf("failed to convert head string %s to int: %s", head, err)
			c.AbortWithStatus(400)
		}

		tail := c.DefaultQuery("tail", "-1")
		tailInt, err := strconv.Atoi(tail)
		if err != nil {
			log.Errorf("failed to convert tail string %s to int: %s", tail, err)
			c.AbortWithStatus(400)
		}

		count := c.DefaultQuery("count", "-1")
		countInt, err := strconv.Atoi(count)

		c.String(http.StatusOK, fmt.Sprintf("pull path: \"%s\", head: %d, tail: %d, offset: %d, count: %d", justPath, headInt, tailInt, offsetInt, countInt))

		// todo: validate that the file path exists
		if _, err := os.Stat(justPath); os.IsNotExist(err) {
			c.AbortWithStatus(404)
		}

		// validate query params
		// head and tail and offset are mutually exclusive
		// head=x and nothing else
		useHead := false
		useTail := false
		useOffset := false
		useCount := false
		if headInt > 0 && tail == "" && offset == "" && count == "" {
			// todo: read file and get first n lines
			useHead = true
		}
		// tail=x and nothing else
		else if tailInt > 0 && head == "" && offset == "" && count == "" {
			// todo: read file and get last n lines
			useTail = true
		}
		// offset=x&count=y and nothing else
		else if offsetInt > 0 && countInt > 0 && head == "" && tail == "" {
			// todo: go to line X in file and get Y lines
			useOffset = true
			useCount = true
		}
		// otherwise the state is invalid
		else {
			log.Errorf("invalid combination of parameters: %s", path)
			c.AbortWithStatus(400)
		}

		// todo: validate that the file path is allowed
		file, err := os.Open(justPath)
		defer file.close()
		if err != nil {
			log.Errorf("failed to open file %s: %s", justPath, err)
			c.AbortWithStatus(500)
		}

		reader := bufio.NewReader(file)

		var line string
		var lines []string
		lineNum := 1
		for {

			line, err = reader.ReadString("\n")
			log.Debugf(" Read %d chars\n", len(line))

			iff err != nil {
				log.Errorf("failed to read line from file: %s", err)
				c.AbortWithStatus(500)
			}

			if useHead == true {
				lines = append(lines, line)
			} else if useTail == true {
				if 
			}

			
		}
	})

	r.GET("/patterns", func(c *gin.Context) {
		c.String(http.StatusOK, "patterns x")
	})

	return r
}

func getLinesFromHead(file *os.File, offset int64, count uint) ([]string, error) {
	var lines []string
	reader := bufio.NewReader(file)
	for {
		line, err = reader.ReadString("\n")
		if err != nil {
			if err == io.EOF {
				break
			}

		}

		append(line, lines)
	}

	
}

func getLinesFromTail(*file os.File, count uint) []string{
	var lines []string
}

func main() {
	log.SetHandler(text.New(os.Stdout))
	r := setupRouter()
	r.Run(":8083")
}
