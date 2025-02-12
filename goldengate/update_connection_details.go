// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateConnectionDetails The information to update a Connection.
type UpdateConnectionDetails interface {

	// An object's Display Name.
	GetDisplayName() *string

	// Metadata about this specific object.
	GetDescription() *string

	// A simple key-value pair that is applied without any predefined name, type, or scope. Exists
	// for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Tags defined for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Refers to the customer's vault OCID.
	// If provided, it references a vault where GoldenGate can manage secrets. Customers must add policies to permit GoldenGate
	// to manage secrets contained within this vault.
	GetVaultId() *string

	// Refers to the customer's master key OCID.
	// If provided, it references a key to manage secrets. Customers must add policies to permit GoldenGate to use this key.
	GetKeyId() *string

	// An array of Network Security Group OCIDs used to define network access for either Deployments or Connections.
	GetNsgIds() []string
}

type updateconnectiondetails struct {
	JsonData       []byte
	DisplayName    *string                           `mandatory:"false" json:"displayName"`
	Description    *string                           `mandatory:"false" json:"description"`
	FreeformTags   map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags    map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	VaultId        *string                           `mandatory:"false" json:"vaultId"`
	KeyId          *string                           `mandatory:"false" json:"keyId"`
	NsgIds         []string                          `mandatory:"false" json:"nsgIds"`
	ConnectionType string                            `json:"connectionType"`
}

// UnmarshalJSON unmarshals json
func (m *updateconnectiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateconnectiondetails updateconnectiondetails
	s := struct {
		Model Unmarshalerupdateconnectiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.Description = s.Model.Description
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.VaultId = s.Model.VaultId
	m.KeyId = s.Model.KeyId
	m.NsgIds = s.Model.NsgIds
	m.ConnectionType = s.Model.ConnectionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateconnectiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ConnectionType {
	case "ORACLE":
		mm := UpdateOracleConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_OBJECT_STORAGE":
		mm := UpdateOciObjectStorageConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MONGODB":
		mm := UpdateMongoDbConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AZURE_DATA_LAKE_STORAGE":
		mm := UpdateAzureDataLakeStorageConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JAVA_MESSAGE_SERVICE":
		mm := UpdateJavaMessageServiceConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GOLDENGATE":
		mm := UpdateGoldenGateConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "POSTGRESQL":
		mm := UpdatePostgresqlConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MICROSOFT_SQLSERVER":
		mm := UpdateMicrosoftSqlserverConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ORACLE_NOSQL":
		mm := UpdateOracleNosqlConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KAFKA_SCHEMA_REGISTRY":
		mm := UpdateKafkaSchemaRegistryConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AMAZON_S3":
		mm := UpdateAmazonS3ConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SNOWFLAKE":
		mm := UpdateSnowflakeConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HDFS":
		mm := UpdateHdfsConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MYSQL":
		mm := UpdateMysqlConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KAFKA":
		mm := UpdateKafkaConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AZURE_SYNAPSE_ANALYTICS":
		mm := UpdateAzureSynapseConnectionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDisplayName returns DisplayName
func (m updateconnectiondetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDescription returns Description
func (m updateconnectiondetails) GetDescription() *string {
	return m.Description
}

//GetFreeformTags returns FreeformTags
func (m updateconnectiondetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m updateconnectiondetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetVaultId returns VaultId
func (m updateconnectiondetails) GetVaultId() *string {
	return m.VaultId
}

//GetKeyId returns KeyId
func (m updateconnectiondetails) GetKeyId() *string {
	return m.KeyId
}

//GetNsgIds returns NsgIds
func (m updateconnectiondetails) GetNsgIds() []string {
	return m.NsgIds
}

func (m updateconnectiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateconnectiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
