// Code generated by ent, DO NOT EDIT.

package certificate

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/hm-edu/pki-service/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Certificate {
	return predicate.Certificate(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Certificate {
	return predicate.Certificate(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Certificate {
	return predicate.Certificate(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Certificate {
	return predicate.Certificate(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldUpdateTime, v))
}

// SslId applies equality check predicate on the "sslId" field. It's identical to SslIdEQ.
func SslId(v int) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldSslId, v))
}

// TransactionId applies equality check predicate on the "transactionId" field. It's identical to TransactionIdEQ.
func TransactionId(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldTransactionId, v))
}

// Serial applies equality check predicate on the "serial" field. It's identical to SerialEQ.
func Serial(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldSerial, v))
}

// CommonName applies equality check predicate on the "commonName" field. It's identical to CommonNameEQ.
func CommonName(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldCommonName, v))
}

// NotBefore applies equality check predicate on the "notBefore" field. It's identical to NotBeforeEQ.
func NotBefore(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldNotBefore, v))
}

// NotAfter applies equality check predicate on the "notAfter" field. It's identical to NotAfterEQ.
func NotAfter(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldNotAfter, v))
}

// IssuedBy applies equality check predicate on the "issuedBy" field. It's identical to IssuedByEQ.
func IssuedBy(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldIssuedBy, v))
}

// Source applies equality check predicate on the "source" field. It's identical to SourceEQ.
func Source(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldSource, v))
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldCreated, v))
}

// Ca applies equality check predicate on the "ca" field. It's identical to CaEQ.
func Ca(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldCa, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldLTE(FieldUpdateTime, v))
}

// SslIdEQ applies the EQ predicate on the "sslId" field.
func SslIdEQ(v int) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldSslId, v))
}

// SslIdNEQ applies the NEQ predicate on the "sslId" field.
func SslIdNEQ(v int) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldSslId, v))
}

// SslIdIn applies the In predicate on the "sslId" field.
func SslIdIn(vs ...int) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldSslId, vs...))
}

// SslIdNotIn applies the NotIn predicate on the "sslId" field.
func SslIdNotIn(vs ...int) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldSslId, vs...))
}

// SslIdGT applies the GT predicate on the "sslId" field.
func SslIdGT(v int) predicate.Certificate {
	return predicate.Certificate(sql.FieldGT(FieldSslId, v))
}

// SslIdGTE applies the GTE predicate on the "sslId" field.
func SslIdGTE(v int) predicate.Certificate {
	return predicate.Certificate(sql.FieldGTE(FieldSslId, v))
}

// SslIdLT applies the LT predicate on the "sslId" field.
func SslIdLT(v int) predicate.Certificate {
	return predicate.Certificate(sql.FieldLT(FieldSslId, v))
}

// SslIdLTE applies the LTE predicate on the "sslId" field.
func SslIdLTE(v int) predicate.Certificate {
	return predicate.Certificate(sql.FieldLTE(FieldSslId, v))
}

// SslIdIsNil applies the IsNil predicate on the "sslId" field.
func SslIdIsNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldIsNull(FieldSslId))
}

// SslIdNotNil applies the NotNil predicate on the "sslId" field.
func SslIdNotNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldNotNull(FieldSslId))
}

// TransactionIdEQ applies the EQ predicate on the "transactionId" field.
func TransactionIdEQ(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldTransactionId, v))
}

// TransactionIdNEQ applies the NEQ predicate on the "transactionId" field.
func TransactionIdNEQ(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldTransactionId, v))
}

// TransactionIdIn applies the In predicate on the "transactionId" field.
func TransactionIdIn(vs ...string) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldTransactionId, vs...))
}

// TransactionIdNotIn applies the NotIn predicate on the "transactionId" field.
func TransactionIdNotIn(vs ...string) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldTransactionId, vs...))
}

// TransactionIdGT applies the GT predicate on the "transactionId" field.
func TransactionIdGT(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldGT(FieldTransactionId, v))
}

