// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

variable "api_id" {
  description = "ID of the API Gateway to which this stage will attach"
  type        = string
}

variable "route_key" {
  description = "Route key for the route. For HTTP APIs, the route key can be either `$default`, or a combination of an HTTP method and resource path, for example, `GET /pets`."
  type        = string
  default     = "$default"
}

variable "target" {
  description = "Target for the route, of the form `integrations/IntegrationID`, where `IntegrationID` is the identifier of an `aws_apigatewayv2_integration` resource."
  type        = string
  default     = null
}
