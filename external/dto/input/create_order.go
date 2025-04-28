package input

type CreateOrder struct {
	Rate   float64
	Amount float64

	/*
		"sell"
		"buy"
	*/
	OrderType string

	/*
		"btc_jpy"
		"eth_jpy"
		"etc_jpy"
		"lsk_jpy"
		"xrp_jpy"
		"xem_jpy"
		"bch_jpy"
		"mona_jpy"
		"iost_jpy"
		"enj_jpy"
		"chz_jpy"
		"imx_jpy"
		"shib_jpy"
		"avax_jpy"
		"plt_jpy"
		"fnct_jpy"
		"dai_jpy"
		"wbtc_jpy"
		"bril_jpy"
		"bc_jpy"
		"doge_jp"
	*/
	Pair string
}