// TransactionIdGTE applies the GTE predicate on the "transactionId" field.
func TransactionIdGTE(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldGTE(FieldTransactionId, v))
}

// TransactionIdLT applies the LT predicate on the "transactionId" field.
func TransactionIdLT(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldLT(FieldTransactionId, v))
}

// TransactionIdLTE applies the LTE predicate on the "transactionId" field.
func TransactionIdLTE(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldLTE(FieldTransactionId, v))
}

// TransactionIdContains applies the Contains predicate on the "transactionId" field.
func TransactionIdContains(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldContains(FieldTransactionId, v))
}

// TransactionIdHasPrefix applies the HasPrefix predicate on the "transactionId" field.
func TransactionIdHasPrefix(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldHasPrefix(FieldTransactionId, v))
}

// TransactionIdHasSuffix applies the HasSuffix predicate on the "transactionId" field.
func TransactionIdHasSuffix(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldHasSuffix(FieldTransactionId, v))
}

// TransactionIdIsNil applies the IsNil predicate on the "transactionId" field.
func TransactionIdIsNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldIsNull(FieldTransactionId))
}

// TransactionIdNotNil applies the NotNil predicate on the "transactionId" field.
func TransactionIdNotNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldNotNull(FieldTransactionId))
}

// TransactionIdEqualFold applies the EqualFold predicate on the "transactionId" field.
func TransactionIdEqualFold(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEqualFold(FieldTransactionId, v))
}

// TransactionIdContainsFold applies the ContainsFold predicate on the "transactionId" field.
func TransactionIdContainsFold(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldContainsFold(FieldTransactionId, v))
}

// SerialEQ applies the EQ predicate on the "serial" field.
func SerialEQ(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldSerial, v))
}

// SerialNEQ applies the NEQ predicate on the "serial" field.
func SerialNEQ(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldSerial, v))
}

// SerialIn applies the In predicate on the "serial" field.
func SerialIn(vs ...string) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldSerial, vs...))
}

// SerialNotIn applies the NotIn predicate on the "serial" field.
func SerialNotIn(vs ...string) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldSerial, vs...))
}

// SerialGT applies the GT predicate on the "serial" field.
func SerialGT(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldGT(FieldSerial, v))
}

// SerialGTE applies the GTE predicate on the "serial" field.
func SerialGTE(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldGTE(FieldSerial, v))
}

// SerialLT applies the LT predicate on the "serial" field.
func SerialLT(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldLT(FieldSerial, v))
}

// SerialLTE applies the LTE predicate on the "serial" field.
func SerialLTE(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldLTE(FieldSerial, v))
}

// SerialContains applies the Contains predicate on the "serial" field.
func SerialContains(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldContains(FieldSerial, v))
}

// SerialHasPrefix applies the HasPrefix predicate on the "serial" field.
func SerialHasPrefix(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldHasPrefix(FieldSerial, v))
}

// SerialHasSuffix applies the HasSuffix predicate on the "serial" field.
func SerialHasSuffix(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldHasSuffix(FieldSerial, v))
}

// SerialIsNil applies the IsNil predicate on the "serial" field.
func SerialIsNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldIsNull(FieldSerial))
}

// SerialNotNil applies the NotNil predicate on the "serial" field.
func SerialNotNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldNotNull(FieldSerial))
}

// SerialEqualFold applies the EqualFold predicate on the "serial" field.
func SerialEqualFold(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEqualFold(FieldSerial, v))
}

// SerialContainsFold applies the ContainsFold predicate on the "serial" field.
func SerialContainsFold(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldContainsFold(FieldSerial, v))
}

// CommonNameEQ applies the EQ predicate on the "commonName" field.
func CommonNameEQ(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldCommonName, v))
}

// CommonNameNEQ applies the NEQ predicate on the "commonName" field.
func CommonNameNEQ(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldCommonName, v))
}

// CommonNameIn applies the In predicate on the "commonName" field.
func CommonNameIn(vs ...string) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldCommonName, vs...))
}

