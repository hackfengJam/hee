// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Swagger™ is a project used to describe and document RESTful APIs.
//
// The Swagger specification defines a set of files required to describe such an API. These files can then be used by the Swagger-UI project to display the API and Swagger-Codegen to generate clients in various languages. Additional utilities can also take advantage of the resulting files, such as testing tools.
// Now in version 2.0, Swagger is more enabling than ever. And it's 100% open source software.

// Package swagger struct definition
package swaggergen

import (
	"errors"
	"fmt"
)

// Swagger list the resource
type Swagger struct {
	SwaggerVersion      string                `json:"swagger,omitempty" yaml:"swagger,omitempty"`
	Infos               Information           `json:"info" yaml:"info"`
	Host                string                `json:"host,omitempty" yaml:"host,omitempty"`
	BasePath            string                `json:"basePath,omitempty" yaml:"basePath,omitempty"`
	Schemes             []string              `json:"schemes,omitempty" yaml:"schemes,omitempty"`
	Consumes            []string              `json:"consumes,omitempty" yaml:"consumes,omitempty"`
	Produces            []string              `json:"produces,omitempty" yaml:"produces,omitempty"`
	Paths               map[string]*Item      `json:"paths" yaml:"paths"`
	Definitions         map[string]Schema     `json:"definitions,omitempty" yaml:"definitions,omitempty"`
	SecurityDefinitions map[string]Security   `json:"securityDefinitions,omitempty" yaml:"securityDefinitions,omitempty"`
	Security            []map[string][]string `json:"security,omitempty" yaml:"security,omitempty"`
	Tags                []Tag                 `json:"tags,omitempty" yaml:"tags,omitempty"`
	ExternalDocs        *ExternalDocs         `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

// Information Provides metadata about the API. The metadata can be used by the clients if needed.
type Information struct {
	Title          string `json:"title,omitempty" yaml:"title,omitempty"`
	Description    string `json:"description,omitempty" yaml:"description,omitempty"`
	Version        string `json:"version,omitempty" yaml:"version,omitempty"`
	TermsOfService string `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`

	Contact Contact  `json:"contact,omitempty" yaml:"contact,omitempty"`
	License *License `json:"license,omitempty" yaml:"license,omitempty"`
}

// Contact information for the exposed API.
type Contact struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	URL   string `json:"url,omitempty" yaml:"url,omitempty"`
	EMail string `json:"email,omitempty" yaml:"email,omitempty"`
}

// License information for the exposed API.
type License struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	URL  string `json:"url,omitempty" yaml:"url,omitempty"`
}

// Item Describes the operations available on a single path.
type Item struct {
	Ref     string     `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Get     *Operation `json:"get,omitempty" yaml:"get,omitempty"`
	Put     *Operation `json:"put,omitempty" yaml:"put,omitempty"`
	Post    *Operation `json:"post,omitempty" yaml:"post,omitempty"`
	Delete  *Operation `json:"delete,omitempty" yaml:"delete,omitempty"`
	Options *Operation `json:"options,omitempty" yaml:"options,omitempty"`
	Head    *Operation `json:"head,omitempty" yaml:"head,omitempty"`
	Patch   *Operation `json:"patch,omitempty" yaml:"patch,omitempty"`

	//NOTE read only after first assignment
	Stash *Operation `json:"-" yaml:"-"`
}

// output Stash as string
func (p *Item) StashString() string {
	if p == nil {
		return "Item:<nil>"
	}
	if p.Stash == nil {
		return "Item:<nil stash>"
	}
	return fmt.Sprintf("Item:%+v", *(p.Stash))
}

// not deep copy, without Stash
func (p *Item) Copy() (np *Item) {
	if p == nil {
		return
	}
	np = new(Item)
	*np = *p
	np.Stash = p.Stash
	return
}

// merge items, should use the returned m as merged one
func (p *Item) Merge(n *Item) (*Item, error) {
	if p == nil {
		return n.Copy(), nil
	}
	if n == nil {
		return p.Copy(), nil
	}

	m := new(Item)
	if p.Get != nil && n.Get != nil && p.Get != n.Get {
		return nil, errors.New("Conflict Get Operations")
	}
	m.Get = p.Get
	if n.Get != nil {
		m.Get = n.Get
	}
	if p.Post != nil && n.Post != nil && p.Post != n.Post {
		return nil, errors.New("Conflict Post Operations")
	}
	m.Post = p.Post
	if n.Post != nil {
		m.Post = n.Post
	}
	if p.Put != nil && n.Put != nil && p.Put != n.Put {
		return nil, errors.New("Conflict Put Operations")
	}
	m.Put = p.Put
	if n.Put != nil {
		m.Put = n.Put
	}
	if p.Delete != nil && n.Delete != nil && p.Delete != n.Delete {
		return nil, errors.New("Conflict Delete Operations")
	}
	m.Delete = p.Delete
	if n.Delete != nil {
		m.Delete = n.Delete
	}
	if p.Options != nil && n.Options != nil && p.Options != n.Options {
		return nil, errors.New("Conflict Options Operations")
	}
	m.Options = p.Options
	if n.Options != nil {
		m.Options = n.Options
	}
	if p.Head != nil && n.Head != nil && p.Head != n.Head {
		return nil, errors.New("Conflict Head Operations")
	}
	m.Head = p.Head
	if n.Head != nil {
		m.Head = n.Head
	}
	if p.Patch != nil && n.Patch != nil && p.Patch != n.Patch {
		return nil, errors.New("Conflict Patch Operations")
	}
	m.Patch = p.Patch
	if n.Patch != nil {
		m.Patch = n.Patch
	}
	return m, nil
}

