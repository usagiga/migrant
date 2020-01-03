package model

import (
	"github.com/upamune/go-esa/esa"
	"github.com/usagiga/migrant/entity"
)

type Crawler interface {
	CrawlIter() (post <-chan *esa.Post)
}

func NewCrawler(crawlerType entity.CrawlerType, options interface{}) Crawler {
	switch crawlerType {
	case entity.CrawlerType_EsaOffline:
		return &EsaOfflineCrawlerImpl{}
	}

	return nil
}