// CommonNameNotIn applies the NotIn predicate on the "commonName" field.
func CommonNameNotIn(vs ...string) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldCommonName, vs...))
}

// CommonNameGT applies the GT predicate on the "commonName" field.
func CommonNameGT(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldGT(FieldCommonName, v))
}

// CommonNameGTE applies the GTE predicate on the "commonName" field.
func CommonNameGTE(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldGTE(FieldCommonName, v))
}

// CommonNameLT applies the LT predicate on the "commonName" field.
func CommonNameLT(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldLT(FieldCommonName, v))
}

// CommonNameLTE applies the LTE predicate on the "commonName" field.
func CommonNameLTE(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldLTE(FieldCommonName, v))
}

// CommonNameContains applies the Contains predicate on the "commonName" field.
func CommonNameContains(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldContains(FieldCommonName, v))
}

// CommonNameHasPrefix applies the HasPrefix predicate on the "commonName" field.
func CommonNameHasPrefix(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldHasPrefix(FieldCommonName, v))
}

// CommonNameHasSuffix applies the HasSuffix predicate on the "commonName" field.
func CommonNameHasSuffix(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldHasSuffix(FieldCommonName, v))
}

// CommonNameEqualFold applies the EqualFold predicate on the "commonName" field.
func CommonNameEqualFold(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEqualFold(FieldCommonName, v))
}

// CommonNameContainsFold applies the ContainsFold predicate on the "commonName" field.
func CommonNameContainsFold(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldContainsFold(FieldCommonName, v))
}

// NotBeforeEQ applies the EQ predicate on the "notBefore" field.
func NotBeforeEQ(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldNotBefore, v))
}

// NotBeforeNEQ applies the NEQ predicate on the "notBefore" field.
func NotBeforeNEQ(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldNotBefore, v))
}

// NotBeforeIn applies the In predicate on the "notBefore" field.
func NotBeforeIn(vs ...time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldNotBefore, vs...))
}

// NotBeforeNotIn applies the NotIn predicate on the "notBefore" field.
func NotBeforeNotIn(vs ...time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldNotBefore, vs...))
}

// NotBeforeGT applies the GT predicate on the "notBefore" field.
func NotBeforeGT(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldGT(FieldNotBefore, v))
}

// NotBeforeGTE applies the GTE predicate on the "notBefore" field.
func NotBeforeGTE(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldGTE(FieldNotBefore, v))
}

// NotBeforeLT applies the LT predicate on the "notBefore" field.
func NotBeforeLT(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldLT(FieldNotBefore, v))
}

// NotBeforeLTE applies the LTE predicate on the "notBefore" field.
func NotBeforeLTE(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldLTE(FieldNotBefore, v))
}

// NotBeforeIsNil applies the IsNil predicate on the "notBefore" field.
func NotBeforeIsNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldIsNull(FieldNotBefore))
}

// NotBeforeNotNil applies the NotNil predicate on the "notBefore" field.
func NotBeforeNotNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldNotNull(FieldNotBefore))
}

// NotAfterEQ applies the EQ predicate on the "notAfter" field.
func NotAfterEQ(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldNotAfter, v))
}

// NotAfterNEQ applies the NEQ predicate on the "notAfter" field.
func NotAfterNEQ(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldNotAfter, v))
}

// NotAfterIn applies the In predicate on the "notAfter" field.
func NotAfterIn(vs ...time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldNotAfter, vs...))
}

// NotAfterNotIn applies the NotIn predicate on the "notAfter" field.
func NotAfterNotIn(vs ...time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldNotAfter, vs...))
}

// NotAfterGT applies the GT predicate on the "notAfter" field.
func NotAfterGT(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldGT(FieldNotAfter, v))
}

// NotAfterGTE applies the GTE predicate on the "notAfter" field.
func NotAfterGTE(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldGTE(FieldNotAfter, v))
}

// NotAfterLT applies the LT predicate on the "notAfter" field.
func NotAfterLT(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldLT(FieldNotAfter, v))
}

