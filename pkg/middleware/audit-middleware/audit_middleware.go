package audit_middleware

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"reflect"
	"time"
	"wabustock/global/marker"
	user_data "wabustock/pkg/utils/user-data"
)

func RegisterCallbacks(db *gorm.DB) error {
	db.Callback().Create().Before("gorm:create").Register("custom_plugin:create_audit_log", createAuditLog)
	db.Callback().Update().Before("gorm:update").Register("custom_plugin:update_audit_log", updateAuditLog)
	db.Callback().Delete().Before("gorm:delete").Register("custom_plugin:delete_audit_log", deleteAuditLog)
	return nil
}

func createAuditLog(db *gorm.DB) {

	//if auditable, ok := db.Statement.Model.(marker.Auditable); ok && auditable.HasAuditModel() {
	//	// Skip logic for models implementing Auditable
	//	return
	//}

	if model, ok := db.Statement.Model.(marker.Auditable); ok {
		if !model.HasAuditModel() {
			return
		}
	} else {
		return
	}

	if db.Statement.Schema != nil && db.Statement.Schema.Table == "audit_logs" || db.Error != nil ||
		db.Statement.Schema.Table == "user_role" {
		return
	}

	//model := db.Statement.Model.(generic_models.AuditModel)
	//
	//userId, getUserIdErr := user_data.GetUserId()
	//
	//if getUserIdErr != nil {
	//	model.CreatedBy = userId
	//}
	//
	//model.CreatedAt = time.Now()
	//model.UpdatedAt = time.Now()

	userId, getUserIdErr := user_data.GetUserId()
	var ifNil *string

	if userId == "" {
		ifNil = nil
	} else {
		ifNil = &userId
	}
	if getUserIdErr != nil {
		updateAuditModelFields(db.Statement.Model, ifNil, "PUT")
	}

	//recordMap, err := getDataBeforeOperation(db)
	//if err != nil {
	//	return
	//}
	//objId := getKeyFromData("id", recordMap)
	//auditLog := &global_entities.AuditLog{
	//	Id:            ulid.Make().String(),
	//	TableName:     db.Statement.Schema.Table,
	//	OperationType: "CREATE",
	//	ObjectId:      objId,
	//	Data:          prepareData(recordMap),
	//	CreatedBy:     user_data.GetUserId(),
	//}
	//if err := db.Session(&gorm.Session{SkipHooks: true, NewDB: true}).Table("audit_logs").Create(auditLog).Error; err != nil {
	//	logrus.Error(fmt.Errorf("error in audit log creation: %s", err.Error()))
	//	return
	//}
}

func updateAuditLog(db *gorm.DB) {
	if model, ok := db.Statement.Model.(marker.Auditable); ok {
		if !model.HasAuditModel() {
			return
		}
	} else {
		return
	}

	if db.Statement.Schema != nil && db.Statement.Schema.Table == "audit_logs" || db.Error != nil {
		return
	}

	//model := db.Statement.Model.(generic_models.AuditModel)

	//userId, getUserIdErr := user_data.GetUserId()
	//
	//if getUserIdErr != nil {
	//	model.UpdatedBy = &userId
	//}
	//
	//model.UpdatedAt = time.Now()

	userId, _ := user_data.GetUserId()
	var ifNil *string

	if userId == "" {
		ifNil = nil
	} else {
		ifNil = &userId
	}
	updateAuditModelFields(db.Statement.Model, ifNil, "PUT")

	//recordMap, err := getDataBeforeOperation(db)
	//if err != nil {
	//	return
	//}
	//
	//
	//objId := getKeyFromData("id", recordMap)
	//auditLog := &global_entities.AuditLog{
	//	Id:            ulid.Make().String(),
	//	TableName:     db.Statement.Schema.Table,
	//	OperationType: "UPDATE",
	//	ObjectId:      objId,
	//	Data:          prepareData(recordMap),
	//	UpdatedBy:     utils.Ptr(user_data.GetUserId()),
	//}
	//if err := db.Session(&gorm.Session{SkipHooks: true, NewDB: true}).Table("audit_logs").Create(auditLog).Error; err != nil {
	//	logrus.Error(fmt.Errorf("error in audit log creation: %s", err.Error()))
	//	return
	//}
}

