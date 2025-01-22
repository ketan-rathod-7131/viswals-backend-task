package dto

type RawUserData struct {
	Id           int64
	FirstName    string
	LastName     string
	Email        string
	ParentUserId int64
	CreatedAt    int64
	DeletedAt    int64
	MergedAt     int64
}
