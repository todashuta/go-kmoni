package kmoni

import "errors"

// Data represents "kyoushin monitor"'s JSON format.
type Data struct {
	Result struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		IsAuth  bool   `json:"is_auth"`
	} `json:"result"`
	ReportTime    string `json:"report_time"`
	RegionCode    string `json:"region_code"`
	RequestTime   string `json:"request_time"`
	RegionName    string `json:"region_name"`
	Longitude     string `json:"longitude"`
	IsCancel      Bool   `json:"is_cancel"`
	Depth         string `json:"depth"`
	Calcintensity string `json:"calcintensity"`
	IsFinal       Bool   `json:"is_final"`
	IsTraining    Bool   `json:"is_training"`
	Latitude      string `json:"latitude"`
	OriginTime    string `json:"origin_time"`
	Security      struct {
		Realm string `json:"realm"`
		Hash  string `json:"hash"`
	} `json:"security"`
	Maginitude      string `json:"magunitude"`
	ReportNum       string `json:"report_time"`
	RequestHypoType string `json:"request_hypo_type"`
	ReportID        string `json:"report_id"`
	AlertFlg        string `json:"alertflg"`
}

type Bool bool

func (b *Bool) UnmarshalJSON(v []byte) error {
	s := string(v)
	switch s {
	case `true`:
		*b = true
	case `false`, `""`:
		*b = false
	default:
		return errors.New("unsupported value: " + s)
	}
	return nil
}
