package down

import (
	log "github.com/sirupsen/logrus"
	"io"
	"lime/pkg/down/site"
	"time"
)

var (
	tSleep   time.Duration
	errSleep time.Duration
)

func Job(syncStore *SyncStore, jobch chan error) {
	defer func(jobch chan error) {
		jobch <- io.EOF
	}(jobch)
	for {
		vi, ci, BookURL, err := syncStore.GetJob()
		if err != nil {
			if err != io.EOF {
				jobch <- err
			}
			return
		}

	LoopStep:
		for {
			content, err := site.Chapter(BookURL)
			if err != nil {
				log.Printf("Error: %s", err)
				time.Sleep(errSleep)
				continue LoopStep
			}
			syncStore.SaveJob(vi, ci, content)
			jobch <- nil
			time.Sleep(tSleep)
			break LoopStep
		}
	}
}
