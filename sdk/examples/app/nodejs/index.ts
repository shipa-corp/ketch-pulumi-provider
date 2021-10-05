import * as pulumi from "@pulumi/pulumi";
import * as ketch from "@shipa-corp/kpulumi";

const framework = new ketch.Framework("f1", {
    framework: {
        name: "f1",
        ingressController: {
            className: "istio",
            serviceEndpoint: "1.2.3.4",
            type: "istio",
        }
    }
});

const app = new ketch.App("a1", {
    app: {
        name: "a1",
        image: "docker.io/shipasoftware/bulletinboard:1.0",
        framework: framework.framework.name,
        cnames: ["a1.ketch.io", "a2.ketch.io"],
        ports: [8080, 8081],
        units: 1,
        processes: [
            {
                name: "web",
                cmds: ["docker-entrypoint.sh", "npm", "start"],
            }
        ],
        routingSettings: {
            weight: 100,
        }
    }
});

export const appName = app.app.name;
export const frameworkName = framework.framework.name;
