package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestSetAndGet(t *testing.T) {
	originalUrl := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	// Persist data mapping
	SaveUrlMapping(shortURL, originalUrl, userUUId)

	// Retrieve initial URL
	resultUrl := GetOriginalUrl(shortURL)

	assert.Equal(t, originalUrl, resultUrl)
}
