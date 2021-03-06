# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['JobArgs', 'Job']

@pulumi.input_type
class JobArgs:
    def __init__(__self__, *,
                 job: pulumi.Input['JobJobArgs']):
        """
        The set of arguments for constructing a Job resource.
        """
        pulumi.set(__self__, "job", job)

    @property
    @pulumi.getter
    def job(self) -> pulumi.Input['JobJobArgs']:
        return pulumi.get(self, "job")

    @job.setter
    def job(self, value: pulumi.Input['JobJobArgs']):
        pulumi.set(self, "job", value)


@pulumi.input_type
class _JobState:
    def __init__(__self__, *,
                 job: Optional[pulumi.Input['JobJobArgs']] = None):
        """
        Input properties used for looking up and filtering Job resources.
        """
        if job is not None:
            pulumi.set(__self__, "job", job)

    @property
    @pulumi.getter
    def job(self) -> Optional[pulumi.Input['JobJobArgs']]:
        return pulumi.get(self, "job")

    @job.setter
    def job(self, value: Optional[pulumi.Input['JobJobArgs']]):
        pulumi.set(self, "job", value)


class Job(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 job: Optional[pulumi.Input[pulumi.InputType['JobJobArgs']]] = None,
                 __props__=None):
        """
        Create a Job resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: JobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Job resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param JobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(JobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 job: Optional[pulumi.Input[pulumi.InputType['JobJobArgs']]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = JobArgs.__new__(JobArgs)

            if job is None and not opts.urn:
                raise TypeError("Missing required property 'job'")
            __props__.__dict__["job"] = job
        super(Job, __self__).__init__(
            'ketch:index/job:Job',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            job: Optional[pulumi.Input[pulumi.InputType['JobJobArgs']]] = None) -> 'Job':
        """
        Get an existing Job resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _JobState.__new__(_JobState)

        __props__.__dict__["job"] = job
        return Job(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def job(self) -> pulumi.Output['outputs.JobJob']:
        return pulumi.get(self, "job")

