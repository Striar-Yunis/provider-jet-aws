/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/crossplane-contrib/provider-jet-aws/apis/accessanalyzer/v1alpha1"
	v1alpha1acm "github.com/crossplane-contrib/provider-jet-aws/apis/acm/v1alpha1"
	v1alpha1acmpca "github.com/crossplane-contrib/provider-jet-aws/apis/acmpca/v1alpha1"
	v1alpha1amp "github.com/crossplane-contrib/provider-jet-aws/apis/amp/v1alpha1"
	v1alpha1amplify "github.com/crossplane-contrib/provider-jet-aws/apis/amplify/v1alpha1"
	v1alpha1apigateway "github.com/crossplane-contrib/provider-jet-aws/apis/apigateway/v1alpha1"
	v1alpha1apigatewayv2 "github.com/crossplane-contrib/provider-jet-aws/apis/apigatewayv2/v1alpha1"
	v1alpha1appautoscaling "github.com/crossplane-contrib/provider-jet-aws/apis/appautoscaling/v1alpha1"
	v1alpha1appconfig "github.com/crossplane-contrib/provider-jet-aws/apis/appconfig/v1alpha1"
	v1alpha1appmesh "github.com/crossplane-contrib/provider-jet-aws/apis/appmesh/v1alpha1"
	v1alpha1apprunner "github.com/crossplane-contrib/provider-jet-aws/apis/apprunner/v1alpha1"
	v1alpha1appstream "github.com/crossplane-contrib/provider-jet-aws/apis/appstream/v1alpha1"
	v1alpha1appsync "github.com/crossplane-contrib/provider-jet-aws/apis/appsync/v1alpha1"
	v1alpha1athena "github.com/crossplane-contrib/provider-jet-aws/apis/athena/v1alpha1"
	v1alpha1autoscaling "github.com/crossplane-contrib/provider-jet-aws/apis/autoscaling/v1alpha1"
	v1alpha2 "github.com/crossplane-contrib/provider-jet-aws/apis/autoscaling/v1alpha2"
	v1alpha1autoscalingplans "github.com/crossplane-contrib/provider-jet-aws/apis/autoscalingplans/v1alpha1"
	v1alpha1backup "github.com/crossplane-contrib/provider-jet-aws/apis/backup/v1alpha1"
	v1alpha1batch "github.com/crossplane-contrib/provider-jet-aws/apis/batch/v1alpha1"
	v1alpha1budgets "github.com/crossplane-contrib/provider-jet-aws/apis/budgets/v1alpha1"
	v1alpha1chime "github.com/crossplane-contrib/provider-jet-aws/apis/chime/v1alpha1"
	v1alpha1cloud9 "github.com/crossplane-contrib/provider-jet-aws/apis/cloud9/v1alpha1"
	v1alpha1cloudformation "github.com/crossplane-contrib/provider-jet-aws/apis/cloudformation/v1alpha1"
	v1alpha1cloudfront "github.com/crossplane-contrib/provider-jet-aws/apis/cloudfront/v1alpha1"
	v1alpha1cloudhsmv2 "github.com/crossplane-contrib/provider-jet-aws/apis/cloudhsmv2/v1alpha1"
	v1alpha1cloudtrail "github.com/crossplane-contrib/provider-jet-aws/apis/cloudtrail/v1alpha1"
	v1alpha1cloudwatch "github.com/crossplane-contrib/provider-jet-aws/apis/cloudwatch/v1alpha1"
	v1alpha1cloudwatchlogs "github.com/crossplane-contrib/provider-jet-aws/apis/cloudwatchlogs/v1alpha1"
	v1alpha1codeartifact "github.com/crossplane-contrib/provider-jet-aws/apis/codeartifact/v1alpha1"
	v1alpha1codebuild "github.com/crossplane-contrib/provider-jet-aws/apis/codebuild/v1alpha1"
	v1alpha1codecommit "github.com/crossplane-contrib/provider-jet-aws/apis/codecommit/v1alpha1"
	v1alpha1codedeploy "github.com/crossplane-contrib/provider-jet-aws/apis/codedeploy/v1alpha1"
	v1alpha1codepipeline "github.com/crossplane-contrib/provider-jet-aws/apis/codepipeline/v1alpha1"
	v1alpha1codestarconnections "github.com/crossplane-contrib/provider-jet-aws/apis/codestarconnections/v1alpha1"
	v1alpha1codestarnotifications "github.com/crossplane-contrib/provider-jet-aws/apis/codestarnotifications/v1alpha1"
	v1alpha1cognitoidentity "github.com/crossplane-contrib/provider-jet-aws/apis/cognitoidentity/v1alpha1"
	v1alpha1cognitoidp "github.com/crossplane-contrib/provider-jet-aws/apis/cognitoidp/v1alpha1"
	v1alpha1configservice "github.com/crossplane-contrib/provider-jet-aws/apis/configservice/v1alpha1"
	v1alpha1cur "github.com/crossplane-contrib/provider-jet-aws/apis/cur/v1alpha1"
	v1alpha1datapipeline "github.com/crossplane-contrib/provider-jet-aws/apis/datapipeline/v1alpha1"
	v1alpha1datasync "github.com/crossplane-contrib/provider-jet-aws/apis/datasync/v1alpha1"
	v1alpha1dax "github.com/crossplane-contrib/provider-jet-aws/apis/dax/v1alpha1"
	v1alpha1devicefarm "github.com/crossplane-contrib/provider-jet-aws/apis/devicefarm/v1alpha1"
	v1alpha1directconnect "github.com/crossplane-contrib/provider-jet-aws/apis/directconnect/v1alpha1"
	v1alpha1dlm "github.com/crossplane-contrib/provider-jet-aws/apis/dlm/v1alpha1"
	v1alpha1dms "github.com/crossplane-contrib/provider-jet-aws/apis/dms/v1alpha1"
	v1alpha1docdb "github.com/crossplane-contrib/provider-jet-aws/apis/docdb/v1alpha1"
	v1alpha1ds "github.com/crossplane-contrib/provider-jet-aws/apis/ds/v1alpha1"
	v1alpha1dynamodb "github.com/crossplane-contrib/provider-jet-aws/apis/dynamodb/v1alpha1"
	v1alpha1ec2 "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha1"
	v1alpha2ec2 "github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2"
	v1alpha1ecr "github.com/crossplane-contrib/provider-jet-aws/apis/ecr/v1alpha1"
	v1alpha2ecr "github.com/crossplane-contrib/provider-jet-aws/apis/ecr/v1alpha2"
	v1alpha2ecrpublic "github.com/crossplane-contrib/provider-jet-aws/apis/ecrpublic/v1alpha2"
	v1alpha2ecs "github.com/crossplane-contrib/provider-jet-aws/apis/ecs/v1alpha2"
	v1alpha1efs "github.com/crossplane-contrib/provider-jet-aws/apis/efs/v1alpha1"
	v1alpha2eks "github.com/crossplane-contrib/provider-jet-aws/apis/eks/v1alpha2"
	v1alpha1elasticache "github.com/crossplane-contrib/provider-jet-aws/apis/elasticache/v1alpha1"
	v1alpha2elasticache "github.com/crossplane-contrib/provider-jet-aws/apis/elasticache/v1alpha2"
	v1alpha1elasticbeanstalk "github.com/crossplane-contrib/provider-jet-aws/apis/elasticbeanstalk/v1alpha1"
	v1alpha1elasticsearch "github.com/crossplane-contrib/provider-jet-aws/apis/elasticsearch/v1alpha1"
	v1alpha1elastictranscoder "github.com/crossplane-contrib/provider-jet-aws/apis/elastictranscoder/v1alpha1"
	v1alpha1elb "github.com/crossplane-contrib/provider-jet-aws/apis/elb/v1alpha1"
	v1alpha1elbv2 "github.com/crossplane-contrib/provider-jet-aws/apis/elbv2/v1alpha1"
	v1alpha2elbv2 "github.com/crossplane-contrib/provider-jet-aws/apis/elbv2/v1alpha2"
	v1alpha1emr "github.com/crossplane-contrib/provider-jet-aws/apis/emr/v1alpha1"
	v1alpha1events "github.com/crossplane-contrib/provider-jet-aws/apis/events/v1alpha1"
	v1alpha1firehose "github.com/crossplane-contrib/provider-jet-aws/apis/firehose/v1alpha1"
	v1alpha1fms "github.com/crossplane-contrib/provider-jet-aws/apis/fms/v1alpha1"
	v1alpha1fsx "github.com/crossplane-contrib/provider-jet-aws/apis/fsx/v1alpha1"
	v1alpha1gamelift "github.com/crossplane-contrib/provider-jet-aws/apis/gamelift/v1alpha1"
	v1alpha1glacier "github.com/crossplane-contrib/provider-jet-aws/apis/glacier/v1alpha1"
	v1alpha2globalaccelerator "github.com/crossplane-contrib/provider-jet-aws/apis/globalaccelerator/v1alpha2"
	v1alpha1glue "github.com/crossplane-contrib/provider-jet-aws/apis/glue/v1alpha1"
	v1alpha1guardduty "github.com/crossplane-contrib/provider-jet-aws/apis/guardduty/v1alpha1"
	v1alpha1iam "github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha1"
	v1alpha2iam "github.com/crossplane-contrib/provider-jet-aws/apis/iam/v1alpha2"
	v1alpha1imagebuilder "github.com/crossplane-contrib/provider-jet-aws/apis/imagebuilder/v1alpha1"
	v1alpha1inspector "github.com/crossplane-contrib/provider-jet-aws/apis/inspector/v1alpha1"
	v1alpha1iot "github.com/crossplane-contrib/provider-jet-aws/apis/iot/v1alpha1"
	v1alpha1kafka "github.com/crossplane-contrib/provider-jet-aws/apis/kafka/v1alpha1"
	v1alpha1kinesis "github.com/crossplane-contrib/provider-jet-aws/apis/kinesis/v1alpha1"
	v1alpha1kinesisanalytics "github.com/crossplane-contrib/provider-jet-aws/apis/kinesisanalytics/v1alpha1"
	v1alpha1kinesisanalyticsv2 "github.com/crossplane-contrib/provider-jet-aws/apis/kinesisanalyticsv2/v1alpha1"
	v1alpha1kinesisvideo "github.com/crossplane-contrib/provider-jet-aws/apis/kinesisvideo/v1alpha1"
	v1alpha1kms "github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha1"
	v1alpha2kms "github.com/crossplane-contrib/provider-jet-aws/apis/kms/v1alpha2"
	v1alpha1lakeformation "github.com/crossplane-contrib/provider-jet-aws/apis/lakeformation/v1alpha1"
	v1alpha1lambda "github.com/crossplane-contrib/provider-jet-aws/apis/lambda/v1alpha1"
	v1alpha1lexmodels "github.com/crossplane-contrib/provider-jet-aws/apis/lexmodels/v1alpha1"
	v1alpha1licensemanager "github.com/crossplane-contrib/provider-jet-aws/apis/licensemanager/v1alpha1"
	v1alpha1lightsail "github.com/crossplane-contrib/provider-jet-aws/apis/lightsail/v1alpha1"
	v1alpha1macie "github.com/crossplane-contrib/provider-jet-aws/apis/macie/v1alpha1"
	v1alpha1macie2 "github.com/crossplane-contrib/provider-jet-aws/apis/macie2/v1alpha1"
	v1alpha1mediaconvert "github.com/crossplane-contrib/provider-jet-aws/apis/mediaconvert/v1alpha1"
	v1alpha1mediapackage "github.com/crossplane-contrib/provider-jet-aws/apis/mediapackage/v1alpha1"
	v1alpha1mediastore "github.com/crossplane-contrib/provider-jet-aws/apis/mediastore/v1alpha1"
	v1alpha2mq "github.com/crossplane-contrib/provider-jet-aws/apis/mq/v1alpha2"
	v1alpha2neptune "github.com/crossplane-contrib/provider-jet-aws/apis/neptune/v1alpha2"
	v1alpha1networkfirewall "github.com/crossplane-contrib/provider-jet-aws/apis/networkfirewall/v1alpha1"
	v1alpha1opsworks "github.com/crossplane-contrib/provider-jet-aws/apis/opsworks/v1alpha1"
	v1alpha1organizations "github.com/crossplane-contrib/provider-jet-aws/apis/organizations/v1alpha1"
	v1alpha1pinpoint "github.com/crossplane-contrib/provider-jet-aws/apis/pinpoint/v1alpha1"
	v1alpha1qldb "github.com/crossplane-contrib/provider-jet-aws/apis/qldb/v1alpha1"
	v1alpha1quicksight "github.com/crossplane-contrib/provider-jet-aws/apis/quicksight/v1alpha1"
	v1alpha1ram "github.com/crossplane-contrib/provider-jet-aws/apis/ram/v1alpha1"
	v1alpha1rds "github.com/crossplane-contrib/provider-jet-aws/apis/rds/v1alpha1"
	v1alpha2rds "github.com/crossplane-contrib/provider-jet-aws/apis/rds/v1alpha2"
	v1alpha1redshift "github.com/crossplane-contrib/provider-jet-aws/apis/redshift/v1alpha1"
	v1alpha1resourcegroups "github.com/crossplane-contrib/provider-jet-aws/apis/resourcegroups/v1alpha1"
	v1alpha2route53 "github.com/crossplane-contrib/provider-jet-aws/apis/route53/v1alpha2"
	v1alpha1route53recoverycontrolconfig "github.com/crossplane-contrib/provider-jet-aws/apis/route53recoverycontrolconfig/v1alpha1"
	v1alpha1route53recoveryreadiness "github.com/crossplane-contrib/provider-jet-aws/apis/route53recoveryreadiness/v1alpha1"
	v1alpha1route53resolver "github.com/crossplane-contrib/provider-jet-aws/apis/route53resolver/v1alpha1"
	v1alpha1s3 "github.com/crossplane-contrib/provider-jet-aws/apis/s3/v1alpha1"
	v1alpha2s3 "github.com/crossplane-contrib/provider-jet-aws/apis/s3/v1alpha2"
	v1alpha1s3control "github.com/crossplane-contrib/provider-jet-aws/apis/s3control/v1alpha1"
	v1alpha1s3outposts "github.com/crossplane-contrib/provider-jet-aws/apis/s3outposts/v1alpha1"
	v1alpha1sagemaker "github.com/crossplane-contrib/provider-jet-aws/apis/sagemaker/v1alpha1"
	v1alpha1schemas "github.com/crossplane-contrib/provider-jet-aws/apis/schemas/v1alpha1"
	v1alpha1secretsmanager "github.com/crossplane-contrib/provider-jet-aws/apis/secretsmanager/v1alpha1"
	v1alpha1securityhub "github.com/crossplane-contrib/provider-jet-aws/apis/securityhub/v1alpha1"
	v1alpha1serverlessrepo "github.com/crossplane-contrib/provider-jet-aws/apis/serverlessrepo/v1alpha1"
	v1alpha1servicecatalog "github.com/crossplane-contrib/provider-jet-aws/apis/servicecatalog/v1alpha1"
	v1alpha1servicediscovery "github.com/crossplane-contrib/provider-jet-aws/apis/servicediscovery/v1alpha1"
	v1alpha1servicequotas "github.com/crossplane-contrib/provider-jet-aws/apis/servicequotas/v1alpha1"
	v1alpha1ses "github.com/crossplane-contrib/provider-jet-aws/apis/ses/v1alpha1"
	v1alpha1sfn "github.com/crossplane-contrib/provider-jet-aws/apis/sfn/v1alpha1"
	v1alpha1shield "github.com/crossplane-contrib/provider-jet-aws/apis/shield/v1alpha1"
	v1alpha1signer "github.com/crossplane-contrib/provider-jet-aws/apis/signer/v1alpha1"
	v1alpha1simpledb "github.com/crossplane-contrib/provider-jet-aws/apis/simpledb/v1alpha1"
	v1alpha1sns "github.com/crossplane-contrib/provider-jet-aws/apis/sns/v1alpha1"
	v1alpha1sqs "github.com/crossplane-contrib/provider-jet-aws/apis/sqs/v1alpha1"
	v1alpha1ssm "github.com/crossplane-contrib/provider-jet-aws/apis/ssm/v1alpha1"
	v1alpha1ssoadmin "github.com/crossplane-contrib/provider-jet-aws/apis/ssoadmin/v1alpha1"
	v1alpha1storagegateway "github.com/crossplane-contrib/provider-jet-aws/apis/storagegateway/v1alpha1"
	v1alpha1swf "github.com/crossplane-contrib/provider-jet-aws/apis/swf/v1alpha1"
	v1alpha1synthetics "github.com/crossplane-contrib/provider-jet-aws/apis/synthetics/v1alpha1"
	v1alpha1timestreamwrite "github.com/crossplane-contrib/provider-jet-aws/apis/timestreamwrite/v1alpha1"
	v1alpha1transfer "github.com/crossplane-contrib/provider-jet-aws/apis/transfer/v1alpha1"
	v1alpha1apis "github.com/crossplane-contrib/provider-jet-aws/apis/v1alpha1"
	v1alpha1waf "github.com/crossplane-contrib/provider-jet-aws/apis/waf/v1alpha1"
	v1alpha1wafregional "github.com/crossplane-contrib/provider-jet-aws/apis/wafregional/v1alpha1"
	v1alpha1wafv2 "github.com/crossplane-contrib/provider-jet-aws/apis/wafv2/v1alpha1"
	v1alpha1worklink "github.com/crossplane-contrib/provider-jet-aws/apis/worklink/v1alpha1"
	v1alpha1workspaces "github.com/crossplane-contrib/provider-jet-aws/apis/workspaces/v1alpha1"
	v1alpha1xray "github.com/crossplane-contrib/provider-jet-aws/apis/xray/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1acm.SchemeBuilder.AddToScheme,
		v1alpha1acmpca.SchemeBuilder.AddToScheme,
		v1alpha1amp.SchemeBuilder.AddToScheme,
		v1alpha1amplify.SchemeBuilder.AddToScheme,
		v1alpha1apigateway.SchemeBuilder.AddToScheme,
		v1alpha1apigatewayv2.SchemeBuilder.AddToScheme,
		v1alpha1appautoscaling.SchemeBuilder.AddToScheme,
		v1alpha1appconfig.SchemeBuilder.AddToScheme,
		v1alpha1appmesh.SchemeBuilder.AddToScheme,
		v1alpha1apprunner.SchemeBuilder.AddToScheme,
		v1alpha1appstream.SchemeBuilder.AddToScheme,
		v1alpha1appsync.SchemeBuilder.AddToScheme,
		v1alpha1athena.SchemeBuilder.AddToScheme,
		v1alpha1autoscaling.SchemeBuilder.AddToScheme,
		v1alpha2.SchemeBuilder.AddToScheme,
		v1alpha1autoscalingplans.SchemeBuilder.AddToScheme,
		v1alpha1backup.SchemeBuilder.AddToScheme,
		v1alpha1batch.SchemeBuilder.AddToScheme,
		v1alpha1budgets.SchemeBuilder.AddToScheme,
		v1alpha1chime.SchemeBuilder.AddToScheme,
		v1alpha1cloud9.SchemeBuilder.AddToScheme,
		v1alpha1cloudformation.SchemeBuilder.AddToScheme,
		v1alpha1cloudfront.SchemeBuilder.AddToScheme,
		v1alpha1cloudhsmv2.SchemeBuilder.AddToScheme,
		v1alpha1cloudtrail.SchemeBuilder.AddToScheme,
		v1alpha1cloudwatch.SchemeBuilder.AddToScheme,
		v1alpha1cloudwatchlogs.SchemeBuilder.AddToScheme,
		v1alpha1codeartifact.SchemeBuilder.AddToScheme,
		v1alpha1codebuild.SchemeBuilder.AddToScheme,
		v1alpha1codecommit.SchemeBuilder.AddToScheme,
		v1alpha1codedeploy.SchemeBuilder.AddToScheme,
		v1alpha1codepipeline.SchemeBuilder.AddToScheme,
		v1alpha1codestarconnections.SchemeBuilder.AddToScheme,
		v1alpha1codestarnotifications.SchemeBuilder.AddToScheme,
		v1alpha1cognitoidentity.SchemeBuilder.AddToScheme,
		v1alpha1cognitoidp.SchemeBuilder.AddToScheme,
		v1alpha1configservice.SchemeBuilder.AddToScheme,
		v1alpha1cur.SchemeBuilder.AddToScheme,
		v1alpha1datapipeline.SchemeBuilder.AddToScheme,
		v1alpha1datasync.SchemeBuilder.AddToScheme,
		v1alpha1dax.SchemeBuilder.AddToScheme,
		v1alpha1devicefarm.SchemeBuilder.AddToScheme,
		v1alpha1directconnect.SchemeBuilder.AddToScheme,
		v1alpha1dlm.SchemeBuilder.AddToScheme,
		v1alpha1dms.SchemeBuilder.AddToScheme,
		v1alpha1docdb.SchemeBuilder.AddToScheme,
		v1alpha1ds.SchemeBuilder.AddToScheme,
		v1alpha1dynamodb.SchemeBuilder.AddToScheme,
		v1alpha1ec2.SchemeBuilder.AddToScheme,
		v1alpha2ec2.SchemeBuilder.AddToScheme,
		v1alpha1ecr.SchemeBuilder.AddToScheme,
		v1alpha2ecr.SchemeBuilder.AddToScheme,
		v1alpha2ecrpublic.SchemeBuilder.AddToScheme,
		v1alpha2ecs.SchemeBuilder.AddToScheme,
		v1alpha1efs.SchemeBuilder.AddToScheme,
		v1alpha2eks.SchemeBuilder.AddToScheme,
		v1alpha1elasticache.SchemeBuilder.AddToScheme,
		v1alpha2elasticache.SchemeBuilder.AddToScheme,
		v1alpha1elasticbeanstalk.SchemeBuilder.AddToScheme,
		v1alpha1elasticsearch.SchemeBuilder.AddToScheme,
		v1alpha1elastictranscoder.SchemeBuilder.AddToScheme,
		v1alpha1elb.SchemeBuilder.AddToScheme,
		v1alpha1elbv2.SchemeBuilder.AddToScheme,
		v1alpha2elbv2.SchemeBuilder.AddToScheme,
		v1alpha1emr.SchemeBuilder.AddToScheme,
		v1alpha1events.SchemeBuilder.AddToScheme,
		v1alpha1firehose.SchemeBuilder.AddToScheme,
		v1alpha1fms.SchemeBuilder.AddToScheme,
		v1alpha1fsx.SchemeBuilder.AddToScheme,
		v1alpha1gamelift.SchemeBuilder.AddToScheme,
		v1alpha1glacier.SchemeBuilder.AddToScheme,
		v1alpha2globalaccelerator.SchemeBuilder.AddToScheme,
		v1alpha1glue.SchemeBuilder.AddToScheme,
		v1alpha1guardduty.SchemeBuilder.AddToScheme,
		v1alpha1iam.SchemeBuilder.AddToScheme,
		v1alpha2iam.SchemeBuilder.AddToScheme,
		v1alpha1imagebuilder.SchemeBuilder.AddToScheme,
		v1alpha1inspector.SchemeBuilder.AddToScheme,
		v1alpha1iot.SchemeBuilder.AddToScheme,
		v1alpha1kafka.SchemeBuilder.AddToScheme,
		v1alpha1kinesis.SchemeBuilder.AddToScheme,
		v1alpha1kinesisanalytics.SchemeBuilder.AddToScheme,
		v1alpha1kinesisanalyticsv2.SchemeBuilder.AddToScheme,
		v1alpha1kinesisvideo.SchemeBuilder.AddToScheme,
		v1alpha1kms.SchemeBuilder.AddToScheme,
		v1alpha2kms.SchemeBuilder.AddToScheme,
		v1alpha1lakeformation.SchemeBuilder.AddToScheme,
		v1alpha1lambda.SchemeBuilder.AddToScheme,
		v1alpha1lexmodels.SchemeBuilder.AddToScheme,
		v1alpha1licensemanager.SchemeBuilder.AddToScheme,
		v1alpha1lightsail.SchemeBuilder.AddToScheme,
		v1alpha1macie.SchemeBuilder.AddToScheme,
		v1alpha1macie2.SchemeBuilder.AddToScheme,
		v1alpha1mediaconvert.SchemeBuilder.AddToScheme,
		v1alpha1mediapackage.SchemeBuilder.AddToScheme,
		v1alpha1mediastore.SchemeBuilder.AddToScheme,
		v1alpha2mq.SchemeBuilder.AddToScheme,
		v1alpha2neptune.SchemeBuilder.AddToScheme,
		v1alpha1networkfirewall.SchemeBuilder.AddToScheme,
		v1alpha1opsworks.SchemeBuilder.AddToScheme,
		v1alpha1organizations.SchemeBuilder.AddToScheme,
		v1alpha1pinpoint.SchemeBuilder.AddToScheme,
		v1alpha1qldb.SchemeBuilder.AddToScheme,
		v1alpha1quicksight.SchemeBuilder.AddToScheme,
		v1alpha1ram.SchemeBuilder.AddToScheme,
		v1alpha1rds.SchemeBuilder.AddToScheme,
		v1alpha2rds.SchemeBuilder.AddToScheme,
		v1alpha1redshift.SchemeBuilder.AddToScheme,
		v1alpha1resourcegroups.SchemeBuilder.AddToScheme,
		v1alpha2route53.SchemeBuilder.AddToScheme,
		v1alpha1route53recoverycontrolconfig.SchemeBuilder.AddToScheme,
		v1alpha1route53recoveryreadiness.SchemeBuilder.AddToScheme,
		v1alpha1route53resolver.SchemeBuilder.AddToScheme,
		v1alpha1s3.SchemeBuilder.AddToScheme,
		v1alpha2s3.SchemeBuilder.AddToScheme,
		v1alpha1s3control.SchemeBuilder.AddToScheme,
		v1alpha1s3outposts.SchemeBuilder.AddToScheme,
		v1alpha1sagemaker.SchemeBuilder.AddToScheme,
		v1alpha1schemas.SchemeBuilder.AddToScheme,
		v1alpha1secretsmanager.SchemeBuilder.AddToScheme,
		v1alpha1securityhub.SchemeBuilder.AddToScheme,
		v1alpha1serverlessrepo.SchemeBuilder.AddToScheme,
		v1alpha1servicecatalog.SchemeBuilder.AddToScheme,
		v1alpha1servicediscovery.SchemeBuilder.AddToScheme,
		v1alpha1servicequotas.SchemeBuilder.AddToScheme,
		v1alpha1ses.SchemeBuilder.AddToScheme,
		v1alpha1sfn.SchemeBuilder.AddToScheme,
		v1alpha1shield.SchemeBuilder.AddToScheme,
		v1alpha1signer.SchemeBuilder.AddToScheme,
		v1alpha1simpledb.SchemeBuilder.AddToScheme,
		v1alpha1sns.SchemeBuilder.AddToScheme,
		v1alpha1sqs.SchemeBuilder.AddToScheme,
		v1alpha1ssm.SchemeBuilder.AddToScheme,
		v1alpha1ssoadmin.SchemeBuilder.AddToScheme,
		v1alpha1storagegateway.SchemeBuilder.AddToScheme,
		v1alpha1swf.SchemeBuilder.AddToScheme,
		v1alpha1synthetics.SchemeBuilder.AddToScheme,
		v1alpha1timestreamwrite.SchemeBuilder.AddToScheme,
		v1alpha1transfer.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1alpha1waf.SchemeBuilder.AddToScheme,
		v1alpha1wafregional.SchemeBuilder.AddToScheme,
		v1alpha1wafv2.SchemeBuilder.AddToScheme,
		v1alpha1worklink.SchemeBuilder.AddToScheme,
		v1alpha1workspaces.SchemeBuilder.AddToScheme,
		v1alpha1xray.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
