package ee_GH

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type ee_GH struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyNegativePrefix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'ee_GH' locale
func New() locales.Translator {
	return &ee_GH{
		locale:                 "ee_GH",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UZS", "VEB", "VEF", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "dzv", "dzd", "ted", "af??", "dam", "mas", "sia", "dea", "any", "kel", "ade", "dzm"},
		monthsNarrow:           []string{"", "d", "d", "t", "a", "d", "m", "s", "d", "a", "k", "a", "d"},
		monthsWide:             []string{"", "dzove", "dzodze", "tedoxe", "af??f??e", "dama", "masa", "siaml??m", "deasiamime", "any??ny??", "kele", "ade??mekp??xe", "dzome"},
		daysAbbreviated:        []string{"k??s", "dzo", "bla", "ku??", "yaw", "fi??", "mem"},
		daysNarrow:             []string{"k", "d", "b", "k", "y", "f", "m"},
		daysShort:              []string{"k??s", "dzo", "bla", "ku??", "yaw", "fi??", "mem"},
		daysWide:               []string{"k??si??a", "dzo??a", "bla??a", "ku??a", "yawo??a", "fi??a", "memle??a"},
		periodsAbbreviated:     []string{"??di", "??etr??"},
		periodsNarrow:          []string{"??", "??"},
		periodsWide:            []string{"??di", "??etr??"},
		erasAbbreviated:        []string{"hY", "Y??"},
		erasNarrow:             []string{"hY", "Y??"},
		erasWide:               []string{"Hafi Yesu Va Do ??g??", "Yesu ????li"},
		timezones:              map[string]string{"MST": "Makau ga??o??o??oanyime", "UYT": "Uruguai ga??o??o??oanyime", "NZDT": "NZDT", "MESZ": "Titina Europe ??kekeme ga??o??ome", "HNOG": "??eto??o??e Grinlan?? ga??o??o??oanyime", "SGT": "SGT", "PDT": "Pacific ??kekme ga??o??ome", "HECU": "Kuba ??kekeme ga??o??ome", "COT": "Kolombia ga??o??o??oanyime", "BT": "BT", "CST": "Titina America ga??o??o??oanyime", "CDT": "Titina America ??kekeme ga??o??ome", "MEZ": "Titina Europe ga??o??o??oanyime", "WARST": "??eto??o??e Argentina dzome????li ga??o??ome", "HNT": "Niufaun??lan?? ga??o??o??oanyime", "ACDT": "Titina Australia ??kekeme ga??o??ome", "HNCU": "Kuba ga??o??o??oanyime", "WART": "??eto??o??e Argentina ga??o??o??oanyime", "OESZ": "??edze??e Europe ??kekeme ga??o??ome", "AEDT": "??edze??e Australia ??kekeme ga??o??ome", "ART": "Argentina ga??o??o??oanyime", "SAST": "Anyiehe Africa ga??o??ome", "EAT": "??edze??e Africa ga??o??ome", "AKDT": "Alaska ??kekeme ga??o??ome", "NZST": "NZST", "HAT": "Niufaun??lan?? ??kekeme ga??o??ome", "AEST": "??edze??e Australia ga??o??o??oanyime", "AST": "Atlantic ga??o??o??oanyime", "ADT": "Atlantic ??kekeme ga??o??ome", "HKST": "H??ng K??ng dzome????li ga??o??ome", "CLT": "Tsile ga??o??o??oanyime", "COST": "Kolombia dzome????li ga??o??ome", "UYST": "Uruguai dzome????li ga??o??ome", "OEZ": "??edze??e Europe ga??o??o??oanyime", "CAT": "Titina Afrika ga??o??ome", "CHADT": "CHADT", "AWST": "??eto??o??e Australia ga??o??o??oanyime", "ACWDT": "Australia ??eto??ofe ??kekeme ga??o??ome", "HEOG": "??eto??o??e Grinlan?? dzome????li ga??o??ome", "HKT": "H??ng K??ng ga??o??o??oanyi me", "ECT": "Ikued?? dzome????li ga??o??ome", "WESZ": "??eto??o??e Europe ??kekeme ga??o??ome", "GFT": "Frentsi Guiana ga??o??ome", "HENOMX": "HENOMX", "HNPMX": "HNPMX", "HEPM": "Saint Pierre kple Mikuelon ??kekeme ga??o??ome", "IST": "IST", "ACST": "Titina Australia ga??o??o??oanyime", "SRT": "Suriname ga??o??ome", "GMT": "Greenwich ga??o??ome", "LHST": "LHST", "TMT": "T??kmenistan ga??o??o??oanyime", "JDT": "Japan ??kekeme ga??o??ome", "HNPM": "Saint Pierre kple Mikuelon ga??o??o??oanyime", "MYT": "MYT", "BOT": "Bolivia ga??o??ome", "HNEG": "??edze??e Grinlan?? ga??o??o??oanyime", "MDT": "Makau ??kekeme ga??o??ome", "HAST": "Hawaii-Aleutia ga??o??o??oanyime", "TMST": "T??kmenistan dzome????li ga??o??ome", "LHDT": "LHDT", "WAST": "??eto??o??e Africa ??kekeme ga??o??ome", "EST": "??edze??e America ga??o??o??oanyime", "WEZ": "??eto??o??e Europe ga??o??o??oanyime", "HEPMX": "HEPMX", "EDT": "??edze??e America ??kekeme ga??o??ome", "WIB": "WIB", "GYT": "Gayana ga??o??ome", "ACWST": "Australia ??eto??ofe ga??o??o??oanyime", "HADT": "Hawaii-Aleutia ??kekeme ga??o??ome", "JST": "Japan ga??o??o??anyime", "CLST": "Tsile dzome????li ga??o??ome", "ChST": "ChST", "AWDT": "??eto??o??e Australia ??kekeme ga??o??ome", "PST": "Pacific ga??o??o??oanyime", "WIT": "WIT", "HNNOMX": "HNNOMX", "ARST": "Argentina dzome????li ga??o??ome", "WAT": "??eto??o??e Afrika ga??o??o??oanyime", "HEEG": "??edze??e Grinlan?? dzome????li ga??o??ome", "?????????": "Amazon dzome????li ga??o??ome", "AKST": "Alaska ga??o??o??oanyime", "WITA": "WITA", "CHAST": "CHAST", "VET": "Venezuela ga??o??ome"},
	}
}

