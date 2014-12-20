package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"time"

	"github.com/blevesearch/bleve"
)

type tpbDoc struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Size string `json:"size"`
	Hash string `json:"hash"`
	Type string `json:"type"`
}

func buildIndexMapping() *bleve.IndexMapping {
	// a generic reusable mapping for english text
	englishTextFieldMapping := bleve.NewTextFieldMapping()
	englishTextFieldMapping.Analyzer = "en"

	// a generic reusable mapping for keyword text
	keywordFieldMapping := bleve.NewTextFieldMapping()
	keywordFieldMapping.Analyzer = "keyword"

	torrentMapping := bleve.NewDocumentMapping()
	torrentMapping.AddFieldMappingsAt("name", englishTextFieldMapping)

	im := bleve.NewIndexMapping()

	im.AddDocumentMapping("torrent", torrentMapping)

	im.TypeField = "type"
	im.DefaultAnalyzer = "en"

	return im
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
	reader.FieldsPerRecord = 6
	reader.Comma = '|'

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			continue
		}

		batch.Index(record[0], tpbDoc{record[0], record[1], record[2], record[5], "torrent"})
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

		if count >= 5000 {
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
