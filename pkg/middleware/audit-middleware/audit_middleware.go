package audit_middleware

import (
	"encoding/json"
	"fmt"
	"github.com/oklog/ulid/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"reflect"
	"wabustock/global/global_entities"
)

func RegisterCallbacks(db *gorm.DB) error {
	db.Callback().Create().After("gorm:create").Register("custom_plugin:create_audit_log", createAuditLog)
	db.Callback().Update().After("gorm:update").Register("custom_plugin:update_audit_log", updateAuditLog)
	db.Callback().Delete().Before("gorm:delete").Register("custom_plugin:delete_audit_log", deleteAuditLog)
	return nil
}

func createAuditLog(db *gorm.DB) {
	if db.Statement.Schema != nil && db.Statement.Schema.Table == "audit_logs" || db.Error != nil ||
		db.Statement.Schema.Table == "user_role" {
		return
	}

	recordMap, err := getDataBeforeOperation(db)
	if err != nil {
		return
	}
	objId := getKeyFromData("id", recordMap)
	auditLog := &global_entities.AuditLog{
		Id:            ulid.Make().String(),
		TableName:     db.Statement.Schema.Table,
		OperationType: "CREATE",
		ObjectId:      objId,
		Data:          prepareData(recordMap),
		UserId:        getCurrentUser(),
	}
	if err := db.Session(&gorm.Session{SkipHooks: true, NewDB: true}).Table("audit_logs").Create(auditLog).Error; err != nil {
		logrus.Error(fmt.Errorf("error in audit log creation: %s", err.Error()))
		return
	}
}

func updateAuditLog(db *gorm.DB) {
	if db.Statement.Schema != nil && db.Statement.Schema.Table == "audit_logs" || db.Error != nil {
		return
	}

	recordMap, err := getDataBeforeOperation(db)
	if err != nil {
		return
	}
	objId := getKeyFromData("id", recordMap)
	auditLog := &global_entities.AuditLog{
		Id:            ulid.Make().String(),
		TableName:     db.Statement.Schema.Table,
		OperationType: "UPDATE",
		ObjectId:      objId,
		Data:          prepareData(recordMap),
		UserId:        getCurrentUser(),
	}
	if err := db.Session(&gorm.Session{SkipHooks: true, NewDB: true}).Table("audit_logs").Create(auditLog).Error; err != nil {
		logrus.Error(fmt.Errorf("error in audit log creation: %s", err.Error()))
		return
	}
}

func deleteAuditLog(db *gorm.DB) {
	if db.Statement.Schema != nil && db.Statement.Schema.Table == "audit_logs" || db.Error != nil {
		return
	}

	recordMap, err := getDataBeforeOperation(db)
	if err != nil {
		return
	}
	objId := getKeyFromData("id", recordMap)
	auditLog := &global_entities.AuditLog{
		Id:            ulid.Make().String(),
		TableName:     db.Statement.Schema.Table,
		OperationType: "DELETE",
		ObjectId:      objId,
		Data:          prepareData(recordMap),
		UserId:        getCurrentUser(),
	}
	if err := db.Session(&gorm.Session{SkipHooks: true, NewDB: true}).Table("audit_logs").Create(auditLog).Error; err != nil {
		logrus.Error(fmt.Errorf("error in audit log creation: %s", err.Error()))
		return
	}
}

func getDataBeforeOperation(db *gorm.DB) (map[string]interface{}, error) {
	objMap := map[string]interface{}{}
	if db.Error == nil && !db.DryRun {
		objectType := reflect.TypeOf(db.Statement.ReflectValue.Interface())

		// Create a new instance of the object type
		targetObj := reflect.New(objectType).Interface()

		primaryKeyValue := ""
		value := db.Statement.ReflectValue

		// Check if the value is a struct
		if value.Kind() == reflect.Struct {
			primaryKeyValue = value.FieldByName("Id").String()
		}

		// Fetch the target object separately
		if err := db.Session(&gorm.Session{SkipHooks: true, NewDB: true}).Where("id = ?", primaryKeyValue).First(&targetObj).Error; err != nil {
			logrus.Error(fmt.Errorf("gorm callback: error while finding target object: %s", err.Error()))
			return nil, err
		}

		jsonBytes, err := json.Marshal(targetObj)
		if err != nil {
			logrus.Error(fmt.Errorf("gorm callback: error while marshalling json data: %s", err.Error()))
			return nil, err
		}
		json.Unmarshal(jsonBytes, &objMap)
	}
	return objMap, nil
}

func getKeyFromData(key string, data map[string]interface{}) string {
	objId, ok := data[key]
	if !ok {
		return ""
	}
	return objId.(string)
}

func prepareData(data map[string]interface{}) datatypes.JSON {
	dataByte, _ := json.Marshal(&data)
	return dataByte
}

func getCurrentUser() string {
	return "admin"
}
