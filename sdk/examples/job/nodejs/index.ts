import * as pulumi from "@pulumi/pulumi";
import * as ketch from "@shipa-corp/kpulumi";

const item = new ketch.Job("js-job-1", {
    job: {
        name: "js-job-1",
        framework: "dev",
        version: "v1",
        description: "Pulumi NodeJs job",
        policy: {
            restartPolicy: "Never"
        },
        containers: [
            {
                name: "pi",
                image: "perl",
                commands: [
                    "perl",
                    "-Mbignum=bpi",
                    "wle",
                    "print bpi(2000)"
                ],
            }
        ],
    }
});

export const jobName = item.job.name;
