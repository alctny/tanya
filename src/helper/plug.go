package helper

type Downloader interface {
	Link() string
	File() string
	ShaCode() string
}
