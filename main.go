package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/cbergoon/merkletree"
)

type transaction struct {
	Hash string `json:"hash"`
}

//CalculateHash hashes the values of a TestContent
func (t transaction) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(t.Hash)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

//Equals tests for equality of two Contents
func (t transaction) Equals(other merkletree.Content) (bool, error) {
	return t.Hash == other.(transaction).Hash, nil
}

func main() {

	n := flag.Int("n", 1000, "Number of leaf hash to generate")
	flag.Parse()

	txList := make([]merkletree.Content, 0)
	for i := 0; i < *n; i++ {
		sha := sha256.New()
		sha.Write([]byte(strconv.Itoa(i)))
		hash := hex.EncodeToString(sha.Sum(nil))

		txList = append(txList, transaction{
			Hash: hash,
		})
	}

	timer := time.Now()

	t, _ := merkletree.NewTree(txList)

	elapsed := time.Since(timer)

	log.Printf("Merkle tree computed for %d hash in %f seconds\n", *n, elapsed.Seconds())

	timer = time.Now()
	t.VerifyTree()
	elapsed = time.Since(timer)
	log.Printf("Verifying all the tree %f seconds\n", elapsed.Seconds())

	rand.Seed(time.Now().UTC().UnixNano())
	r := rand.Intn(*n)

	timer = time.Now()

	t.VerifyContent(txList[r])

	elapsed = time.Since(timer)

	log.Printf("Verifying a transaction inside the tree in %f seconds\n", elapsed.Seconds())
}
