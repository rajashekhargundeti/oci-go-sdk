// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// The Organizations API allows you to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and its resources.
//

package tenantmanagercontrolplane

import (
	"strings"
)

// OrganizationTenancyRoleEnum Enum with underlying type: string
type OrganizationTenancyRoleEnum string

// Set of constants representing the allowable values for OrganizationTenancyRoleEnum
const (
	OrganizationTenancyRoleParent OrganizationTenancyRoleEnum = "PARENT"
	OrganizationTenancyRoleChild  OrganizationTenancyRoleEnum = "CHILD"
	OrganizationTenancyRoleNone   OrganizationTenancyRoleEnum = "NONE"
)

var mappingOrganizationTenancyRoleEnum = map[string]OrganizationTenancyRoleEnum{
	"PARENT": OrganizationTenancyRoleParent,
	"CHILD":  OrganizationTenancyRoleChild,
	"NONE":   OrganizationTenancyRoleNone,
}

var mappingOrganizationTenancyRoleEnumLowerCase = map[string]OrganizationTenancyRoleEnum{
	"parent": OrganizationTenancyRoleParent,
	"child":  OrganizationTenancyRoleChild,
	"none":   OrganizationTenancyRoleNone,
}

// GetOrganizationTenancyRoleEnumValues Enumerates the set of values for OrganizationTenancyRoleEnum
func GetOrganizationTenancyRoleEnumValues() []OrganizationTenancyRoleEnum {
	values := make([]OrganizationTenancyRoleEnum, 0)
	for _, v := range mappingOrganizationTenancyRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetOrganizationTenancyRoleEnumStringValues Enumerates the set of values in String for OrganizationTenancyRoleEnum
func GetOrganizationTenancyRoleEnumStringValues() []string {
	return []string{
		"PARENT",
		"CHILD",
		"NONE",
	}
}

// GetMappingOrganizationTenancyRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOrganizationTenancyRoleEnum(val string) (OrganizationTenancyRoleEnum, bool) {
	enum, ok := mappingOrganizationTenancyRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
