package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strconv"
)

type BaseController struct {
	*revel.Controller
}

type CtrlErr map[string]interface{}

func (r BaseController) GetParam(key, defaultValue string) string {
	value := r.Params.Query.Get(key)
	if len(value) == 0 {
		value = defaultValue
	}
	log.Printf("GetParam: key=%s, v=%s\n", key, value)
	return value
}

func parseUintOrDefault(intStr string, _default uint64) uint64 {
	if value, err := strconv.ParseUint(intStr, 0, 64); err != nil {
		return _default
	} else {
		return value
	}
}

func parseIntOrDefault(intStr string, _default int64) int64 {
	if value, err := strconv.ParseInt(intStr, 0, 64); err != nil {
		return _default
	} else {
		return value
	}
}

func convertToObjectIdHex(id string) (result bson.ObjectId, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Unable to convert %v to object id", id)
		}
	}()

	return bson.ObjectIdHex(id), err
}

func buildErrResponse(err error, errorCode string) CtrlErr {
	ctrlErr := CtrlErr{}
	ctrlErr["error_message"] = err.Error()
	ctrlErr["error_code"] = errorCode
	return ctrlErr
}
