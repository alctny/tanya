package plug

import (
	"ecila/util"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	URL_GODEV = `https://go.dev/dl/`

	OS_GODEV_MACOS   = `macOS`
	OS_GODEV_Linux   = `Linux`
	OS_GODEV_WINDOWS = `Windows`
	OS_GODEV_FREEBSD = `FreeBSD`

	ARCH_GODEV_X86     = `x86`
	ARCH_GODEV_X8664   = `x86-64`
	ARCH_GODEV_ARM64   = "ARM64"
	ARCH_GODEV_PPC64LE = "ppc64le"
	ARCH_GODEV_S390X   = "s390x"

	KIND_GODEV_INSTALER = "Installer"
	KIND_GODEV_ARCHIVE  = "Archive"
	KIND_GODEV_SOURCE   = "Source"

	COLUMN_FILENAME_GODEV = 1
	COLUMN_KIND_GODEV     = 2
	COLUMN_OS_GODEV       = 3
	COLUMN_ARCH_GODEV     = 4
	COLUMN_SIZE_GODEV     = 5
	COLUMN_SHA_GODEV      = 6

	// BASEURL_GODEV = "dl.google.com/go/go1.19.5.darwin-arm64.tar.gz"
)

func DownloadGolangSDK(os string, arch string, kind string) error {
	resp, err := http.Get(URL_GODEV)
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
		if strings.TrimSpace(row[COLUMN_KIND_GODEV]) != kind ||
			strings.TrimSpace(row[COLUMN_OS_GODEV]) != os ||
			strings.TrimSpace(row[COLUMN_ARCH_GODEV]) != arch {
			return
		}
		// url, ok := s.Find("a").Attr("href")
		// if ok {
		filename := strings.TrimSpace(row[COLUMN_FILENAME_GODEV])
		err = util.Download(fmt.Sprintf("https://dl.google.com/go/%s", filename), filename)
		if err != nil {
			panic(err)
		}
		safe, err := util.Hash64Sum(filename, strings.TrimSpace(row[COLUMN_SHA_GODEV]))
		if err != nil {
			panic(err)
		}
		if !safe {
			// TODO: 主动删除
			panic("the sha256 check not ok, please delete it")
		}
		// }
	})

	return nil
}

func InstallGolangSDK() {

}
