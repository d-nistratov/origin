package policy

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/spf13/cobra"

	kcmdutil "github.com/GoogleCloudPlatform/kubernetes/pkg/kubectl/cmd/util"

	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
	"github.com/openshift/origin/pkg/client"
	"github.com/openshift/origin/pkg/cmd/util/clientcmd"
)

const WhoCanRecommendedName = "who-can"

type whoCanOptions struct {
	bindingNamespace string
	client           client.Interface

	verb     string
	resource string
}

func NewCmdWhoCan(name, fullName string, f *clientcmd.Factory, out io.Writer) *cobra.Command {
	options := &whoCanOptions{}

	cmd := &cobra.Command{
		Use:   name + " <verb> <resource>",
		Short: "Indicates which users can perform the action on the resource.",
		Long:  `Indicates which users can perform the action on the resource.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := options.complete(args); err != nil {
				kcmdutil.CheckErr(kcmdutil.UsageError(cmd, err.Error()))
			}

			var err error
			if options.client, _, err = f.Clients(); err != nil {
				kcmdutil.CheckErr(err)
			}
			if options.bindingNamespace, err = f.DefaultNamespace(); err != nil {
				kcmdutil.CheckErr(err)
			}
			if err := options.run(); err != nil {
				kcmdutil.CheckErr(err)
			}
		},
	}

	return cmd
}

func (o *whoCanOptions) complete(args []string) error {
	if len(args) != 2 {
		return errors.New("You must specify two arguments: <verb> <resource>")
	}

	o.verb = args[0]
	o.resource = args[1]
	return nil
}

func (o *whoCanOptions) run() error {
	resourceAccessReview := &authorizationapi.ResourceAccessReview{}
	resourceAccessReview.Resource = o.resource
	resourceAccessReview.Verb = o.verb

	resourceAccessReviewResponse, err := o.client.ResourceAccessReviews(o.bindingNamespace).Create(resourceAccessReview)
	if err != nil {
		return err
	}

	if len(resourceAccessReviewResponse.Users) == 0 {
		fmt.Printf("Users:  none\n\n")
	} else {
		fmt.Printf("Users:  %s\n\n", strings.Join(resourceAccessReviewResponse.Users.List(), "\n        "))
	}

	if len(resourceAccessReviewResponse.Groups) == 0 {
		fmt.Printf("Groups: none\n\n")
	} else {
		fmt.Printf("Groups: %s\n\n", strings.Join(resourceAccessReviewResponse.Groups.List(), "\n        "))
	}

	return nil
}
