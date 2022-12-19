package loader

import "time"

type Model struct {
	CreateDate  time.Time `json:"create_date" hexya:"type=datetime;display_name=Created On;noCopy"`
	CreateUID   int64     `json:"create_uid" hexya:"display_name=Created By;noCopy"`
	WriteDate   time.Time `json:"write_date" hexya:"type=datetime;display_name=Updated On;noCopy"`
	WriteUID    int64     `json:"write_uid" hexya:"display_name=Updated By;noCopy"`
	LastUpdate  time.Time `json:"__last_update" hexya:"type=datetime;display_name=Updated On;noCopy"`
	DisplayName string    `json:"display_name" hexya:"type=compute;display_name=Display Name;noCopy"`
}