// NotAfterLTE applies the LTE predicate on the "notAfter" field.
func NotAfterLTE(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldLTE(FieldNotAfter, v))
}

// NotAfterIsNil applies the IsNil predicate on the "notAfter" field.
func NotAfterIsNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldIsNull(FieldNotAfter))
}

// NotAfterNotNil applies the NotNil predicate on the "notAfter" field.
func NotAfterNotNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldNotNull(FieldNotAfter))
}

// IssuedByEQ applies the EQ predicate on the "issuedBy" field.
func IssuedByEQ(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldIssuedBy, v))
}

// IssuedByNEQ applies the NEQ predicate on the "issuedBy" field.
func IssuedByNEQ(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldIssuedBy, v))
}

// IssuedByIn applies the In predicate on the "issuedBy" field.
func IssuedByIn(vs ...string) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldIssuedBy, vs...))
}

// IssuedByNotIn applies the NotIn predicate on the "issuedBy" field.
func IssuedByNotIn(vs ...string) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldIssuedBy, vs...))
}

// IssuedByGT applies the GT predicate on the "issuedBy" field.
func IssuedByGT(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldGT(FieldIssuedBy, v))
}

// IssuedByGTE applies the GTE predicate on the "issuedBy" field.
func IssuedByGTE(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldGTE(FieldIssuedBy, v))
}

// IssuedByLT applies the LT predicate on the "issuedBy" field.
func IssuedByLT(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldLT(FieldIssuedBy, v))
}

// IssuedByLTE applies the LTE predicate on the "issuedBy" field.
func IssuedByLTE(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldLTE(FieldIssuedBy, v))
}

// IssuedByContains applies the Contains predicate on the "issuedBy" field.
func IssuedByContains(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldContains(FieldIssuedBy, v))
}

// IssuedByHasPrefix applies the HasPrefix predicate on the "issuedBy" field.
func IssuedByHasPrefix(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldHasPrefix(FieldIssuedBy, v))
}

// IssuedByHasSuffix applies the HasSuffix predicate on the "issuedBy" field.
func IssuedByHasSuffix(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldHasSuffix(FieldIssuedBy, v))
}

// IssuedByIsNil applies the IsNil predicate on the "issuedBy" field.
func IssuedByIsNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldIsNull(FieldIssuedBy))
}

// IssuedByNotNil applies the NotNil predicate on the "issuedBy" field.
func IssuedByNotNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldNotNull(FieldIssuedBy))
}

// IssuedByEqualFold applies the EqualFold predicate on the "issuedBy" field.
func IssuedByEqualFold(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEqualFold(FieldIssuedBy, v))
}

// IssuedByContainsFold applies the ContainsFold predicate on the "issuedBy" field.
func IssuedByContainsFold(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldContainsFold(FieldIssuedBy, v))
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldSource, v))
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldSource, v))
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...string) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldSource, vs...))
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...string) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldSource, vs...))
}

// SourceGT applies the GT predicate on the "source" field.
func SourceGT(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldGT(FieldSource, v))
}

// SourceGTE applies the GTE predicate on the "source" field.
func SourceGTE(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldGTE(FieldSource, v))
}

// SourceLT applies the LT predicate on the "source" field.
func SourceLT(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldLT(FieldSource, v))
}

// SourceLTE applies the LTE predicate on the "source" field.
func SourceLTE(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldLTE(FieldSource, v))
}

// SourceContains applies the Contains predicate on the "source" field.
func SourceContains(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldContains(FieldSource, v))
}

// SourceHasPrefix applies the HasPrefix predicate on the "source" field.
func SourceHasPrefix(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldHasPrefix(FieldSource, v))
}

// SourceHasSuffix applies the HasSuffix predicate on the "source" field.
func SourceHasSuffix(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldHasSuffix(FieldSource, v))
}

// SourceIsNil applies the IsNil predicate on the "source" field.
func SourceIsNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldIsNull(FieldSource))
}

