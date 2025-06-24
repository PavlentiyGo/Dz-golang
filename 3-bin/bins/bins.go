package bins

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}
type BinsList struct {
	bins []Bin
}

func newBin(id string, private bool, name string) *Bin {
	return &Bin{id: id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}
func newBinsList(bin *Bin) *BinsList {
	return &BinsList{bins: []Bin{*bin}}
}