func deleteAuditLog(db *gorm.DB) {
	if model, ok := db.Statement.Model.(marker.Auditable); ok {
		if !model.HasAuditModel() {
			return
		}
	} else {
		return
	}

	if db.Statement.Schema != nil && db.Statement.Schema.Table == "audit_logs" || db.Error != nil {
		return
	}

	//model := db.Statement.Model.(generic_models.AuditModel)
	//
	//userId, getUserIdErr := user_data.GetUserId()
	//
	//if getUserIdErr != nil {
	//	model.UpdatedBy = &userId
	//}

	userId, getUserIdErr := user_data.GetUserId()
	var ifNil *string

	if userId == "" {
		ifNil = nil
	} else {
		ifNil = &userId
	}
	if getUserIdErr != nil {
		updateAuditModelFields(db.Statement.Model, ifNil, "PUT")
	}

	//model.DeletedAt = time.Now()
	//model.UpdatedAt = time.Now()

	//recordMap, err := getDataBeforeOperation(db)
	//if err != nil {
	//	return
	//}
	//objId := getKeyFromData("id", recordMap)
	//auditLog := &global_entities.AuditLog{
	//	Id:            ulid.Make().String(),
	//	TableName:     db.Statement.Schema.Table,
	//	OperationType: "DELETE",
	//	ObjectId:      objId,
	//	Data:          prepareData(recordMap),
	//	UpdatedBy:     utils.Ptr(user_data.GetUserId()),
	//}
	//if err := db.Session(&gorm.Session{SkipHooks: true, NewDB: true}).Table("audit_logs").Create(auditLog).Error; err != nil {
	//	logrus.Error(fmt.Errorf("error in audit log creation: %s", err.Error()))
	//	return
	//}
}

func getDataBeforeOperation(db *gorm.DB) (map[string]interface{}, error) {
	objMap := map[string]interface{}{}
	if db.Error == nil && !db.DryRun {
		objectType := reflect.TypeOf(db.Statement.ReflectValue.Interface())

		// Create a new instance of the object type
		targetObj := reflect.New(objectType).Interface()

		primaryKeyValue := uuid.Nil
		value := db.Statement.ReflectValue

		// Check if the value is a struct
		if value.Kind() == reflect.Struct {
			primaryKeyValue, _ = uuid.Parse(value.FieldByName("Id").String())
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

func updateAuditModelFields(model interface{}, userId *string, saveType string) {
	value := reflect.ValueOf(model)

	// If model is a pointer, get the underlying value
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// Check if the value is a struct
	if value.Kind() != reflect.Struct {
		fmt.Println("Expected a struct, got:", value.Kind())
		return
	}

	if saveType == "POST" {
		// Set CreatedBy field
		createdByField := value.FieldByName("CreatedBy")
		if createdByField.IsValid() && createdByField.CanSet() {
			createdByField.Set(reflect.ValueOf(userId))

		}

	}

	if saveType == "POST" {

		// Set CreatedAt field
		createdAtField := value.FieldByName("CreatedAt")
		if createdAtField.IsValid() && createdAtField.CanSet() {
			createdAtField.Set(reflect.ValueOf(time.Now()))
		}
	}

	if saveType == "PUT" {

		// Set UpdatedAt field
		updatedAtField := value.FieldByName("UpdatedAt")
		if updatedAtField.IsValid() && updatedAtField.CanSet() {
			updatedAtField.Set(reflect.ValueOf(time.Now()))
		}

		// Set UpdatedBy field
		updatedByField := value.FieldByName("UpdatedBy")
		if updatedByField.IsValid() && updatedByField.CanSet() {

			updatedByField.Set(reflect.ValueOf(userId))
		}
	}

	if saveType == "DELETE" {

		//// Set UpdatedAt field
		//updatedAtField := value.FieldByName("DeletedAt")
		//if updatedAtField.IsValid() && updatedAtField.CanSet() {
		//	updatedAtField.Set(reflect.ValueOf(time.Now()))
		//}

		// Set UpdatedBy field
		updatedByField := value.FieldByName("UpdatedBy")
		if updatedByField.IsValid() && updatedByField.CanSet() {
			updatedByField.Set(reflect.ValueOf(userId))
		}
	}
}
