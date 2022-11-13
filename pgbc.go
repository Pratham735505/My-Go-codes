package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"time"
)

const (
	difficulty = 10
)

type Block struct {
	Index     int
	Timestamp int64
	Hash      []byte
	Data      string
	PrevHash  []byte
}

// bytes manipulation based on https://github.com/Jeiwan/blockchain_go

func genhash(data string, prev []byte) []byte {
	head := bytes.Join([][]byte{prev, []byte(data)}, []byte{})

	h32 := sha256.Sum256(head)

	fmt.Printf("Header hash: %x\n", h32)

	return h32[:]
}

func mine(hash []byte) []byte {
	target := big.NewInt(1)
	target = target.Lsh(target, uint(256-difficulty))

	fmt.Printf("target: %x\n", target)

	var nonce int64

	for nonce = 0; nonce < math.MaxInt64; nonce++ {
		testNum := big.NewInt(0)
		testNum.Add(testNum.SetBytes(hash), big.NewInt(nonce))
		testHash := sha256.Sum256(testNum.Bytes())

		fmt.Printf("\rproof: %x (nonce: %d)", testHash, nonce)

		if target.Cmp(testNum.SetBytes(testHash[:])) > 0 {
			fmt.Println("\nFound!")
			return testHash[:]
		}
	}

	return []byte{}
}

func NewBlock(id int, data string, prev []byte) *Block {
	return &Block{
		id,
		time.Now().Unix(),
		mine(genhash(data, prev)),
		data,
		prev,
	}
}

func main() {
	bdatas := []string{"\n 1\n Genesis\n   Name:Pratham Gupta\n    Email:prathamgupta735505@gmail.com\n    Phno:7355057737\n    Course:BCA\n",
		"\n 2\n Block2\n   Name:Saksham Gupta\n    Email:guptasaksham@gmail.com\n    Phno:697315186\n    Course:BTech\n",
		"\n 3\n Block3\n   Name:Ayushman Tripathi\n    Email:at2347@gmail.com\n    Phno:9456183465\n    Course:BA\n",
		"\n 4\n Block3\n   Name:Srajan Saxena\n    Email:srajansaxena54356@gmail.com\n    Phno:9461853225\n    Course:BA\n",
		"\n 5\n Block4\n   Name:Anurag Gupta\n    Email:Beinganurag23@gmail.com\n    Phno:9451223856\n    Course:BCA\n",
		"\n 6\n Block5\n   Name:Kushagra Dixit\n    Email:Kushagradixit018@gmail.com\n    Phno:6387022930\n   Course:BCA\n",
                "\n 5\n Block6\n   Name:Shivansh Srivastav\n   Email:creamwalabiscuit@gmail.com\n  Phno.:9305255488\n   Course:BCA"}

	prev := []byte{}

	for i, d := range bdatas {
		b := NewBlock(i, d, prev)
		fmt.Printf("Id: %d\nHash; %x\nData: %s\nPrevious: %x\n",
			b.Index,
			b.Hash,
			b.Data,
			b.PrevHash,
		)
		prev = b.Hash
	}
}
