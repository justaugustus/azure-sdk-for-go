package redis

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 0.12.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/azure-sdk-for-go/Godeps/_workspace/src/github.com/Azure/go-autorest/autorest"
	"github.com/Azure/azure-sdk-for-go/Godeps/_workspace/src/github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// KeyType enumerates the values for key type.
type KeyType string

const (
	// Primary specifies the primary state for key type.
	Primary KeyType = "Primary"
	// Secondary specifies the secondary state for key type.
	Secondary KeyType = "Secondary"
)

// SkuFamily enumerates the values for sku family.
type SkuFamily string

const (
	// C specifies the c state for sku family.
	C SkuFamily = "C"
	// P specifies the p state for sku family.
	P SkuFamily = "P"
)

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Basic specifies the basic state for sku name.
	Basic SkuName = "Basic"
	// Premium specifies the premium state for sku name.
	Premium SkuName = "Premium"
	// Standard specifies the standard state for sku name.
	Standard SkuName = "Standard"
)

// AccessKeys is redis cache access keys.
type AccessKeys struct {
	PrimaryKey   *string `json:"primaryKey,omitempty"`
	SecondaryKey *string `json:"secondaryKey,omitempty"`
}

// CreateOrUpdateParameters is parameters supplied to the CreateOrUpdate Redis
// operation.
type CreateOrUpdateParameters struct {
	ID         *string             `json:"id,omitempty"`
	Name       *string             `json:"name,omitempty"`
	Type       *string             `json:"type,omitempty"`
	Location   *string             `json:"location,omitempty"`
	Tags       *map[string]*string `json:"tags,omitempty"`
	Properties *Properties         `json:"properties,omitempty"`
}

// ListKeysResult is the response of redis list keys operation.
type ListKeysResult struct {
	autorest.Response `json:"-"`
	PrimaryKey        *string `json:"primaryKey,omitempty"`
	SecondaryKey      *string `json:"secondaryKey,omitempty"`
}

// ListResult is the response of list redis operation.
type ListResult struct {
	autorest.Response `json:"-"`
	Value             *[]ResourceType `json:"value,omitempty"`
	NextLink          *string         `json:"nextLink,omitempty"`
}

// ListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ListResult) ListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// Properties is parameters supplied to CreateOrUpdate redis operation.
type Properties struct {
	RedisVersion       *string             `json:"redisVersion,omitempty"`
	Sku                *Sku                `json:"sku,omitempty"`
	RedisConfiguration *map[string]*string `json:"redisConfiguration,omitempty"`
	EnableNonSslPort   *bool               `json:"enableNonSslPort,omitempty"`
	TenantSettings     *map[string]*string `json:"tenantSettings,omitempty"`
	ShardCount         *int                `json:"shardCount,omitempty"`
	VirtualNetwork     *string             `json:"virtualNetwork,omitempty"`
	Subnet             *string             `json:"subnet,omitempty"`
	StaticIP           *string             `json:"staticIP,omitempty"`
}

// ReadableProperties is parameters describing a redis instance
type ReadableProperties struct {
	ProvisioningState  *string             `json:"provisioningState,omitempty"`
	HostName           *string             `json:"hostName,omitempty"`
	Port               *int                `json:"port,omitempty"`
	SslPort            *int                `json:"sslPort,omitempty"`
	RedisVersion       *string             `json:"redisVersion,omitempty"`
	Sku                *Sku                `json:"sku,omitempty"`
	RedisConfiguration *map[string]*string `json:"redisConfiguration,omitempty"`
	EnableNonSslPort   *bool               `json:"enableNonSslPort,omitempty"`
	TenantSettings     *map[string]*string `json:"tenantSettings,omitempty"`
	ShardCount         *int                `json:"shardCount,omitempty"`
	VirtualNetwork     *string             `json:"virtualNetwork,omitempty"`
	Subnet             *string             `json:"subnet,omitempty"`
	StaticIP           *string             `json:"staticIP,omitempty"`
}

// ReadablePropertiesWithAccessKey is properties generated only in response to
// CreateOrUpdate redis operation.
type ReadablePropertiesWithAccessKey struct {
	AccessKeys         *AccessKeys         `json:"accessKeys,omitempty"`
	ProvisioningState  *string             `json:"provisioningState,omitempty"`
	HostName           *string             `json:"hostName,omitempty"`
	Port               *int                `json:"port,omitempty"`
	SslPort            *int                `json:"sslPort,omitempty"`
	RedisVersion       *string             `json:"redisVersion,omitempty"`
	Sku                *Sku                `json:"sku,omitempty"`
	RedisConfiguration *map[string]*string `json:"redisConfiguration,omitempty"`
	EnableNonSslPort   *bool               `json:"enableNonSslPort,omitempty"`
	TenantSettings     *map[string]*string `json:"tenantSettings,omitempty"`
	ShardCount         *int                `json:"shardCount,omitempty"`
	VirtualNetwork     *string             `json:"virtualNetwork,omitempty"`
	Subnet             *string             `json:"subnet,omitempty"`
	StaticIP           *string             `json:"staticIP,omitempty"`
}

// RegenerateKeyParameters is specifies which redis access keys to reset.
type RegenerateKeyParameters struct {
	KeyType KeyType `json:"keyType,omitempty"`
}

// Resource is
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// ResourceType is a single redis item in List or Get Operation.
type ResourceType struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	Properties        *ReadableProperties `json:"properties,omitempty"`
}

// ResourceWithAccessKey is a redis item in CreateOrUpdate Operation response.
type ResourceWithAccessKey struct {
	autorest.Response `json:"-"`
	ID                *string                          `json:"id,omitempty"`
	Name              *string                          `json:"name,omitempty"`
	Type              *string                          `json:"type,omitempty"`
	Location          *string                          `json:"location,omitempty"`
	Tags              *map[string]*string              `json:"tags,omitempty"`
	Properties        *ReadablePropertiesWithAccessKey `json:"properties,omitempty"`
}

// Sku is sku parameters supplied to the create redis operation.
type Sku struct {
	Name     SkuName   `json:"name,omitempty"`
	Family   SkuFamily `json:"family,omitempty"`
	Capacity *int      `json:"capacity,omitempty"`
}

// SubResource is
type SubResource struct {
	ID *string `json:"id,omitempty"`
}