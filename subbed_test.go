package ptt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubbed(t *testing.T) {
	for _, tc := range []struct {
		ttitle string
		subbed bool
	}{
		{"The.Walking.Dead.S06E07.SUBFRENCH.HDTV.x264-AMB3R.mkv", true},
		{"www.MovCr.to - Bikram Yogi, Guru, Predator (2019) 720p WEB_DL x264 ESubs [Dual Audio]-[Hindi + Eng] - 950MB - MovCr.mkv", true},
		{"[www.1TamilMV.pics]_The.Great.Indian.Suicide.2023.Tamil.TRUE.WEB-DL.4K.SDR.HEVC.(DD+5.1.384Kbps.&.AAC).3.2GB.ESub.mkv", true},
		{"Evil.Dead.Rise.2023.PLSUB.720p.MA.WEB-DL.H264.E-AC3-CMRG", true},
		{"You.[Uncut].S01.SweSub.1080p.x264-Justiso", true},
		{"A.Good.Day.To.Die.Hard.2013.SWESUB.DANSUB.FiNSUB.720p.WEB-DL.-Ro", true},
		{"Seinfeld.COMPLETE.SLOSUBS.DVDRip.XviD", true},

		{"[HorribleSubs] White Album 2 - 06 [1080p].mkv", false},
	} {
		t.Run(tc.ttitle, func(t *testing.T) {
			result := Parse(tc.ttitle)
			assert.Equal(t, tc.subbed, result.Subbed)
		})
	}
}
