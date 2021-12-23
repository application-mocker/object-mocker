package model

import (
	"object-mocker/utils"
)

type Data struct {
	Id       string `json:"id"`
	CreateAt int64  `json:"create_at"` // the nano-second of create time
	UpdateAt int64  `json:"update_at"` // the nano-second of last update time
	DeleteAt int64  `json:"delete_at"` // the nano-second of delete time, if delete_at less than zero, the data not delete

	DataValue map[string]interface{} `json:"data"` // the custom datas
}

func (d *Data) Copy() *Data {
	data := &Data{}
	data.Id = d.Id
	data.DeleteAt = d.DeleteAt
	data.CreateAt = d.CreateAt
	data.UpdateAt = d.UpdateAt
	data.DataValue = d.DataValue

	return data
}

// NewData return a new data without Data.DataValue. And the Data.DataValue is empty but not nil.
func NewData() (*Data, error) {
	d := &Data{}
	var err error

	if d.Id, err = utils.NewUUIDString(); err != nil {
		return nil, err
	}

	createTime := utils.NowTime()
	d.CreateAt = createTime
	d.UpdateAt = createTime
	d.DeleteAt = -1
	d.DataValue = map[string]interface{}{}
	return d, nil
}

// NewDataWithDataValue return a new data with special value, if value is nil, the Data.DataValue of returns is nil.
func NewDataWithDataValue(value map[string]interface{}) (*Data, error) {
	d, err := NewData()
	if err != nil {
		return nil, err
	}

	d.DataValue = value
	return d, nil
}

// IsDelete return true when d is deleted. If the d.DeleteAt less than zero, is not be deleted.
func (d *Data) IsDelete() bool {
	return d.DeleteAt >= 0
}

// Delete set delete_at to now time.
func (d *Data) Delete() {
	d.DeleteAt = utils.NowTime()
}

// ToJson return the string of json object.
func (d *Data) ToJson() (string, error) {
	return utils.ToJson(d)
}

// String return the string of json object. Like Data.ToJson. But if marshal error, return "".
func (d *Data) String() string {
	if s, err := d.ToJson(); err != nil {
		return ""
	} else {
		return s
	}
}

// UpdateValue will update the Data.DataValue, and set Data.UpdateAt to now time by utils.NowTime.
func (d *Data) UpdateValue(value map[string]interface{}) {
	d.UpdateAt = utils.NowTime()
	d.DataValue = value
}