// Locale returns the current translators string locale
func (ee *ee_GH) Locale() string {
	return ee.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ee_GH'
func (ee *ee_GH) PluralsCardinal() []locales.PluralRule {
	return ee.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ee_GH'
func (ee *ee_GH) PluralsOrdinal() []locales.PluralRule {
	return ee.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ee_GH'
func (ee *ee_GH) PluralsRange() []locales.PluralRule {
	return ee.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ee_GH'
func (ee *ee_GH) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ee_GH'
func (ee *ee_GH) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ee_GH'
func (ee *ee_GH) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ee *ee_GH) MonthAbbreviated(month time.Month) string {
	return ee.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ee *ee_GH) MonthsAbbreviated() []string {
	return ee.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ee *ee_GH) MonthNarrow(month time.Month) string {
	return ee.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ee *ee_GH) MonthsNarrow() []string {
	return ee.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ee *ee_GH) MonthWide(month time.Month) string {
	return ee.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ee *ee_GH) MonthsWide() []string {
	return ee.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ee *ee_GH) WeekdayAbbreviated(weekday time.Weekday) string {
	return ee.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ee *ee_GH) WeekdaysAbbreviated() []string {
	return ee.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ee *ee_GH) WeekdayNarrow(weekday time.Weekday) string {
	return ee.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ee *ee_GH) WeekdaysNarrow() []string {
	return ee.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ee *ee_GH) WeekdayShort(weekday time.Weekday) string {
	return ee.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ee *ee_GH) WeekdaysShort() []string {
	return ee.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ee *ee_GH) WeekdayWide(weekday time.Weekday) string {
	return ee.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ee *ee_GH) WeekdaysWide() []string {
	return ee.daysWide
}

// Decimal returns the decimal point of number
func (ee *ee_GH) Decimal() string {
	return ee.decimal
}

// Group returns the group of number
func (ee *ee_GH) Group() string {
	return ee.group
}

// Group returns the minus sign of number
func (ee *ee_GH) Minus() string {
	return ee.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ee_GH' and handles both Whole and Real numbers based on 'v'
func (ee *ee_GH) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ee_GH' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ee *ee_GH) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ee_GH'
func (ee *ee_GH) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ee.currencies[currency]
	l := len(s) + len(symbol) + 0
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ee.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ee.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, ee.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ee.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ee_GH'
// in accounting notation.
func (ee *ee_GH) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ee.currencies[currency]
	l := len(s) + len(symbol) + 2
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ee.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ee.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, ee.currencyNegativePrefix[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ee.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ee.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ee_GH'
func (ee *ee_GH) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ee_GH'
func (ee *ee_GH) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ee.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x6c, 0x69, 0x61}...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ee_GH'
func (ee *ee_GH) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ee.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x6c, 0x69, 0x61}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ee_GH'
func (ee *ee_GH) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ee.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, ee.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x6c, 0x69, 0x61}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ee_GH'
func (ee *ee_GH) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ee.periodsAbbreviated[0]...)
	} else {
		b = append(b, ee.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20, 0x67, 0x61}...)
	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ee.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ee_GH'
func (ee *ee_GH) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ee.periodsAbbreviated[0]...)
	} else {
		b = append(b, ee.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20, 0x67, 0x61}...)
	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ee.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ee.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ee_GH'
func (ee *ee_GH) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ee.periodsAbbreviated[0]...)
	} else {
		b = append(b, ee.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20, 0x67, 0x61}...)
	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ee.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ee.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ee_GH'
func (ee *ee_GH) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, ee.periodsAbbreviated[0]...)
	} else {
		b = append(b, ee.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20, 0x67, 0x61}...)
	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ee.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ee.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ee.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
