package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

type blockchain struct {
	blocks []*Block

	/**
	blocks []block과 같은 형태로 슬라이스를 선언하면,
	슬라이스가 참조하는 배열이 복사되어 메모리 사용량이 늘어나게 됩니다.
	이 때문에, 슬라이스가 매우 길어지는 경우에는 메모리 사용량이 매우 증가하게 되어 문제가 발생할 수 있습니다.

	따라서, blocks []*block과 같이 포인터 슬라이스로 변경하면,
	실제 block 객체가 아니라 객체의 '메모리 주소값'만 저장하게 되므로, 메모리 사용량이 줄어들게 됩니다.

	이렇게 하면 인스턴스의 복사가 아니라
	인스턴스의 메모리 주소값을 가져와서 메모리를 안정화시키며 최적화에 이점이 있습니다.
	*/
}

var b *blockchain  // singleton 패턴
var once sync.Once // 한번만 사용하기 위함

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLatestHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks > 0 {
		return GetBlockchain().blocks[totalBlocks-1].Hash
	}
	return ""
}

func createBlock(data string) *Block { // 메모리 주소 반환!
	newBlock := &Block{data, "", getLatestHash(), len(GetBlockchain().blocks) + 1}
	newBlock.calculateHash()
	return newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis Block")
		}) // 단 한번만 실행할수 있도록 (병렬의 상황에서도.)
	}
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}

func (b *blockchain) GetBlock(height int) *Block {
	return b.blocks[height-1]
}
