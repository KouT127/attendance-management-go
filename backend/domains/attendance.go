package domains

type Attendance struct {
	Id     uint `json:"id"`
	UserId string `json:"userId"`
	Kind   uint8 `json:"kind"`
	Remark string `json:"remark"`
}
