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

// The request body contains details of the image list entry that you want to create.
type ImageListEntryPostRequest struct {

	// <p>User-defined parameters, in JSON format, that can be passed to an instance of this machine image when it is launched. This field can be used, for example, to specify the location of a database server and login details. Instance metadata, including user-defined data is available at http://192.0.0.192/ within an instance. See <a target=\"_blank\" href=\"http://www.oracle.com/pls/topic/lookup?ctx=stcomputecs&id=STCSG-GUID-268FE284-E5A0-4A18-BA58-345660925FB7\">Retrieving User-Defined Instance Attributes</a> in <em>Using Oracle Compute Cloud Service (IaaS)</em>.
	Attributes string `json:"attributes,omitempty"`

	// <p>A list of machine images. Specify the three-part name of each machine image.
	Machineimages []string `json:"machineimages,omitempty"`

	// The unique version of the entry in the image list.
	Version int32 `json:"version,omitempty"`
}