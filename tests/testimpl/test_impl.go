package testimpl

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	apitypes "github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	awsClient := GetAWSApiGatewayV2Client(t)

	t.Run("TestApiGatewayV2Exists", func(t *testing.T) {
		awsApiGatewayId := terraform.Output(t, ctx.TerratestTerraformOptions(), "api_gateway_id")
		awsApiGatewayProtocolType := terraform.Output(t, ctx.TerratestTerraformOptions(), "api_protocol_type")

		apiGateway, err := awsClient.GetApi(context.TODO(), &apigatewayv2.GetApiInput{
			ApiId: &awsApiGatewayId,
		})
		if err != nil {
			t.Errorf("Failure during GetApi: %v", err)
		}

		assert.Equal(t, *apiGateway.ApiId, awsApiGatewayId, "Expected ID did not match actual ID!")
		assert.Equal(t, apiGateway.ProtocolType, apitypes.ProtocolType(awsApiGatewayProtocolType), "Expected protocol type did not match actual!")
	})

	t.Run("TestApiGatewayV2RouteExists", func(t *testing.T) {
		awsApiGatewayId := terraform.Output(t, ctx.TerratestTerraformOptions(), "api_gateway_id")
		awsApiGatewayRouteId := terraform.Output(t, ctx.TerratestTerraformOptions(), "api_gateway_route_id")
		awsApiGatewayRouteKey := terraform.Output(t, ctx.TerratestTerraformOptions(), "api_gateway_route_key")

		apiGatewayRoute, err := awsClient.GetRoute(context.TODO(), &apigatewayv2.GetRouteInput{
			ApiId:   &awsApiGatewayId,
			RouteId: &awsApiGatewayRouteId,
		})
		if err != nil {
			t.Errorf("Failure during GetApi: %v", err)
		}

		assert.Equal(t, *apiGatewayRoute.RouteId, awsApiGatewayRouteId, "Expected Route ID did not match actual Route ID!")
		assert.Equal(t, *apiGatewayRoute.RouteKey, awsApiGatewayRouteKey, "Expected Route Key did not match actual Route Key!")
	})
}

func GetAWSApiGatewayV2Client(t *testing.T) *apigatewayv2.Client {
	awsApiGatewayV2Client := apigatewayv2.NewFromConfig(GetAWSConfig(t))
	return awsApiGatewayV2Client
}

func GetAWSConfig(t *testing.T) (cfg aws.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	require.NoErrorf(t, err, "unable to load SDK config, %v", err)
	return cfg
}
