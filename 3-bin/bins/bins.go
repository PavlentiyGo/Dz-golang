package bins

import "time"

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}
type BinList struct {
	bins []Bin
}

func newBin(id string, private bool, name string) *Bin {
	return &Bin{Id: id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}
func newBinList(bin *Bin) *BinList {
	return &BinList{bins: []Bin{*bin}}
}
