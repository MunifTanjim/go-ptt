package ptt

import "strings"

func normalize_audio(audio []string) []string {
	isChanged := false
	for i := range audio {
		switch audio[i] {
		case "AC3":
			audio[i] = "DD"
			isChanged = true
		case "EAC3":
			audio[i] = "DDP"
			isChanged = true
		}
	}
	if !isChanged {
		return audio
	}
	nAudio := []string{}
	seenMap := map[string]struct{}{}
	for _, item := range audio {
		if _, seen := seenMap[item]; !seen {
			nAudio = append(nAudio, item)
			seenMap[item] = struct{}{}
		}
	}
	return nAudio
}

func normalize_codec(codec string) string {
	codec = strings.ToLower(codec)
	switch codec {
	case "h264", "x264":
		return "avc"
	case "h265", "x265":
		return "hevc"
	default:
		return codec
	}
}

func normalize_resolution(resolution string) string {
	resolution = strings.ToLower(resolution)
	switch resolution {
	case "2160p":
		return "4k"
	case "1440p":
		return "2k"
	default:
		return resolution
	}
}

func (r *Result) Normalize() *Result {
	if r.Error() != nil {
		return r
	}
	if !r.is_normalized {
		r.Audio = normalize_audio(r.Audio)
		r.Codec = normalize_codec(r.Codec)
		r.Resolution = normalize_resolution(r.Resolution)
		r.is_normalized = true
	}
	return r
}
