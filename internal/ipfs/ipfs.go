package ipfs

import (
	"bytes"
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/redish101/reblog/internal/env"
)

var sh *shell.Shell

func Init() {
	sh = shell.NewShell(env.IPFSAPIURL)
}

func UploadToIPFS(content []byte) (string, error) {
	if sh == nil {
		return "", fmt.Errorf("IPFS shell not initialized")
	}
	reader := bytes.NewReader(content)
	cid, err := sh.Add(reader)
	if err != nil {
		return "", fmt.Errorf("failed to add content to IPFS: %w", err)
	}
	return cid, nil
}
