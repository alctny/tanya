package main

import "ecila/plug"

func main() {
	err := plug.DownloadGolangSDK(plug.OS_GODEV_MACOS, plug.ARCH_GODEV_X8664, plug.KIND_GODEV_ARCHIVE)
	if err != nil {
		panic(err)
	}
}
