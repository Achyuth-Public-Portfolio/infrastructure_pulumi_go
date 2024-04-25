package eks

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi-eks/sdk/go/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateEKSCluster(ctx *pulumi.Context, role *iam.Role) (*eks.Cluster, error) {

	instanceProfile, err := iam.NewInstanceProfile(ctx, "Instance-profile-1", &iam.InstanceProfileArgs{
		Role: role.Name,
	})


	cluster, err := eks.NewCluster(ctx, "EKS Cluster", &eks.ClusterArgs{
		SkipDefaultNodeGroup: pulumi.BoolRef(true),
		InstanceRoles: nil,
	})

	eks.NewNodeGroup(ctx, "Fixed Node Group", &eks.NodeGroupArgs{
		Cluster: cluster,
		InstanceType:    pulumi.String("t2.nano"),
		DesiredCapacity: pulumi.Int(2),
		MinSize:         pulumi.Int(1),
		MaxSize:         pulumi.Int(3),
		SpotPrice:       pulumi.String("1"),
		Labels: map[string]string{
			"ondemand": "true",
		},
		InstanceProfile: instanceProfile,
	})



	return cluster, err
}