/*
 * Red Hat Advanced Cluster Security Service Fleet Manager
 *
 * Red Hat Advanced Cluster Security (RHACS) Service Fleet Manager is a Rest API to manage instances of ACS components.
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

import (
	"time"
)

// CentralRequest struct for CentralRequest
type CentralRequest struct {
	Id   string `json:"id,omitempty"`
	Kind string `json:"kind,omitempty"`
	Href string `json:"href,omitempty"`
	// Values: [accepted, preparing, provisioning, ready, failed, deprovision, deleting]
	Status string `json:"status,omitempty"`
	// Name of Cloud used to deploy. For example AWS
	CloudProvider string `json:"cloud_provider,omitempty"`
	MultiAz       bool   `json:"multi_az"`
	// Values will be regions of specific cloud provider. For example: us-east-1 for AWS
	Region       string    `json:"region,omitempty"`
	Owner        string    `json:"owner,omitempty"`
	Name         string    `json:"name,omitempty"`
	UiHost       string    `json:"ui_host,omitempty"`
	DataHost     string    `json:"data_host,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	FailedReason string    `json:"failed_reason,omitempty"`
	Version      string    `json:"version,omitempty"`
	InstanceType string    `json:"instance_type,omitempty"`
}
