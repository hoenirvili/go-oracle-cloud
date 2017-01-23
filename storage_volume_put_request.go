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

// The request body contains details of the storage volume that you want to update. 
type StorageVolumePutRequest struct {

	// <p>The description of the storage volume.
	Description string `json:"description,omitempty"`

	// The three-part name of the object (<code>/Compute-<em>identity_domain</em>/<em>user</em>/<em>object</em></code>).
	Name string `json:"name,omitempty"`

	// Specify the operating system platform for a bootable storage volume, such as Linux or Windows.
	Platform string `json:"platform,omitempty"`

	// <p>The storage-pool property.<p>For storage volumes that require low latency and high IOPS, such as for storing database files, specify <code>/oracle/public/storage/latency</code>.<p>For all other storage volumes, specify <code>/oracle/public/storage/default</code>.
	Properties []string `json:"properties,omitempty"`

	// <p>The size of this storage volume. Use one of the following abbreviations for the unit of measurement:<ul><li><code>B</code> or <code>b</code> for bytes</li><li><code>K</code> or <code>k</code> for kilobytes</li><li><code>M</code> or <code>m</code> for megabytes</li><li><code>G</code> or <code>g</code> for gigabytes</li><li><code>T</code> or <code>t</code> for terabytes</li></ul><p>For example, to create a volume of size 10 gigabytes, you can specify <code>10G</code>, or <code>10240M</code>, or <code>10485760K</code>, and so on.<p>The allowed range is from 1 GB to 2 TB, in increments of 1 GB.
	Size string `json:"size,omitempty"`

	// Multipart name of the storage volume snapshot if this storage volume is a clone.
	Snapshot string `json:"snapshot,omitempty"`

	// Account of the parent snapshot from which the storage volume is restored.
	SnapshotAccount string `json:"snapshot_account,omitempty"`

	// Id of the parent snapshot from which the storage volume is restored or cloned.
	SnapshotId string `json:"snapshot_id,omitempty"`

	// <p>Strings that you can use to tag the storage volume.
	Tags []string `json:"tags,omitempty"`
}