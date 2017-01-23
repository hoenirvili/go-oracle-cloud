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

type StorageSnapshotResponse struct {

	// Account to use for snapshots
	Account string `json:"account,omitempty"`

	// Description of this storage snapshot.
	Description string `json:"description,omitempty"`

	// The name of the machine image that's used in the boot volume from which this snapshot is taken
	MachineimageName string `json:"machineimage_name,omitempty"`

	// Multipart name of the storage snapshot.
	Name string `json:"name,omitempty"`

	// The OS platform this snapshot is compatible with.
	Platform string `json:"platform,omitempty"`

	// <ul><li><code>/oracle/private/storage/snapshot/collocated</code> indicates that the snapshot is stored in the same physical location as the original storage volume.</li><li><code>/oracle/public/storage/snapshot/default</code> indicates that the snapshot aren't stored in the same location as the original storage volume. Instead, they are reduced and stored in the associated Oracle Storage Cloud Service instance.</li></ul>
	Property string `json:"property,omitempty"`

	// The size of the storage snapshot in bytes.
	Size string `json:"size,omitempty"`

	// A unique id for this snapshot auto-generated by the server.
	SnapshotId string `json:"snapshot_id,omitempty"`

	// Timestamp of the storage snapshot, generated by storage server. The snapshot will contain data written to the original volume before this time.
	SnapshotTimestamp string `json:"snapshot_timestamp,omitempty"`

	// Timestamp when the operation is started.
	StartTimestamp string `json:"start_timestamp,omitempty"`

	// Current status of the storage snapshot.
	Status string `json:"status,omitempty"`

	// Details about the latest state of the storage volume snapshot.
	StatusDetail string `json:"status_detail,omitempty"`

	// Indicates the time that the current view of the storage volume snapshot was generated.
	StatusTimestamp string `json:"status_timestamp,omitempty"`

	// Comma-separated list of strings that describe the storage snapshot.
	Tags []string `json:"tags,omitempty"`

	// Uniform Resource Identifier
	Uri string `json:"uri,omitempty"`

	// The name of the parent storage volume of this storage snapshot.
	Volume string `json:"volume,omitempty"`
}