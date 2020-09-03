package args

import "github.com/astaxie/beego/validation"

type VoteBalanceRsp struct {
	BalanceTotal  float64 `json:"balance_total"`
	VoteTotal     float64 `json:"vote_total"`
	EarningsTotal float64 `json:"earnings_total"`
}

type VoteCoinDepositListArgs struct {
	UserId   int
	Symbol   int `form:"symbol" json:"symbol"`
	PageSize int `form:"page_size" json:"page_size"`
	PageNum  int `form:"page_num" json:"page_num"`
}

func (t *VoteCoinDepositListArgs) Valid(v *validation.Validation) {
	if t.PageNum < 1 {
		v.SetError("PageNum", "PageNum 需要大于等于1")
	}
	if t.PageSize <= 0 {
		v.SetError("PageSize", "PageSize 需要大于0")
	}
	if t.Symbol != 1 && t.Symbol != 2 {
		v.SetError("Symbol", "Symbol 只能为1或2")
	}
}
