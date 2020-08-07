// Copyright (c) 2020 WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
//
// WSO2 Inc. licenses this file to you under the Apache License,
// Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

public type Subscription record {|
    int id;
    int apiId = -1;
    int appId = -1;
    string policyId = "";
    string state = "";
    string tenantDomain = "carbon.super";
    int timestamp = 0;
|};

public type Application record {|
    int id;
    string name = "";
    string owner = "";
    string policyId = "";
    string tokenType = "";
    string tenantDomain = "carbon.super";
    json[] groupIds?;
    map<json> attributes?;
|};


public type KeyMap record {|
    int appId;
    string consumerKey = "";
    string keyType = "";
    string tenantDomain = "carbon.super";
    string keyManager = "Default";
|};

public type Api record {|
    int id;
    string provider = "";
    string name = "";
    string apiVersion = "";
    string context = "";
    string tenantDomain = "carbon.super";
    string policyId = "";
    any urlMaping?;
|};
