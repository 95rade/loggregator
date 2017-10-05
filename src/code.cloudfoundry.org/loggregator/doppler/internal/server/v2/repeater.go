package v2

import (
	"code.cloudfoundry.org/loggregator/plumbing/v2"
)

// Repeater connects a reader to a writer.
type Repeater struct {
	r Reader
	w Writer
}

// Reader reads envelopes.
type Reader func() *loggregator_v2.Envelope

// Writer writes envelopes.
type Writer func(*loggregator_v2.Envelope)

// NewRepeater is the constructor for Transponder.
func NewRepeater(r Reader, w Writer) *Repeater {
	return &Repeater{
		r: r,
		w: w,
	}
}

// Start blocks indefinitely while transmitting data from the reader to the
// writer.
func (r *Repeater) Start() {
	for {
		r.w(r.r())
	}
}
