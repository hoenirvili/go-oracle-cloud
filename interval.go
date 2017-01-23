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

// There are two kinds of Intervals. Each Interval has its own JSON format. Your Interval field should look like one of the following:<p><ul><li><code>\"interval\":{\"Hourly\":{\"hourlyInterval\":2}}</code></li><li><code>{\"DailyWeekly\":{\"daysOfWeek\":[\"MONDAY\"],\"timeOfDay\":\"03:15\",\"userTimeZone\":\"America/Los_Angeles\"}}</code></li></ul><p>Days of the week is any day of the week fully capitalized (MONDAY, TUESDAY, etc). The user time zone is any IANA user timezone.  For example user time zones see <a href=\"https://en.wikipedia.org/wiki/IANA_time_zone_database\">List of IANA time zones</a>.
type Interval struct {
}