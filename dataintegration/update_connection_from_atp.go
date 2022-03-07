// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v61/common"
	"strings"
)

// UpdateConnectionFromAtp The details to update an Autonomous Transaction Processing data asset connection.
type UpdateConnectionFromAtp struct {

	// Generated key that can be used in API calls to identify connection. On scenarios where reference to the connection is needed, a value can be passed in create.
	Key *string `mandatory:"true" json:"key"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"true" json:"objectVersion"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// User-defined description for the connection.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// The properties for the connection.
	ConnectionProperties []ConnectionProperty `mandatory:"false" json:"connectionProperties"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	// The Autonomous Transaction Processing instance service name.
	TnsAlias *string `mandatory:"false" json:"tnsAlias"`

	// Array of service names that are available for selection in the tnsAlias property.
	TnsNames []string `mandatory:"false" json:"tnsNames"`

	// The user name for the connection.
	Username *string `mandatory:"false" json:"username"`

	// The password for the connection.
	Password *string `mandatory:"false" json:"password"`

	PasswordSecret *SensitiveAttribute `mandatory:"false" json:"passwordSecret"`
}

//GetKey returns Key
func (m UpdateConnectionFromAtp) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m UpdateConnectionFromAtp) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m UpdateConnectionFromAtp) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m UpdateConnectionFromAtp) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m UpdateConnectionFromAtp) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m UpdateConnectionFromAtp) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetObjectVersion returns ObjectVersion
func (m UpdateConnectionFromAtp) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetIdentifier returns Identifier
func (m UpdateConnectionFromAtp) GetIdentifier() *string {
	return m.Identifier
}

//GetConnectionProperties returns ConnectionProperties
func (m UpdateConnectionFromAtp) GetConnectionProperties() []ConnectionProperty {
	return m.ConnectionProperties
}

//GetRegistryMetadata returns RegistryMetadata
func (m UpdateConnectionFromAtp) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m UpdateConnectionFromAtp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateConnectionFromAtp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateConnectionFromAtp) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateConnectionFromAtp UpdateConnectionFromAtp
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeUpdateConnectionFromAtp
	}{
		"ORACLE_ATP_CONNECTION",
		(MarshalTypeUpdateConnectionFromAtp)(m),
	}

	return json.Marshal(&s)
}
