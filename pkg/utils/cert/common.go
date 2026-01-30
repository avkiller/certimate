package cert

import (
	"encoding/pem"
)

func decodePEMBlocks(data []byte) []*pem.Block {
	blocks := make([]*pem.Block, 0)
	for {
		block, rest := pem.Decode(data)
		if block == nil {
			break
		}

		blocks = append(blocks, block)
		data = rest
	}

	return blocks
}
