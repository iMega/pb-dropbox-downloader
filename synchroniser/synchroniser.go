package synchroniser

import (
	"fmt"
	"io"
	"sync/atomic"

	"github.com/go-git/go-billy/v5"
)

// DropboxSynchroniser Dropbox data synchroniser app structure.
type DropboxSynchroniser struct {
	progress       *Progress
	storage        DataStorage
	files          billy.Filesystem
	dropbox        Dropbox
	maxParallelism int
	output         io.Writer
	version        string
}

// NewSynchroniser creates and initialize new instance of DropboxSynchroniser create.
func NewSynchroniser(options ...synchroniserOption) *DropboxSynchroniser {
	s := &DropboxSynchroniser{
		maxParallelism: 1,
		output:         io.Discard,
		progress:       &Progress{},
	}

	for _, option := range options {
		option(s)
	}

	return s
}

func (s *DropboxSynchroniser) printf(format string, a ...interface{}) {
	fmt.Fprintln(s.output, fmt.Sprintf(format, a...))
}

type Progress struct {
	fn      ProgressFn
	current uint32
	total   int
}

func (p *Progress) SetTotal(val int) {
	p.total = val
}

func (p *Progress) Increase(text string) {
	atomic.AddUint32(&p.current, 1)

	p.fn(text, int(p.current), p.total)
}

type ProgressFn func(text string, current, total int)

func WithProgress(fn ProgressFn) synchroniserOption {
	return func(ds *DropboxSynchroniser) {
		ds.progress = &Progress{fn: fn}
	}
}
