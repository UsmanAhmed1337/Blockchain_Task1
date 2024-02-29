package blockchain

import (
    "fmt"
    "time"
)

type Block struct {
    Index     int
    Timestamp time.Time
    Data      string
    PrevHash  string
    Hash      string
}

type Blockchain struct {
    Blocks []*Block
}

func NewBlock(index int, data, prevHash string) *Block {
    block := &Block{
        Index:     index,
        Timestamp: time.Now(),
        Data:      data,
        PrevHash:  prevHash,
        Hash:      "", 
    }
    return block
}

func (bc *Blockchain) DisplayAllBlocks() {
    fmt.Println("Blockchain:")
    for _, block := range bc.Blocks {
        fmt.Printf("Index: %d\n", block.Index)
        fmt.Printf("Timestamp: %s\n", block.Timestamp.String())
        fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("PrevHash: %s\n", block.PrevHash)
        fmt.Printf("Hash: %s\n\n", block.Hash)
    }
}

func (bc *Blockchain) ModifyBlock(index int, newData string) {
    if index >= 0 && index < len(bc.Blocks) {
        bc.Blocks[index].Data = newData
    } else {
        fmt.Println("Invalid block index")
    }
}
