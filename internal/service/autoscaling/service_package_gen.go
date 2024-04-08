// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package autoscaling

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	autoscaling_sdkv2 "github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceGroup,
			TypeName: "aws_autoscaling_group",
		},
		{
			Factory:  DataSourceGroups,
			TypeName: "aws_autoscaling_groups",
		},
		{
			Factory:  DataSourceLaunchConfiguration,
			TypeName: "aws_launch_configuration",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAttachment,
			TypeName: "aws_autoscaling_attachment",
			Name:     "Attachment",
		},
		{
			Factory:  ResourceGroup,
			TypeName: "aws_autoscaling_group",
		},
		{
			Factory:  ResourceGroupTag,
			TypeName: "aws_autoscaling_group_tag",
		},
		{
			Factory:  ResourceLifecycleHook,
			TypeName: "aws_autoscaling_lifecycle_hook",
		},
		{
			Factory:  ResourceNotification,
			TypeName: "aws_autoscaling_notification",
		},
		{
			Factory:  ResourcePolicy,
			TypeName: "aws_autoscaling_policy",
		},
		{
			Factory:  ResourceSchedule,
			TypeName: "aws_autoscaling_schedule",
		},
		{
			Factory:  ResourceTrafficSourceAttachment,
			TypeName: "aws_autoscaling_traffic_source_attachment",
		},
		{
			Factory:  ResourceLaunchConfiguration,
			TypeName: "aws_launch_configuration",
			Name:     "Launch Configuration",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.AutoScaling
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*autoscaling_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return autoscaling_sdkv2.NewFromConfig(cfg, func(o *autoscaling_sdkv2.Options) {
		if endpoint := config["endpoint"].(string); endpoint != "" {
			o.BaseEndpoint = aws_sdkv2.String(endpoint)
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
