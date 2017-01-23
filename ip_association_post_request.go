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

// The request body contains details of the IP association that you want to create. 
type IpAssociationPostRequest struct {

	// <ul><li>To associate a temporary IP address from the pool, specify ippool:/oracle/public/ippool.</li><li>To associate a persistent IP address, specify ipreservation:ipreservation_name, where ipreservation_name is three-part name of an existing IP reservation in the <code>/Compute-identity_domain/user/object_name</code> format. For more information about how to create an IP reservation, see <a class=\"xref\" href=\"op-ip-reservation--post.html\">Create an IP Reservation</a>.</li><ul>
	Parentpool string `json:"parentpool,omitempty"`

	// <p>The three-part name of the vcable ID of the instance that you want to associate with an IP address. The three-part name is in the format: <code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object</em></code>.<p>For more information about the vcable of an instance, see <a class=\"xref\" href=\"op-instance-{name}-get.html\">Retrieve Details of an Instance</a>.
	Vcable string `json:"vcable,omitempty"`
}