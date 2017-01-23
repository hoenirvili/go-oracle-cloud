/* 
 * REST API for Oracle Compute Cloud Service (IaaS)
 *
 * Use the Oracle Compute Cloud Service REST API to provision and manage instances and the associated resources
 *
 * OpenAPI spec version: 1.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

import (
	"net/url"
	"encoding/json"
	"fmt"
	"strings"
)

type InstanceConsoleApi struct {
	Configuration Configuration
}

func NewInstanceConsoleApi() *InstanceConsoleApi {
	configuration := NewConfiguration()
	return &InstanceConsoleApi{
		Configuration: *configuration,
	}
}

func NewInstanceConsoleApiWithBasePath(basePath string) *InstanceConsoleApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &InstanceConsoleApi{
		Configuration: *configuration,
	}
}

/**
 * Retrieve Details of an Instance Console
 * Retrieves the messages that appear when an instance boots. Use these messages to diagnose unresponsive instances and failures in the boot up process.
 *
 * @param name Name of the instance in one of the following format: &lt;code&gt;/Compute-&lt;em&gt;identityDomain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;id&lt;/em&gt;&lt;/code&gt; or &lt;code&gt;/Compute-&lt;em&gt;identityDomain&lt;/em&gt;/&lt;em&gt;user&lt;/em&gt;/&lt;em&gt;provided_name&lt;/em&gt;/&lt;em&gt;id&lt;/em&gt;&lt;/code&gt;, where id is an autogenerated ID.
 * @return *InstanceConsoleResponse
 */
func (a InstanceConsoleApi) GetInstanceConsole(name string) (*InstanceConsoleResponse, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/instanceconsole/{name}"
	path = strings.Replace(path, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/oracle-compute-v3+json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/oracle-compute-v3+json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(InstanceConsoleResponse)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}
