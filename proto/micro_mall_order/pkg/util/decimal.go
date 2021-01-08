package util

import (
	"github.com/shopspring/decimal"
)

func init() {
	// 小数精度，默认16
	decimal.DivisionPrecision = 16
}

// 小于等于
func DecimalLessThanOrEqual(d1 decimal.Decimal, d2 decimal.Decimal) bool {
	return d1.LessThanOrEqual(d2)
}

// 小于
func DecimalLessThan(d1 decimal.Decimal, d2 decimal.Decimal) bool {
	return d1.LessThan(d2)
}

// 大于等于
func DecimalGreaterThanOrEqual(d1 decimal.Decimal, d2 decimal.Decimal) bool {
	return d1.GreaterThanOrEqual(d2)
}

// 大于
func DecimalGreaterThan(d1 decimal.Decimal, d2 decimal.Decimal) bool {
	return d1.GreaterThan(d2)
}

// 加法
func DecimalAdd(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Add(d2)
}

// 连续加法
func DecimalAddMore(ds ...decimal.Decimal) decimal.Decimal {
	var all decimal.Decimal
	for i := 0; i < len(ds); i++ {
		if i == 0 {
			all = ds[i]
		} else {
			all = DecimalAdd(all, ds[i])
		}
	}
	return all
}

// 减法
func DecimalSub(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Sub(d2)
}

// 连续减法
func DecimalSubMore(ds ...decimal.Decimal) decimal.Decimal {
	var diff decimal.Decimal
	for i := 0; i < len(ds); i++ {
		if i == 0 {
			diff = ds[0]
		} else {
			diff = DecimalSub(diff, ds[i])
		}
	}
	return diff
}

// 乘法
func DecimalMul(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Mul(d2)
}

// 连续乘法
func DecimalMulMore(ds ...decimal.Decimal) decimal.Decimal {
	var result decimal.Decimal
	for i := 0; i < len(ds); i++ {
		if i == 0 {
			result = ds[0]
		} else {
			result = DecimalMul(result, ds[i])
		}
	}
	return result
}

// 除法
func DecimalDiv(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Div(d2)
}

// 连续除法
func DecimalDivMore(ds ...decimal.Decimal) decimal.Decimal {
	var result decimal.Decimal
	for i := 0; i < len(ds); i++ {
		if i == 0 {
			result = ds[0]
		} else {
			result = DecimalDiv(result, ds[i])
		}
	}
	return result
}

// int
func DecimalToInt64(d decimal.Decimal) int64 {
	return d.IntPart()
}

func DecimalToString(d decimal.Decimal) string {
	return d.String()
}

// float
func DecimalToFloat64(d decimal.Decimal) float64 {
	f, exact := d.Float64()
	if !exact {
		return f
	}
	return f
}
