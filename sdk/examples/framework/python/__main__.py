# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import pulumi
import pulumi_ketch as ketch

item = ketch.Framework(
    "python-framework-1",
    ketch.FrameworkArgs(
        framework=ketch.FrameworkFrameworkArgs(
            name="py-pulumi-framework-1",
            ingress_controller=ketch.FrameworkFrameworkIngressControllerArgs(
                class_name="istio",
                service_endpoint="1.2.3.4",
                type="istio"
            )
        )
    )
)

pulumi.export("framework-name", item.framework.name)