// SourceNotNil applies the NotNil predicate on the "source" field.
func SourceNotNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldNotNull(FieldSource))
}

// SourceEqualFold applies the EqualFold predicate on the "source" field.
func SourceEqualFold(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEqualFold(FieldSource, v))
}

// SourceContainsFold applies the ContainsFold predicate on the "source" field.
func SourceContainsFold(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldContainsFold(FieldSource, v))
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldCreated, v))
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldCreated, v))
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldCreated, vs...))
}

// CreatedNotIn applies the NotIn predicate on the "created" field.
func CreatedNotIn(vs ...time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldCreated, vs...))
}

// CreatedGT applies the GT predicate on the "created" field.
func CreatedGT(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldGT(FieldCreated, v))
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldGTE(FieldCreated, v))
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldLT(FieldCreated, v))
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v time.Time) predicate.Certificate {
	return predicate.Certificate(sql.FieldLTE(FieldCreated, v))
}

// CreatedIsNil applies the IsNil predicate on the "created" field.
func CreatedIsNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldIsNull(FieldCreated))
}

// CreatedNotNil applies the NotNil predicate on the "created" field.
func CreatedNotNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldNotNull(FieldCreated))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldStatus, vs...))
}

// CaEQ applies the EQ predicate on the "ca" field.
func CaEQ(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEQ(FieldCa, v))
}

// CaNEQ applies the NEQ predicate on the "ca" field.
func CaNEQ(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldNEQ(FieldCa, v))
}

// CaIn applies the In predicate on the "ca" field.
func CaIn(vs ...string) predicate.Certificate {
	return predicate.Certificate(sql.FieldIn(FieldCa, vs...))
}

// CaNotIn applies the NotIn predicate on the "ca" field.
func CaNotIn(vs ...string) predicate.Certificate {
	return predicate.Certificate(sql.FieldNotIn(FieldCa, vs...))
}

// CaGT applies the GT predicate on the "ca" field.
func CaGT(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldGT(FieldCa, v))
}

// CaGTE applies the GTE predicate on the "ca" field.
func CaGTE(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldGTE(FieldCa, v))
}

// CaLT applies the LT predicate on the "ca" field.
func CaLT(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldLT(FieldCa, v))
}

// CaLTE applies the LTE predicate on the "ca" field.
func CaLTE(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldLTE(FieldCa, v))
}

// CaContains applies the Contains predicate on the "ca" field.
func CaContains(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldContains(FieldCa, v))
}

// CaHasPrefix applies the HasPrefix predicate on the "ca" field.
func CaHasPrefix(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldHasPrefix(FieldCa, v))
}

// CaHasSuffix applies the HasSuffix predicate on the "ca" field.
func CaHasSuffix(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldHasSuffix(FieldCa, v))
}

// CaIsNil applies the IsNil predicate on the "ca" field.
func CaIsNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldIsNull(FieldCa))
}

// CaNotNil applies the NotNil predicate on the "ca" field.
func CaNotNil() predicate.Certificate {
	return predicate.Certificate(sql.FieldNotNull(FieldCa))
}

// CaEqualFold applies the EqualFold predicate on the "ca" field.
func CaEqualFold(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldEqualFold(FieldCa, v))
}

// CaContainsFold applies the ContainsFold predicate on the "ca" field.
func CaContainsFold(v string) predicate.Certificate {
	return predicate.Certificate(sql.FieldContainsFold(FieldCa, v))
}

// HasDomains applies the HasEdge predicate on the "domains" edge.
func HasDomains() predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, DomainsTable, DomainsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDomainsWith applies the HasEdge predicate on the "domains" edge with a given conditions (other predicates).
func HasDomainsWith(preds ...predicate.Domain) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		step := newDomainsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Certificate) predicate.Certificate {
	return predicate.Certificate(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Certificate) predicate.Certificate {
	return predicate.Certificate(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Certificate) predicate.Certificate {
	return predicate.Certificate(sql.NotPredicates(p))
}
