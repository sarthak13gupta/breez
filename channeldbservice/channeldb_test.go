package channeldbservice

import (
	"encoding/hex"
	"fmt"
	"github.com/breez/breez/channeldbservice"
	bolt "go.etcd.io/bbolt"
	"os"
	"path/filepath"
	"testing"
)

// To run this test, you need to set the LND_DIR environment variables.
func TestRemoveBadElement(t *testing.T) {
	workingDir := os.Getenv("LND_DIR")
	channeldb := "channel.db"

	dbPath := filepath.Join(workingDir, channeldb)
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		t.Fatalf(fmt.Sprintf("opening database received error: %v", err), "ERROR")
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("graph-edge"))
		if err != nil {
			t.Fatalf("CreateBucketIfNotExists failed: %v", err)
		}
		hex, _ := hex.DecodeString("02fe80fb6a2dc0fb6e9bec49c76d048889c91355d4e900fcb026bf095665790325")
		err = b.Put(hex, []byte{})
		if err != nil {
			t.Fatalf("put failed: %v", err)
		}
		return nil
	})
	db.Close()
	if err != nil {
		t.Fatalf("updating db failed: %v", err)
	}
	_, chanDBCleanUp, err := channeldbservice.Get(workingDir)
	if err != nil {
		t.Fatalf("updating db failed: %v", err)
	}
	defer chanDBCleanUp()
}
