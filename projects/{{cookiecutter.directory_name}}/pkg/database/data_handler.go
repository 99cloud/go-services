package database

import (
	"{{cookiecutter.project_slug}}/pkg/client"
	"{{cookiecutter.project_slug}}/pkg/logger"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

var OrmRegisterList = map[string]interface{}{}

type IDBQueryConditionOperator interface {
	ToQuery() string
	GetValue() []interface{}
	PositionCount() int
}

type IQueryConditionConnector interface {
	GetWhere(db *gorm.DB, isAnd bool) (*gorm.DB, error)
	GetDB(db *gorm.DB) (*gorm.DB, error)
}

type QueryConditionOr struct {
	// or contains one or more and
	Ors []IQueryConditionConnector
}

func (q QueryConditionOr) GetDB(db *gorm.DB) (*gorm.DB, error) {
	return q.GetWhere(db, false)
}

func (q QueryConditionOr) GetWhere(db *gorm.DB, ignore bool) (*gorm.DB, error) {
	var resultCondition *gorm.DB
	var err error
	for index := range q.Ors {
		resultCondition, err = q.Ors[index].GetWhere(db, index == 0)
		if err != nil {
			return db, err
		}
	}
	return resultCondition, err
}

type QueryConditionAnd struct {
	// and contains one or more operator like in,not in, between
	//ConditionOperators []IDBQueryConditionOperator
	ConditionOperators []IDBQueryConditionOperator
}

func (q QueryConditionAnd) GetDB(db *gorm.DB) (*gorm.DB, error) {
	return q.GetWhere(db, false)
}

func (q QueryConditionAnd) GetWhere(db *gorm.DB, isAnd bool) (*gorm.DB, error) {
	queryString := ""
	var values []interface{}
	for _, operator := range q.ConditionOperators {
		queryString += operator.ToQuery() + " AND "
		if operator.PositionCount() > len(operator.GetValue()) {
			return db, errors.New(fmt.Sprintf("Expected %d arguments but got %d", operator.PositionCount(), len(operator.GetValue())))
		}
		for index := 0; index < operator.PositionCount(); index++ {
			values = append(values, operator.GetValue()[index])
		}
	}
	// remove last " AND "
	queryString = queryString[:len(queryString)-4]
	if isAnd {
		return db.Where(queryString, values...), nil
	} else {
		return db.Or(queryString, values...), nil
	}
}

type DBQueryConditionInOperator struct {
	FieldName string
	Value     []interface{}
}

type DBQueryConditionNotInOperator struct {
	FieldName string
	Value     []interface{}
}

type DBQueryConditionLikeOperator struct {
	FieldName string
	Value     []interface{}
}

type DBQueryConditionNotLikeOperator struct {
	FieldName string
	Value     []interface{}
}

type DBQueryConditionEqualOperator struct {
	FieldName string
	Value     []interface{}
}

type DBQueryConditionNotEqualOperator struct {
	FieldName string
	Value     []interface{}
}

type DBQueryConditionBetweenOperator struct {
	FieldName string
	Value     []interface{}
}

type DBQueryConditionGtOrGtEOperator struct {
	FieldName string
	Equal     bool
	Value     []interface{}
}

type DBQueryConditionLtOrLtEOperator struct {
	FieldName string
	Equal     bool
	Value     []interface{}
}

func (operator DBQueryConditionInOperator) GetValue() []interface{} {
	return operator.Value
}

func (operator DBQueryConditionNotInOperator) GetValue() []interface{} {
	return operator.Value
}

func (operator DBQueryConditionNotLikeOperator) GetValue() []interface{} {
	return operator.Value
}

func (operator DBQueryConditionLikeOperator) GetValue() []interface{} {
	return operator.Value
}

func (operator DBQueryConditionEqualOperator) GetValue() []interface{} {
	return operator.Value
}

func (operator DBQueryConditionNotEqualOperator) GetValue() []interface{} {
	return operator.Value
}

func (operator DBQueryConditionBetweenOperator) GetValue() []interface{} {
	return operator.Value
}

func (operator DBQueryConditionGtOrGtEOperator) GetValue() []interface{} {
	return operator.Value
}

func (operator DBQueryConditionLtOrLtEOperator) GetValue() []interface{} {
	return operator.Value
}

func (operator DBQueryConditionInOperator) ToQuery() string {
	return fmt.Sprintf("%s IN (?)", operator.FieldName)
}

func (operator DBQueryConditionNotInOperator) ToQuery() string {
	return fmt.Sprintf("%s NOT IN (?)", operator.FieldName)
}

func (operator DBQueryConditionNotLikeOperator) ToQuery() string {
	return fmt.Sprintf("%s NOT LIKE (?)", operator.FieldName)
}

func (operator DBQueryConditionLikeOperator) ToQuery() string {
	return fmt.Sprintf("%s LIKE (?)", operator.FieldName)
}

func (operator DBQueryConditionEqualOperator) ToQuery() string {
	return fmt.Sprintf("%s = (?)", operator.FieldName)
}

func (operator DBQueryConditionNotEqualOperator) ToQuery() string {
	return fmt.Sprintf("%s <> (?)", operator.FieldName)
}

func (operator DBQueryConditionBetweenOperator) ToQuery() string {
	return fmt.Sprintf("%s BETWEEN ? AND ?", operator.FieldName)
}

func (operator DBQueryConditionGtOrGtEOperator) ToQuery() string {
	innerOperator := ">"
	if operator.Equal {
		innerOperator += "="
	}
	return fmt.Sprintf("%s %s ?", operator.FieldName, innerOperator)
}

func (operator DBQueryConditionLtOrLtEOperator) ToQuery() string {
	innerOperator := "<"
	if operator.Equal {
		innerOperator += "="
	}
	return fmt.Sprintf("%s %s ?", operator.FieldName, innerOperator)
}

func (operator DBQueryConditionInOperator) PositionCount() int {
	return 1
}

func (operator DBQueryConditionNotInOperator) PositionCount() int {
	return 1
}

func (operator DBQueryConditionNotLikeOperator) PositionCount() int {
	return 1
}

func (operator DBQueryConditionLikeOperator) PositionCount() int {
	return 1
}

func (operator DBQueryConditionEqualOperator) PositionCount() int {
	return 1
}

func (operator DBQueryConditionNotEqualOperator) PositionCount() int {
	return 1
}

func (operator DBQueryConditionBetweenOperator) PositionCount() int {
	return 2
}

func (operator DBQueryConditionLtOrLtEOperator) PositionCount() int {
	return 1
}

func (operator DBQueryConditionGtOrGtEOperator) PositionCount() int {
	return 1
}

func GetFieldDbColumnName(model interface{}, fieldName string) (string, error) {
	db, err := GetGorm()
	if err != nil {
		return "", err
	}
	modelScope := db.NewScope(model)
	fieldInfo, ok := modelScope.FieldByName(fieldName)
	if !ok {
		return "", errors.New(fmt.Sprintf("Failed to find field with given name %s", fieldName))
	}
	return fieldInfo.DBName, nil
}

func OrmRegister(db *gorm.DB) {
	for _, model := range OrmRegisterList {
		db.AutoMigrate(model)
	}
}

func GetGorm() (*gorm.DB, error) {
	dbClient, err := client.ClientSets().Database(OrmRegister)
	// dbClient.LogMode(true)
	if err != nil {
		logger.Error(nil, "Error")
		return nil, err
	}
	return dbClient, nil
}
