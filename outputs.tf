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

output "api_gateway_route_id" {
  description = "Identifier of the Route"
  value       = aws_apigatewayv2_route.route.id
}

output "api_gateway_route_key" {
  description = "Key of the Route"
  value       = aws_apigatewayv2_route.route.route_key
}
