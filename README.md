# video-enconder

This is a simple video encoder that uses ffmpeg to encode videos to a specific format.

The idea is to have a simple way to encode videos to a specific format, so uploading fragmented videos to google storage is easier using go routines and channels.

## Setup

To use this project you need to have ffmpeg installed on your machine or use a docker image.

### Install dependencies

```bash
go mod tidy
```

### Run tests

```bash
go test ./...
```
