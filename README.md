[![GitHub Workflow Status: CI](https://img.shields.io/github/actions/workflow/status/MunifTanjim/go-ptt/ci.yml?branch=main&label=CI&style=for-the-badge)](https://github.com/MunifTanjim/go-ptt/actions/workflows/ci.yml)
[![Go Reference](https://img.shields.io/badge/go-reference-%23007d9c?style=for-the-badge&logo=go)](https://pkg.go.dev/github.com/MunifTanjim/go-ptt)
[![License](https://img.shields.io/github/license/MunifTanjim/go-ptt?style=for-the-badge)](https://github.com/MunifTanjim/go-ptt/blob/main/LICENSE)

# go-ptt

Parse Torrent Title

## `Result` Reference

- **Audio** (`[]string`): `DTS Lossless`, `DTS Lossy`, `Atmos`, `TrueHD`, `FLAC`, `DDP`, `EAC3`, `DD`, `AC3`, `AAC`, `PCM`, `OPUS`, `HQ`, `MP3`
- **BitDepth** (`string`): `8bit`, `10bit`, `12bit`
- **Channels** (`[]string`): `2.0`, `5.1`, `7.1`, `stereo`, `mono`
- **Codec** (`string`): `AVC` (`avc`, `h264`, `x264`,), `HEVC` (`hevc`, `h265`, `x265`), `MPEG-2` (`mgpeg2`), `DivX` (`divx`, `dvix`), `Xvid` (`xvid`)
- **Commentary** (`bool`)
- **Complete** (`bool`)
- **Container** (`string`): `mkv`, `avi`, `mp4`, `wmv`, `mpg`, `mpeg`
- **Convert** (`bool`)
- **Date** (`string`): `YYYY-MM-DD`
- **Documentary** (`bool`)
- **Dubbed** (`bool`)
- **Edition** (`string`): `Anniversary Edition`, `Ultimate Edition`, `Director's Cut`, `Extended Edition`, `Collector's Edition`, `Theatrical`, `Uncut`, `IMAX`, `Diamond Edition`, `Remastered`
- **EpisodeCode** (`string`): 8-character alphanumeric code
- **Episodes** (`[]int`): Array of episode numbers
- **Extended** (`bool`)
- **Extension** (`string`): `3g2`, `3gp`, `avi`, `flv`, `mkv`, `mk3d`, `mov`, `mp2`, `mp4`, `m4v`, `mpe`, `mpeg`, `mpg`, `mpv`, `webm`, `wmv`, `ogm`, `divx`, `ts`, `m2ts`, `iso`, `vob`, `sub`, `idx`, `ttxt`, `txt`, `smi`, `srt`, `ssa`, `ass`, `vtt`, `nfo`, `html`
- **Group** (`string`): Release group name
- **HDR** (`[]string`): `DV`, `HDR10+`, `HDR`, `SDR`
- **Hardcoded** (`bool`)
- **Languages** (`[]string`):
  - Special: `multi subs`, `multi audio`, `dual audio`
  - ISO 639-1: `en`, `ja`, `ko`, `zh`, `zh-tw`, `fr`, `es`, `es-419`, `pt`, `it`, `de`, `ru`, `uk`, `nl`, `da`, `fi`, `sv`, `no`, `el`, `lt`, `lv`, `et`, `pl`, `cs`, `sk`, `hu`, `ro`, `bg`, `sr`, `hr`, `sl`, `hi`, `te`, `ta`, `ml`, `kn`, `mr`, `gu`, `pa`, `bn`, `vi`, `id`, `th`, `ms`, `ar`, `tr`, `he`, `fa`
- **Network** (`string`): `Apple TV`, `Amazon`, `Netflix`, `Nickelodeon`, `Disney`, `HBO`, `Hulu`, `CBS`, `NBC`, `AMC`, `PBS`, `Crunchyroll`, `VICE`, `Sony`, `Hallmark`, `Adult Swim`, `Animal Planet`, `Cartoon Network`
- **Proper** (`bool`)
- **Quality** (`string`):
  - Cam: `CAM`, `TeleSync`, `TeleCine`, `SCR`
  - Web: `WEB`, `WEB-DL`, `WEBRip`, `WEB-DLRip`
  - Broadcast: `HDTV`, `HDTVRip`, `PDTV`, `TVRip`, `SATRip`
  - Physical: `BluRay`, `BluRay REMUX`, `REMUX`, `BRRip`, `BDRip`, `UHDRip`, `HDRip`, `DVD`, `DVDRip`, `PPVRip`, `R5`
  - Other: `XviD`, `DivX`
- **Region** (`string`): `R0`-`R9`, `R2J`, `PAL`, `NTSC`, `SECAM`
- **ReleaseTypes** (`[]string`): `OAD` (`ODA`),`OVA`(`OAV`), `ONA`, 
- **Remastered** (bool)
- **Repack** (`bool`)
- **Resolution** (`string`): `4k` (`2160p`), `2k` (`1440p`), `1080p`, `720p`, `576p`, `480p`, `360p`, `240p`
- **Retail** (`bool`)
- **Seasons** (`[]int`): Array of season numbers
- **Site** (`string`): Source website/URL
- **Size** (`string`): Size with unit (e.g., `2.3GB`)
- **Subbed** (`bool`)
- **ThreeD** (`string`): `3D`, `3D HSBS`, `3D SBS`, `3D HOU`, `3D OU`
- **Title** (`string`): Cleaned title
- **Uncensored** (`bool`)
- **Unrated** (`bool`)
- **Upscaled** (`bool`)
- **Volumes** (`[]int`): Array of volume numbers
- **Year** (`string`): `YYYY` or `YYYY-YYYY`

## Acknowledgement

- [TheBeastLT/parse-torrent-title](https://github.com/TheBeastLT/parse-torrent-title)
- [dreulavelle/PTT](https://github.com/dreulavelle/PTT)

## License

Licensed under the MIT License. Check the [LICENSE](./LICENSE) file for details.
