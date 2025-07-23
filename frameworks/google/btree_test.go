package the_google

import (
	"github.com/google/btree"
	"testing"
)

func Test_btree1(t *testing.T) {
	tr := btree.New(2)
	tr.ReplaceOrInsert(btree.Int(1))
	tr.ReplaceOrInsert(btree.Int(2))
	tr.ReplaceOrInsert(btree.Int(3))
	tr.Delete(btree.Int(2))
	tr.Ascend(func(a btree.Item) bool {
		t.Logf("%v", a)
		return true
	})
}
