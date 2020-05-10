package down

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io"
	"io/ioutil"
	"lime/pkg/down/store"
	"sync"
)

type SyncStore struct {
	lock    *sync.Mutex
	IsTWork bool
	jobs    [][]bool
	Store   *store.Store
}

func (s *SyncStore) Init() {
	s.jobs = make([][]bool, len(s.Store.Volumes))
	s.lock = &sync.Mutex{}
	for k := range s.jobs {
		s.jobs[k] = make([]bool, len(s.Store.Volumes[k].Chapters))
	}
}

func (s *SyncStore) GetJob() (vi, ci int, url string, err error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for vi, vol := range s.Store.Volumes {
		for ci, ch := range vol.Chapters {
			if !s.jobs[vi][ci] {
				if len(ch.Text) == 0 {
					s.jobs[vi][ci] = true
					return vi, ci, ch.URL, nil
				}
			}
		}
	}
	return 0, 0, "", io.EOF
}

func (s *SyncStore) SaveJob(vi, ci int, text []string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Store.Volumes[vi].Chapters[ci].Text = text
	bbb, err := yaml.Marshal(*(s.Store))
	if err != nil {
		panic(err)
	}
	var filename = fmt.Sprintf("%s.%s", s.Store.BookName, store.FileExt)
	err = ioutil.WriteFile(filename, bbb, 0775)
	if err != nil {
		panic(err)
	}
}
