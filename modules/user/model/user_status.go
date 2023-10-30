package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type UserStatus int

const (
	UserStatusDoing UserStatus = iota
	UserStatusDeleted
)

var allUserStatuses = [3]string{"Doing", "Deleted"}

func (item *UserStatus) String() string {
	return allUserStatuses[*item]
}

func parseStr2ItemStatus(s string) (UserStatus, error) {
	for i := range allUserStatuses {
		if allUserStatuses[i] == s {
			return UserStatus(i), nil
		}
	}
	return UserStatus(0), errors.New("invalid status string")
}

/* Scan đọc dữ liệu từ Mysql để đem lên UserStatus */

func (item *UserStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}

	v, err := parseStr2ItemStatus(string(bytes))

	if err != nil {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}

	*item = v

	return nil

}

/* Ngược lại với hàm Scan */

func (item *UserStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.String(), nil
}

/* Json Encoding -  từ data structure từ UserStatus sang json value */

func (item *UserStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

/* Json Decoding -  ngược lại với MarshalJSON */

func (item *UserStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")
	itemValue, err := parseStr2ItemStatus(str)
	if err != nil {
		return err
	}

	*item = itemValue

	return nil
}
