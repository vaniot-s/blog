package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

type Model struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt LocalTime `gorm:"created_at" json:"createdAt"`
	UpdatedAt LocalTime `gorm:"updated_at" json:"updatedAt"`
	DeletedAt LocalTime `gorm:"deleted_at" json:"DeletedAt"`

}

type LocalTime struct {
	time.Time
}
func (t LocalTime) MarshalJSON() ([]byte, error) {
	//格式化秒
	seconds := t.Unix()
	return []byte(strconv.FormatInt(seconds, 10)), nil
}
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}