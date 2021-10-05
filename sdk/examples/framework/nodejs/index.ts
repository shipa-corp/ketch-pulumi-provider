import * as pulumi from "@pulumi/pulumi";
import * as ketch from "@shipa-corp/kpulumi";

const item = new ketch.Framework("ketch-framework-1", {
    framework: {
        name: "pulumi-framework-1",
        ingressController: {
            className: "istio",
            serviceEndpoint: "1.2.3.4",
            type: "istio",
        }
    }
});

export const frameworkName = item.framework.name;
