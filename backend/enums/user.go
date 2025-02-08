package enums

type Gender int
type Status string
type Role int

const (
	Male   Gender = 0
	Female Gender = 1
)

const (
	Active   Status = "active"
	Inactive Status = "inactive"
	Deleted  Status = "deleted"
)

const (
	SystemAdmin    Role = 2
	DuTruong       Role = 4
	HuynhTruong    Role = 8
	PhanDoanTruong Role = 16
	NganhTruong    Role = 32
	KyLuat         Role = 64
	SinhHoat       Role = 128
	PhungVu        Role = 256
	HocTap         Role = 512
	BanDieuHanh    Role = 1024
)
