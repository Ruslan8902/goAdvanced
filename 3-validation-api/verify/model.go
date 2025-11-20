package verify

import "encoding/json"

type Db interface {
	WriteStorage([]byte)
	ReadStorage() ([]byte, error)
}

type EmailHash struct {
	Email string `json:"email" validate:"email"`
	Hash  string `json:"hash"`
}

type EmailHashList struct {
	EmailHashs []EmailHash
}

type EmailHashListWithDb struct {
	EmailHashList
	Db Db
}

func NewBinListWithDb(db Db) *EmailHashListWithDb {
	binArrBytes, _ := db.ReadStorage()

	var binArr []EmailHash
	err := json.Unmarshal(binArrBytes, &binArr)
	if err != nil {
		return &EmailHashListWithDb{
			EmailHashList: EmailHashList{
				EmailHashs: []EmailHash{},
			},
			Db: db,
		}
	}

	return &EmailHashListWithDb{
		EmailHashList: EmailHashList{
			EmailHashs: binArr,
		},
		Db: db,
	}
}
