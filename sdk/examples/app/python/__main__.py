# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import pulumi
import pulumi_ketch as ketch

framework = ketch.Framework(
    "ketch-py-framework-1",
    ketch.FrameworkArgs(
        framework=ketch.FrameworkFrameworkArgs(
            name="ketch-py-framework-1",
            ingress_controller=ketch.FrameworkFrameworkIngressControllerArgs(
                class_name="istio",
                service_endpoint="1.2.3.4",
                type="istio"
            )
        )
    )
)

app = ketch.App(
    "ketch-py-app-1",
    ketch.AppArgs(
        app=ketch.AppAppArgs(
            name="ketch-py-app-1",
            image="docker.io/shipasoftware/bulletinboard:1.0",
            framework=framework.framework.name,
            cnames=["py-app-1.ketch.io", "py-app-2.ketch.io"],
            ports=[8002, 8003],
            units=5,
            processes=[
                ketch.AppAppProcessArgs(
                    name="web",
                    cmds=["docker-entrypoint.sh", "npm", "start"]
                )
            ],
            routing_settings=ketch.AppAppRoutingSettingsArgs(
                weight=100
            ),
            version=1
        )
    )
)

pulumi.export("app-name", app.app.name)
