package balance

import (
	"context"
	"net/http"

	http_client "github.com/sg3t41/go-coincheck/internal/client/http"
)

// GetResponseV2 はアカウントの残高情報を持つ構造体（全通貨対応版）
type GetResponseV2 struct {
	Success bool `json:"success"`
	
	// JPY
	JPY          string `json:"jpy"`
	JPYReserved  string `json:"jpy_reserved"`
	JPYLending   string `json:"jpy_lending"`
	JPYLendInUse string `json:"jpy_lend_in_use"`
	JPYLent      string `json:"jpy_lent"`
	JPYDebt      string `json:"jpy_debt"`
	JPYTsumitate string `json:"jpy_tsumitate"`
	
	// BTC
	BTC          string `json:"btc"`
	BTCReserved  string `json:"btc_reserved"`
	BTCLending   string `json:"btc_lending"`
	BTCLendInUse string `json:"btc_lend_in_use"`
	BTCLent      string `json:"btc_lent"`
	BTCDebt      string `json:"btc_debt"`
	BTCTsumitate string `json:"btc_tsumitate"`
	
	// ETH
	ETH          string `json:"eth,omitempty"`
	ETHReserved  string `json:"eth_reserved,omitempty"`
	ETHLending   string `json:"eth_lending,omitempty"`
	ETHLendInUse string `json:"eth_lend_in_use,omitempty"`
	ETHLent      string `json:"eth_lent,omitempty"`
	ETHDebt      string `json:"eth_debt,omitempty"`
	ETHTsumitate string `json:"eth_tsumitate,omitempty"`
	
	// ETC
	ETC          string `json:"etc,omitempty"`
	ETCReserved  string `json:"etc_reserved,omitempty"`
	ETCLending   string `json:"etc_lending,omitempty"`
	ETCLendInUse string `json:"etc_lend_in_use,omitempty"`
	ETCLent      string `json:"etc_lent,omitempty"`
	ETCDebt      string `json:"etc_debt,omitempty"`
	ETCTsumitate string `json:"etc_tsumitate,omitempty"`
	
	// LSK
	LSK          string `json:"lsk,omitempty"`
	LSKReserved  string `json:"lsk_reserved,omitempty"`
	LSKLending   string `json:"lsk_lending,omitempty"`
	LSKLendInUse string `json:"lsk_lend_in_use,omitempty"`
	LSKLent      string `json:"lsk_lent,omitempty"`
	LSKDebt      string `json:"lsk_debt,omitempty"`
	LSKTsumitate string `json:"lsk_tsumitate,omitempty"`
	
	// XRP
	XRP          string `json:"xrp,omitempty"`
	XRPReserved  string `json:"xrp_reserved,omitempty"`
	XRPLending   string `json:"xrp_lending,omitempty"`
	XRPLendInUse string `json:"xrp_lend_in_use,omitempty"`
	XRPLent      string `json:"xrp_lent,omitempty"`
	XRPDebt      string `json:"xrp_debt,omitempty"`
	XRPTsumitate string `json:"xrp_tsumitate,omitempty"`
	
	// XEM
	XEM          string `json:"xem,omitempty"`
	XEMReserved  string `json:"xem_reserved,omitempty"`
	XEMLending   string `json:"xem_lending,omitempty"`
	XEMLendInUse string `json:"xem_lend_in_use,omitempty"`
	XEMLent      string `json:"xem_lent,omitempty"`
	XEMDebt      string `json:"xem_debt,omitempty"`
	XEMTsumitate string `json:"xem_tsumitate,omitempty"`
	
	// LTC
	LTC          string `json:"ltc,omitempty"`
	LTCReserved  string `json:"ltc_reserved,omitempty"`
	LTCLending   string `json:"ltc_lending,omitempty"`
	LTCLendInUse string `json:"ltc_lend_in_use,omitempty"`
	LTCLent      string `json:"ltc_lent,omitempty"`
	LTCDebt      string `json:"ltc_debt,omitempty"`
	LTCTsumitate string `json:"ltc_tsumitate,omitempty"`
	
	// BCH
	BCH          string `json:"bch,omitempty"`
	BCHReserved  string `json:"bch_reserved,omitempty"`
	BCHLending   string `json:"bch_lending,omitempty"`
	BCHLendInUse string `json:"bch_lend_in_use,omitempty"`
	BCHLent      string `json:"bch_lent,omitempty"`
	BCHDebt      string `json:"bch_debt,omitempty"`
	BCHTsumitate string `json:"bch_tsumitate,omitempty"`
	
	// MONA
	MONA          string `json:"mona,omitempty"`
	MONAReserved  string `json:"mona_reserved,omitempty"`
	MONALending   string `json:"mona_lending,omitempty"`
	MONALendInUse string `json:"mona_lend_in_use,omitempty"`
	MONALent      string `json:"mona_lent,omitempty"`
	MONADebt      string `json:"mona_debt,omitempty"`
	MONATsumitate string `json:"mona_tsumitate,omitempty"`
	
	// XLM
	XLM          string `json:"xlm,omitempty"`
	XLMReserved  string `json:"xlm_reserved,omitempty"`
	XLMLending   string `json:"xlm_lending,omitempty"`
	XLMLendInUse string `json:"xlm_lend_in_use,omitempty"`
	XLMLent      string `json:"xlm_lent,omitempty"`
	XLMDebt      string `json:"xlm_debt,omitempty"`
	XLMTsumitate string `json:"xlm_tsumitate,omitempty"`
	
	// QTUM
	QTUM          string `json:"qtum,omitempty"`
	QTUMReserved  string `json:"qtum_reserved,omitempty"`
	QTUMLending   string `json:"qtum_lending,omitempty"`
	QTUMLendInUse string `json:"qtum_lend_in_use,omitempty"`
	QTUMLent      string `json:"qtum_lent,omitempty"`
	QTUMDebt      string `json:"qtum_debt,omitempty"`
	QTUMTsumitate string `json:"qtum_tsumitate,omitempty"`
	
	// BAT
	BAT          string `json:"bat,omitempty"`
	BATReserved  string `json:"bat_reserved,omitempty"`
	BATLending   string `json:"bat_lending,omitempty"`
	BATLendInUse string `json:"bat_lend_in_use,omitempty"`
	BATLent      string `json:"bat_lent,omitempty"`
	BATDebt      string `json:"bat_debt,omitempty"`
	BATTsumitate string `json:"bat_tsumitate,omitempty"`
	
	// IOST
	IOST          string `json:"iost,omitempty"`
	IOSTReserved  string `json:"iost_reserved,omitempty"`
	IOSTLending   string `json:"iost_lending,omitempty"`
	IOSTLendInUse string `json:"iost_lend_in_use,omitempty"`
	IOSTLent      string `json:"iost_lent,omitempty"`
	IOSTDebt      string `json:"iost_debt,omitempty"`
	IOSTTsumitate string `json:"iost_tsumitate,omitempty"`
	
	// ENJ
	ENJ          string `json:"enj,omitempty"`
	ENJReserved  string `json:"enj_reserved,omitempty"`
	ENJLending   string `json:"enj_lending,omitempty"`
	ENJLendInUse string `json:"enj_lend_in_use,omitempty"`
	ENJLent      string `json:"enj_lent,omitempty"`
	ENJDebt      string `json:"enj_debt,omitempty"`
	ENJTsumitate string `json:"enj_tsumitate,omitempty"`
	
	// OMG
	OMG          string `json:"omg,omitempty"`
	OMGReserved  string `json:"omg_reserved,omitempty"`
	OMGLending   string `json:"omg_lending,omitempty"`
	OMGLendInUse string `json:"omg_lend_in_use,omitempty"`
	OMGLent      string `json:"omg_lent,omitempty"`
	OMGDebt      string `json:"omg_debt,omitempty"`
	OMGTsumitate string `json:"omg_tsumitate,omitempty"`
	
	// PLT
	PLT          string `json:"plt,omitempty"`
	PLTReserved  string `json:"plt_reserved,omitempty"`
	PLTLending   string `json:"plt_lending,omitempty"`
	PLTLendInUse string `json:"plt_lend_in_use,omitempty"`
	PLTLent      string `json:"plt_lent,omitempty"`
	PLTDebt      string `json:"plt_debt,omitempty"`
	PLTTsumitate string `json:"plt_tsumitate,omitempty"`
	
	// SAND
	SAND          string `json:"sand,omitempty"`
	SANDReserved  string `json:"sand_reserved,omitempty"`
	SANDLending   string `json:"sand_lending,omitempty"`
	SANDLendInUse string `json:"sand_lend_in_use,omitempty"`
	SANDLent      string `json:"sand_lent,omitempty"`
	SANDDebt      string `json:"sand_debt,omitempty"`
	SANDTsumitate string `json:"sand_tsumitate,omitempty"`
	
	// DOT
	DOT          string `json:"dot,omitempty"`
	DOTReserved  string `json:"dot_reserved,omitempty"`
	DOTLending   string `json:"dot_lending,omitempty"`
	DOTLendInUse string `json:"dot_lend_in_use,omitempty"`
	DOTLent      string `json:"dot_lent,omitempty"`
	DOTDebt      string `json:"dot_debt,omitempty"`
	DOTTsumitate string `json:"dot_tsumitate,omitempty"`
	
	// FNCT
	FNCT          string `json:"fnct,omitempty"`
	FNCTReserved  string `json:"fnct_reserved,omitempty"`
	FNCTLending   string `json:"fnct_lending,omitempty"`
	FNCTLendInUse string `json:"fnct_lend_in_use,omitempty"`
	FNCTLent      string `json:"fnct_lent,omitempty"`
	FNCTDebt      string `json:"fnct_debt,omitempty"`
	FNCTTsumitate string `json:"fnct_tsumitate,omitempty"`
	
	// CHZ
	CHZ          string `json:"chz,omitempty"`
	CHZReserved  string `json:"chz_reserved,omitempty"`
	CHZLending   string `json:"chz_lending,omitempty"`
	CHZLendInUse string `json:"chz_lend_in_use,omitempty"`
	CHZLent      string `json:"chz_lent,omitempty"`
	CHZDebt      string `json:"chz_debt,omitempty"`
	CHZTsumitate string `json:"chz_tsumitate,omitempty"`
	
	// LINK
	LINK          string `json:"link,omitempty"`
	LINKReserved  string `json:"link_reserved,omitempty"`
	LINKLending   string `json:"link_lending,omitempty"`
	LINKLendInUse string `json:"link_lend_in_use,omitempty"`
	LINKLent      string `json:"link_lent,omitempty"`
	LINKDebt      string `json:"link_debt,omitempty"`
	LINKTsumitate string `json:"link_tsumitate,omitempty"`
	
	// MKR
	MKR          string `json:"mkr,omitempty"`
	MKRReserved  string `json:"mkr_reserved,omitempty"`
	MKRLending   string `json:"mkr_lending,omitempty"`
	MKRLendInUse string `json:"mkr_lend_in_use,omitempty"`
	MKRLent      string `json:"mkr_lent,omitempty"`
	MKRDebt      string `json:"mkr_debt,omitempty"`
	MKRTsumitate string `json:"mkr_tsumitate,omitempty"`
	
	// MATIC
	MATIC          string `json:"matic,omitempty"`
	MATICReserved  string `json:"matic_reserved,omitempty"`
	MATICLending   string `json:"matic_lending,omitempty"`
	MATICLendInUse string `json:"matic_lend_in_use,omitempty"`
	MATICLent      string `json:"matic_lent,omitempty"`
	MATICDebt      string `json:"matic_debt,omitempty"`
	MATICTsumitate string `json:"matic_tsumitate,omitempty"`
	
	// APE
	APE          string `json:"ape,omitempty"`
	APEReserved  string `json:"ape_reserved,omitempty"`
	APELending   string `json:"ape_lending,omitempty"`
	APELendInUse string `json:"ape_lend_in_use,omitempty"`
	APELent      string `json:"ape_lent,omitempty"`
	APEDebt      string `json:"ape_debt,omitempty"`
	APETsumitate string `json:"ape_tsumitate,omitempty"`
	
	// AXS
	AXS          string `json:"axs,omitempty"`
	AXSReserved  string `json:"axs_reserved,omitempty"`
	AXSLending   string `json:"axs_lending,omitempty"`
	AXSLendInUse string `json:"axs_lend_in_use,omitempty"`
	AXSLent      string `json:"axs_lent,omitempty"`
	AXSDebt      string `json:"axs_debt,omitempty"`
	AXSTsumitate string `json:"axs_tsumitate,omitempty"`
	
	// IMX
	IMX          string `json:"imx,omitempty"`
	IMXReserved  string `json:"imx_reserved,omitempty"`
	IMXLending   string `json:"imx_lending,omitempty"`
	IMXLendInUse string `json:"imx_lend_in_use,omitempty"`
	IMXLent      string `json:"imx_lent,omitempty"`
	IMXDebt      string `json:"imx_debt,omitempty"`
	IMXTsumitate string `json:"imx_tsumitate,omitempty"`
	
	// WBTC
	WBTC          string `json:"wbtc,omitempty"`
	WBTCReserved  string `json:"wbtc_reserved,omitempty"`
	WBTCLending   string `json:"wbtc_lending,omitempty"`
	WBTCLendInUse string `json:"wbtc_lend_in_use,omitempty"`
	WBTCLent      string `json:"wbtc_lent,omitempty"`
	WBTCDebt      string `json:"wbtc_debt,omitempty"`
	WBTCTsumitate string `json:"wbtc_tsumitate,omitempty"`
	
	// AVAX
	AVAX          string `json:"avax,omitempty"`
	AVAXReserved  string `json:"avax_reserved,omitempty"`
	AVAXLending   string `json:"avax_lending,omitempty"`
	AVAXLendInUse string `json:"avax_lend_in_use,omitempty"`
	AVAXLent      string `json:"avax_lent,omitempty"`
	AVAXDebt      string `json:"avax_debt,omitempty"`
	AVAXTsumitate string `json:"avax_tsumitate,omitempty"`
	
	// SHIB
	SHIB          string `json:"shib,omitempty"`
	SHIBReserved  string `json:"shib_reserved,omitempty"`
	SHIBLending   string `json:"shib_lending,omitempty"`
	SHIBLendInUse string `json:"shib_lend_in_use,omitempty"`
	SHIBLent      string `json:"shib_lent,omitempty"`
	SHIBDebt      string `json:"shib_debt,omitempty"`
	SHIBTsumitate string `json:"shib_tsumitate,omitempty"`
}

// GETV2 は全通貨対応版の残高取得
func (t balance) GETV2(ctx context.Context) (*GetResponseV2, error) {
	req, err := t.client.CreateRequest(ctx, http_client.RequestInput{
		Method:  http.MethodGet,
		Path:    "/api/accounts/balance",
		Private: true,
	})
	if err != nil {
		return nil, err
	}

	var res GetResponseV2
	if err := t.client.Do(req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}