// Operation Describes a single API operation on a path.
type Operation struct {
	Tags        []string              `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary     string                `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string                `json:"description,omitempty" yaml:"description,omitempty"`
	OperationID string                `json:"operationId,omitempty" yaml:"operationId,omitempty"`
	Consumes    []string              `json:"consumes,omitempty" yaml:"consumes,omitempty"`
	Produces    []string              `json:"produces,omitempty" yaml:"produces,omitempty"`
	Schemes     []string              `json:"schemes,omitempty" yaml:"schemes,omitempty"`
	Parameters  []Parameter           `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Responses   map[string]Response   `json:"responses,omitempty" yaml:"responses,omitempty"`
	Security    []map[string][]string `json:"security,omitempty" yaml:"security,omitempty"`
	Deprecated  bool                  `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Author      []string              `json:"-" yaml:"-"`
}

// not deep copy
func (p *Operation) Copy() (np *Operation) {
	if p == nil {
		return
	}
	np = new(Operation)
	*np = *p
	return
}

// Parameter Describes a single operation parameter.
type Parameter struct {
	In          string          `json:"in,omitempty" yaml:"in,omitempty"`
	Name        string          `json:"name,omitempty" yaml:"name,omitempty"`
	Description string          `json:"description,omitempty" yaml:"description,omitempty"`
	Required    bool            `json:"required,omitempty" yaml:"required,omitempty"`
	Schema      *Schema         `json:"schema,omitempty" yaml:"schema,omitempty"`
	Type        string          `json:"type,omitempty" yaml:"type,omitempty"`
	Format      string          `json:"format,omitempty" yaml:"format,omitempty"`
	Items       *ParameterItems `json:"items,omitempty" yaml:"items,omitempty"`
	Default     interface{}     `json:"default,omitempty" yaml:"default,omitempty"`
}

// ParameterItems A limited subset of JSON-Schema's items object. It is used by parameter definitions that are not located in "body".
// http://swagger.io/specification/#itemsObject
type ParameterItems struct {
	Type             string            `json:"type,omitempty" yaml:"type,omitempty"`
	Format           string            `json:"format,omitempty" yaml:"format,omitempty"`
	Items            []*ParameterItems `json:"items,omitempty" yaml:"items,omitempty"` //Required if type is "array". Describes the type of items in the array.
	CollectionFormat string            `json:"collectionFormat,omitempty" yaml:"collectionFormat,omitempty"`
	Default          string            `json:"default,omitempty" yaml:"default,omitempty"`
}

// Schema Object allows the definition of input and output data types.
type Schema struct {
	Ref         string               `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Title       string               `json:"title,omitempty" yaml:"title,omitempty"`
	Format      string               `json:"format,omitempty" yaml:"format,omitempty"`
	Description string               `json:"description,omitempty" yaml:"description,omitempty"`
	Required    []string             `json:"required,omitempty" yaml:"required,omitempty"`
	Type        string               `json:"type,omitempty" yaml:"type,omitempty"`
	Items       *Schema              `json:"items,omitempty" yaml:"items,omitempty"`
	Properties  map[string]Propertie `json:"properties,omitempty" yaml:"properties,omitempty"`
	Enum        []interface{}        `json:"enum,omitempty" yaml:"enum,omitempty"`
}

// Propertie are taken from the JSON Schema definition but their definitions were adjusted to the Swagger Specification
type Propertie struct {
	Ref                  string               `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Title                string               `json:"title,omitempty" yaml:"title,omitempty"`
	Description          string               `json:"description,omitempty" yaml:"description,omitempty"`
	Default              interface{}          `json:"default,omitempty" yaml:"default,omitempty"`
	Type                 string               `json:"type,omitempty" yaml:"type,omitempty"`
	Example              string               `json:"example,omitempty" yaml:"example,omitempty"`
	Required             []string             `json:"required,omitempty" yaml:"required,omitempty"`
	Format               string               `json:"format,omitempty" yaml:"format,omitempty"`
	ReadOnly             bool                 `json:"readOnly,omitempty" yaml:"readOnly,omitempty"`
	Properties           map[string]Propertie `json:"properties,omitempty" yaml:"properties,omitempty"`
	Items                *Propertie           `json:"items,omitempty" yaml:"items,omitempty"`
	AdditionalProperties *Propertie           `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
}

// Response as they are returned from executing this operation.
type Response struct {
	Description string  `json:"description" yaml:"description"`
	Schema      *Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
	Ref         string  `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}

// Security Allows the definition of a security scheme that can be used by the operations
type Security struct {
	Type             string            `json:"type,omitempty" yaml:"type,omitempty"` // Valid values are "basic", "apiKey" or "oauth2".
	Description      string            `json:"description,omitempty" yaml:"description,omitempty"`
	Name             string            `json:"name,omitempty" yaml:"name,omitempty"`
	In               string            `json:"in,omitempty" yaml:"in,omitempty"`     // Valid values are "query" or "header".
	Flow             string            `json:"flow,omitempty" yaml:"flow,omitempty"` // Valid values are "implicit", "password", "application" or "accessCode".
	AuthorizationURL string            `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"`
	TokenURL         string            `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
	Scopes           map[string]string `json:"scopes,omitempty" yaml:"scopes,omitempty"` // The available scopes for the OAuth2 security scheme.
}

// Tag Allows adding meta data to a single tag that is used by the Operation Object
type Tag struct {
	Name         string        `json:"name,omitempty" yaml:"name,omitempty"`
	Description  string        `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocs `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}

// ExternalDocs include Additional external documentation
type ExternalDocs struct {
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	URL         string `json:"url,omitempty" yaml:"url,omitempty"`
}
