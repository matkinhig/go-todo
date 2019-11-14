package types

type HeaderVNPAY struct {
	LoaiBanGhi   string
	MaSoBank     string
	NgayGiaoDich string
}

type BodyVNPAY struct {
	LoaiBanGhi             string
	MTI                    string
	AccountNumber          string
	ProcessingCode         string
	Amount                 string
	SystemTrace            string
	TransactionCode        string
	TransactionTime        string
	TransactionDate        string
	PaymentDate            string
	DeviceType             string
	MaKhoiTaoGiaoDich      string
	SoGiaoDich             string
	ResponseCode           string
	ProcessingCompareCode  string
	MaSoThietBiChapNhanThe string
	ThongTinBonus          string
	MerchantName           string
	Date                   string
	BankCode               string
}

type FooterVNPAY struct {
	LoaiBanGhi     string
	SoDongGiaoDich string
	NguoiTao       string
	GioTaoFile     string
	NgayTaoFile    string
}
