package plug

type ArchType string
type OSType string
type Bit string
type FileType string

type Config struct {
	Arch ArchType
	OS   OSType
	Bit  Bit
	Kind FileType
}

const (
	DOWNLOAD_PATH = "download_list.txt"
	SHA_FILE      = "sha.txt"
)

const (
	ARCH_AMD ArchType = "amd"
	ARCH_ARM ArchType = "arm"
)

const (
	OS_MAC OSType = "mac"
)

const (
	BIT_32 Bit = "32"
	BIT_64 Bit = "64"
)
