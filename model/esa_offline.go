package model

import (
	"github.com/upamune/go-esa/esa"
	"github.com/usagiga/migrant/lib/path"
	"io/ioutil"
	"path/filepath"
	"regexp"
)

type EsaOfflineCrawlerImpl struct {

}

func (c *EsaOfflineCrawlerImpl) CrawlIter() (post <-chan *esa.Post) {
	ch := make(chan *esa.Post)

	go func() {
		paths := path.Dirwalk("backup")
		frontMatterRule := regexp.MustCompile("(?s)---.*?---\\n\\n")

		for _, file := range paths {
			category := filepath.Dir(file)
			name := path.GetFileNameWithoutExt(file)
			rawBody, _ := ioutil.ReadFile(file)

			// Front Matter の削除
			body := frontMatterRule.ReplaceAllLiteral(rawBody, nil)

			ch <- &esa.Post{
				Category: category,
				BodyMd:   string(body),
				Name:     name,
				Message:  "Cloned from dump of esa.",
				Tags:     nil,
				Wip:      false,
			}
		}

		close(ch)
	}()

	return ch
}
