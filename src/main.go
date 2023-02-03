package main

import (
	"github.com/lwifew/aded/plug"
)

func main() {
	plug.InitGolangPlug(
		plug.ARM64,
		plug.MACOS,
		plug.BIT_64,
		plug.ARCHIVE_GO,
	)
	err := plug.GoPlug.Parse()
	if err != nil {
		panic(err)
	}

	// https://code.visualstudio.com/sha/download?build=stable&os=linux-x64
	// alice's development envrioument downloader
}
