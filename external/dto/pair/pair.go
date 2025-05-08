package pair

import "fmt"

type Code int

const (
	PairBTCJPY Code = iota
	PairETHJPY
	PairETCJPY
	PairLSKJPY
	PairXRPJPY
	PairXEMJPY
	PairBCHJPY
	PairMONAJPY
	PairIOSTJPY
	PairENJJPY
	PairCHZJPY
	PairIMXJPY
	PairSHIBJPY
	PairAVAXJPY
	PairPLTJPY
	PairFNCTJPY
	PairDAIJPY
	PairWBTCJPY
	PairBRILJPY
	PairBCJPY
	PairDOGEJPY
)

var pairMap = map[string]Code{
	"btc_jpy":  PairBTCJPY,
	"eth_jpy":  PairETHJPY,
	"etc_jpy":  PairETCJPY,
	"lsk_jpy":  PairLSKJPY,
	"xrp_jpy":  PairXRPJPY,
	"xem_jpy":  PairXEMJPY,
	"bch_jpy":  PairBCHJPY,
	"mona_jpy": PairMONAJPY,
	"iost_jpy": PairIOSTJPY,
	"enj_jpy":  PairENJJPY,
	"chz_jpy":  PairCHZJPY,
	"imx_jpy":  PairIMXJPY,
	"shib_jpy": PairSHIBJPY,
	"avax_jpy": PairAVAXJPY,
	"plt_jpy":  PairPLTJPY,
	"fnct_jpy": PairFNCTJPY,
	"dai_jpy":  PairDAIJPY,
	"wbtc_jpy": PairWBTCJPY,
	"bril_jpy": PairBRILJPY,
	"bc_jpy":   PairBCJPY,
	"doge_jpy": PairDOGEJPY,
}

type Pair interface {
	Value() string
	Equal(Pair) bool
}

type pair struct {
	value Code
}

func New(s string) (Pair, error) {
	if code, ok := pairMap[s]; ok {
		return pair{value: code}, nil
	}
	return nil, fmt.Errorf("invalid pair string: %s", s)
}

func New2(code Code) (Pair, error) {
	return pair{value: code}, nil
}

func FromString(s string) (Pair, error) {
	if code, ok := pairMap[s]; ok {
		return pair{value: code}, nil
	}
	return nil, fmt.Errorf("invalid pair string: %s", s)
}

func (p pair) Value() string {
	for key, value := range pairMap {
		if value == p.value {
			return key
		}
	}
	return ""
}

func (p pair) Equal(other Pair) bool {
	if otherPair, ok := other.(pair); ok {
		return p.value == otherPair.value
	}
	return false
}
