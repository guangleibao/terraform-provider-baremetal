// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"testing"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	"github.com/oracle/terraform-provider-baremetal/client/mocks"
)

func TestLoadBalancerBackendsDatasource(t *testing.T) {
	client := &mocks.BareMetalClient{}
	providers := map[string]terraform.ResourceProvider{
		"baremetal": Provider(func(d *schema.ResourceData) (interface{}, error) {
			return client, nil
		}),
	}
	resourceName := "data.baremetal_load_balancer_backends.t"
	config := `
data "baremetal_load_balancer_backends" "t" {
  load_balancer_id = "ocid1.loadbalancer.stub_id"
  backendset_name  = "stub_backendset_name"
}
`
	config += testProviderConfig

	loadbalancerID := "ocid1.loadbalancer.stub_id"
	backendsetName := "stub_backendset_name"
	list := &baremetal.ListBackends{
		Backends: []baremetal.Backend{
			{Name: "stub_name1"},
			{Name: "stub_name2"},
		},
	}
	client.On(
		"ListBackends",
		loadbalancerID,
		backendsetName,
	).Return(list, nil)

	resource.UnitTest(t, resource.TestCase{
		PreventPostDestroyRefresh: true,
		Providers:                 providers,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "load_balancer_id", loadbalancerID),
					resource.TestCheckResourceAttr(resourceName, "backends.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "backends.0.name", "stub_name1"),
					resource.TestCheckResourceAttr(resourceName, "backends.1.name", "stub_name2"),
				),
			},
		},
	})
}
