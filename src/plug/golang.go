package plug

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	URL = `https://go.dev/dl/`
	// BASEURL_GODEV = "dl.google.com/go/go1.19.5.darwin-arm64.tar.gz"
)

const (
	X86     = "x86"
	X8664   = "x86-64"
	ARM64   = "ARM64"
	PPC64LE = "ppc64le"
	S390X   = "s390x"
)

const (
	MACOS   = "macOS"
	Linux   = "Linux"
	WINDOWS = "Windows"
	FREEBSD = "FreeBSD"
)

const (
	FILENAME = 1
	KIND     = 2
	OS       = 3
	ARCH     = 4
	SIZE     = 5
	SHA      = 6
)

const (
	INSTALER_GO = "Installer"
	ARCHIVE_GO  = "Archive"
	SOURCE_GO   = "Source"
)

type golangPlug struct {
	Config
	Link string
	File string
	Sha  string
}

var GoPlug = &golangPlug{}

func InitGolangPlug(arch ArchType, os OSType, bit Bit, kind FileType) {
	GoPlug.Config = Config{
		Arch: arch,
		OS:   os,
		Bit:  bit,
		Kind: kind,
	}
}

func (g *golangPlug) Parse() error {
	resp, err := http.Get(URL)
	if err != nil {
		return err
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	doc.Find(".toggleVisible").First().Find("tr").Each(func(i int, s *goquery.Selection) {
		row := strings.Split(s.Text(), "\n")
		if len(row) != 8 {
			return
		}
		kind := strings.TrimSpace(row[KIND])
		os := strings.TrimSpace(row[OS])
		arch := strings.TrimSpace(row[ARCH])
		if kind != string(g.Kind) || os != string(g.OS) || arch != string(g.Arch) {
			return
		}
		link, exist := s.Find(".download").Attr("href")
		if !exist {
			panic("parse error")
		}
		GoPlug.File = row[FILENAME]
		GoPlug.Sha = row[SHA]
		GoPlug.Link = link
	})

	return nil
}

// TODO
// Install install go sdk to local
func (g *golangPlug) Install() {

}
