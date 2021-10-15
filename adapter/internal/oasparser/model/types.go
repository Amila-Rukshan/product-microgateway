/*
 *  Copyright (c) 2021, WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package model

// ProjectAPI contains the extracted from an API project zip
type ProjectAPI struct {
	APIYaml            APIYaml
	Deployments        []Deployment
	OpenAPIJsn         []byte
	UpstreamCerts      []byte
	InterceptorCerts   []byte
	APIType            string              // read from api.yaml and formatted to upper case
	APILifeCycleStatus string              // read from api.yaml and formatted to upper case
	ProductionEndpoint string              // read from api.yaml or env variable
	SandboxEndpoint    string              // read from api.yaml or env variable
	OrganizationID     string              // read from api.yaml or config
	EndpointSecurity   APIEndpointSecurity // derived from api.yaml or config
}

// EndpointSecurity contains parameters of endpoint security at api.json
type EndpointSecurity struct {
	Password         string `json:"password,omitempty"`
	Type             string `json:"type,omitempty"`
	Enabled          bool   `json:"enabled,omitempty"`
	Username         string `json:"username,omitempty"`
	CustomParameters string `json:"customparameters,omitempty"`
}

// APIEndpointSecurity represents the structure of endpoint_security param in api.yaml
type APIEndpointSecurity struct {
	Production EndpointSecurity
	Sandbox    EndpointSecurity
}

// ApimMeta represents APIM meta information of files received from APIM
type ApimMeta struct {
	Type    string `yaml:"type"`
	Version string `yaml:"version"`
}

// DeploymentEnvironments represents content of deployment_environments.yaml file
// of an API_CTL Project
type DeploymentEnvironments struct {
	ApimMeta
	Data []Deployment `yaml:"data"`
}

// Deployment represents deployment information of an API_CTL project
type Deployment struct {
	DisplayOnDevportal    bool   `yaml:"displayOnDevportal"`
	DeploymentVhost       string `yaml:"deploymentVhost"`
	DeploymentEnvironment string `yaml:"deploymentEnvironment"`
}

// APIYaml contains everything necessary to extract api.json/api.yaml file
// To support both api.json and api.yaml we convert yaml to json and then use json.Unmarshal()
// Therefore, the params are defined to support json.Unmarshal()
type APIYaml struct {
	ApimMeta
	Data struct {
		ID                         string   `json:"Id,omitempty"`
		Name                       string   `json:"name,omitempty"`
		Context                    string   `json:"context,omitempty"`
		Version                    string   `json:"version,omitempty"`
		RevisionID                 int      `json:"revisionId,omitempty"`
		APIType                    string   `json:"type,omitempty"`
		LifeCycleStatus            string   `json:"lifeCycleStatus,omitempty"`
		EndpointImplementationType string   `json:"endpointImplementationType,omitempty"`
		AuthorizationHeader        string   `json:"authorizationHeader,omitempty"`
		SecurityScheme             []string `json:"securityScheme,omitempty"`
		OrganizationID             string   `json:"organizationId,omitempty"`
		EndpointConfig             struct {
			EndpointType     string `json:"endpoint_type,omitempty"`
			EndpointSecurity struct {
				Production struct {
					Password         string `json:"password,omitempty"`
					Type             string `json:"type,omitempty"`
					Enabled          bool   `json:"enabled,omitempty"`
					Username         string `json:"username,omitempty"`
					CustomParameters string `json:"customparameters,omitempty"`
				} `json:"production,omitempty"`
				Sandbox struct {
					Password         string `json:"password,omitempty"`
					Type             string `json:"type,omitempty"`
					Enabled          bool   `json:"enabled,omitempty"`
					Username         string `json:"username,omitempty"`
					CustomParameters string `json:"customparameters,omitempty"`
				} `json:"sandbox,omitempty"`
			} `json:"endpoint_security,omitempty"`
			ProductionEndpoints struct {
				Endpoint string `json:"url,omitempty"`
				Config   struct {
					RetryDelay   string `json:"retryDelay,omitempty"`
					RetryTimeOut string `json:"retryTimeOut,omitempty"`
				} `json:"config,omitempty"`
			} `json:"production_endpoints,omitempty"`
			SandBoxEndpoints struct {
				Endpoint string `json:"url,omitempty"`
				Config   struct {
					RetryDelay   string `json:"retryDelay,omitempty"`
					RetryTimeOut string `json:"retryTimeOut,omitempty"`
				} `json:"config,omitempty"`
			} `json:"sandbox_endpoints,omitempty"`
		} `json:"endpointConfig,omitempty"`
	} `json:"data"`
}
