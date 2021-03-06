package client

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/fields"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"

	authorizationapi "github.com/openshift/origin/pkg/authorization/api"
)

type FakeClusterRoles struct {
	Fake *Fake
}

func (c *FakeClusterRoles) List(label labels.Selector, field fields.Selector) (*authorizationapi.ClusterRoleList, error) {
	obj, err := c.Fake.Invokes(FakeAction{Action: "list-clusterRoles"}, &authorizationapi.ClusterRoleList{})
	return obj.(*authorizationapi.ClusterRoleList), err
}

func (c *FakeClusterRoles) Get(name string) (*authorizationapi.ClusterRole, error) {
	obj, err := c.Fake.Invokes(FakeAction{Action: "get-clusterRole"}, &authorizationapi.ClusterRole{})
	return obj.(*authorizationapi.ClusterRole), err
}

func (c *FakeClusterRoles) Create(role *authorizationapi.ClusterRole) (*authorizationapi.ClusterRole, error) {
	obj, err := c.Fake.Invokes(FakeAction{Action: "create-clusterRole", Value: role}, &authorizationapi.ClusterRole{})
	return obj.(*authorizationapi.ClusterRole), err
}

func (c *FakeClusterRoles) Delete(name string) error {
	c.Fake.Actions = append(c.Fake.Actions, FakeAction{Action: "delete-clusterRole", Value: name})
	return nil
}
