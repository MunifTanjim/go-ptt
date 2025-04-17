package ptt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizer(t *testing.T) {
	for _, tc := range []struct {
		name   string
		input  Result
		output Result
	}{
		{"audio AC3", Result{Audio: []string{"AC3"}}, Result{Audio: []string{"DD"}}},
		{"audio EAC3", Result{Audio: []string{"EAC3"}}, Result{Audio: []string{"DDP"}}},
		{"audio dedupe", Result{Audio: []string{"AC3", "EAC3", "DD", "DDP"}}, Result{Audio: []string{"DD", "DDP"}}},

		{"codec h264", Result{Codec: "avc"}, Result{Codec: "AVC"}},
		{"codec h264", Result{Codec: "h264"}, Result{Codec: "AVC"}},
		{"codec x264", Result{Codec: "x264"}, Result{Codec: "AVC"}},
		{"codec h265", Result{Codec: "hevc"}, Result{Codec: "HEVC"}},
		{"codec h265", Result{Codec: "h265"}, Result{Codec: "HEVC"}},
		{"codec x265", Result{Codec: "x265"}, Result{Codec: "HEVC"}},

		{"resolution 2160p", Result{Resolution: "2160p"}, Result{Resolution: "4k"}},
		{"resolution 1440p", Result{Resolution: "1440p"}, Result{Resolution: "2k"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.input.is_normalized, false)
			tc.output.is_normalized = true
			assert.Equal(t, tc.input.Normalize(), &tc.output)
			assert.Equal(t, tc.input.is_normalized, true)
		})
	}
}
