package user

type Gender int
type Status string

const (
	Male   Gender = 0
	Female Gender = 1
)

const (
	Active   Status = "active"
	Inactive Status = "inactive"
	Deleted  Status = "deleted"
)
