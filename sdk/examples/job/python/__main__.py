# Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

import pulumi
import pulumi_ketch as ketch

item = ketch.Job(
    "py-job-1",
    ketch.JobArgs(
        job=ketch.JobJobArgs(
            name="py-job-1",
            version="v1",
            type="Job",
            framework="t-f3",
            description="Pulumi Python job",
            policy=ketch.JobJobPolicyArgs(
                restart_policy="Never"
            ),
            containers=[
                ketch.JobJobContainerArgs(
                    name="pi",
                    image="perl",
                    commands=[
                        "perl",
                        "-Mbignum=bpi",
                        "wle",
                        "print bpi(2000)"
                    ]
                )
            ]
        )
    )
)

pulumi.export("job-name", item.job.name)
