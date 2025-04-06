package ptt

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestUncensored(t *testing.T) {
	for _, tc := range []struct {
		ttitle     string
		uncensored bool
	}{
		{"Identity.Thief.2013.Vostfr.UNRATED.BluRay.720p.DTS.x264-Nenuko", false},
		{"Charlie.les.filles.lui.disent.merci.2007.UNCENSORED.TRUEFRENCH.DVDRiP.AC3.Libe", true},
		{"Have I Got News For You S53E02 EXTENDED 720p HDTV x264-QPEL", false},
	} {
		t.Run(tc.ttitle, func(t *testing.T) {
			result := Parse(tc.ttitle)
			assert.Equal(t, tc.uncensored, result.Uncensored)
		})
	}
}
