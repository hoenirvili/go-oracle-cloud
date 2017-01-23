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

type IpReservationResponse struct {

	// Shows the default account for your identity domain.
	Account string `json:"account,omitempty"`

	// Public IP address.<p>An IP reservation is a public IP address that you can attach to an Oracle Compute Cloud Service instance that requires access to or from the Internet.
	Ip string `json:"ip,omitempty"`

	// <p>The three-part name of the object (<code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object</em></code>).
	Name string `json:"name,omitempty"`

	// <code>/oracle/public/ippool</code><p>Pool of public IP addresses
	Parentpool string `json:"parentpool,omitempty"`

	// <code>true</code> indicates that the IP reservation has a persistent public IP address. You can associate either a temporary or a persistent public IP address with an instance when you create the instance.<p>Temporary public IP addresses are assigned dynamically from a pool of public IP addresses. When you associate a temporary public IP address with an instance, if the instance is restarted or is deleted and created again later, its public IP address might change.
	Permanent bool `json:"permanent,omitempty"`

	// Not used
	Quota string `json:"quota,omitempty"`

	// A comma-separated list of strings which helps you to identify IP reservation.
	Tags []string `json:"tags,omitempty"`

	// Uniform Resource Identifier
	Uri string `json:"uri,omitempty"`

	// <code>true</code> indicates that the IP reservation is associated with an instance.
	Used bool `json:"used,omitempty"`
}