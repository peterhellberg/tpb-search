package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/blevesearch/bleve"
)

func buildIndexMapping() *bleve.IndexMapping {
	// a generic reusable mapping for keyword text
	keywordFieldMapping := bleve.NewTextFieldMapping()
	keywordFieldMapping.Analyzer = "keyword"

	torrentMapping := bleve.NewDocumentMapping()
	torrentMapping.AddFieldMappingsAt("category", keywordFieldMapping)

	im := bleve.NewIndexMapping()

	im.AddDocumentMapping("torrent", torrentMapping)
	im.TypeField = "type"
	im.DefaultAnalyzer = "standard"
	im.DefaultType = "torrent"
	im.DefaultField = "name"

	return im
}

type tpbDoc struct {
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	Hash     string `json:"hash"`
	Category string `json:"category"`
	Type     string `json:"type"`
}

func indexTPB(i bleve.Index) error {
	log.Printf("Indexing...")

	count := 0
	startTime := time.Now()

	batch := bleve.NewBatch()
	batchCount := 0

	dumpFile, _ := os.Open(*dump)
	defer dumpFile.Close()

	reader := csv.NewReader(dumpFile)
	reader.LazyQuotes = true
	reader.FieldsPerRecord = 7
	reader.Comma = '|'

	for {
		r, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			continue
		}

		size, err := strconv.ParseInt(r[1], 10, 0)
		if err != nil {
			fmt.Println("%#v", size)
			size = 0
		}

		batch.Index(r[2], tpbDoc{
			Name:     r[0],
			Size:     size,
			Hash:     r[2],
			Category: r[4],
			Type:     "torrent",
		})
		batchCount++

		if batchCount >= *batchSize {
			err = i.Batch(batch)
			if err != nil {
				return err
			}
			batch = bleve.NewBatch()
			batchCount = 0
		}

		count++

		if count%1000 == 0 {
			indexDuration := time.Since(startTime)
			indexDurationSeconds := float64(indexDuration) / float64(time.Second)
			timePerDoc := float64(indexDuration) / float64(count)
			log.Printf("Indexed %d documents in %.2fs (average %.2fms/doc)", count, indexDurationSeconds, timePerDoc/float64(time.Millisecond))
		}

		if *indexLimit > 0 && count >= *indexLimit {
			break
		}
	}

	// flush the last batch
	if batchCount > 0 {
		err := i.Batch(batch)
		if err != nil {
			log.Fatal(err)
		}
	}

	indexDuration := time.Since(startTime)
	indexDurationSeconds := float64(indexDuration) / float64(time.Second)
	timePerDoc := float64(indexDuration) / float64(count)
	log.Printf("Finished indexing %d documents in %.2fs (average %.2fms/doc)", count, indexDurationSeconds, timePerDoc/float64(time.Millisecond))

	log.Printf("Still listening on http://%v", bind)

	return nil
}
