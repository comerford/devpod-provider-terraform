# TERRAFORM Provider for DevPod modified to use ECS rather than EC2

The [original version](https://github.com/loft-sh/devpod-provider-terraform/) of this provider provisions an EC2 instance to run Devpod, this one provisions and uses an ECS cluster instead

## Getting started

The provider is available for auto-installation using 

```sh
devpod provider add comerford/devpod-provider-terraform -o TERRAFORM_PROJECT=$PROJECT_LOCATION -o REGION=$AWS_REGION
devpod provider use comerford/devpod-provider-terraform
```

Follow the on-screen instructions to complete the setup.

Needed variables will be:

- TERRAFORM_PROJECT
- REGION

TERRAFORM_PROJECT points to a git repo or directory where the terraform project
that defines the infra is stored.

In this repo it would point to: `./examples/terraform-aws/`

### Creating your first devpod env with terraform

After the initial setup, just use:

```sh
devpod up .
```

You'll need to wait for the machine and environment setup.

### Notes

With the terraform provider, all the power is in the terraform project. So it
will be there where you will place your defaults for these variables:

- IMAGE_URL
- MEM_ALLOCATION
- CPU_ALLOCATION

The memory and CPU allocations will need to be one of the valid combinations for FARGATE

Note that **stop/start is not supported right now on terraform provider**
So the right thing to do is to handle data saving inside your terraform code
(eg. use external data buckets)

### Customize the ECS Cluster

This provider has the following options:

|    NAME           | REQUIRED |          DESCRIPTION                  |         DEFAULT         |
|-------------------|----------|---------------------------------------|-------------------------|
| MEM_ALLOCATION            | false    | The amount of memory to use.                 | 40  |
| IMAGE_URL                 | false    | The container image to use.                |     |
| LAUNCH_TYPE               | false    | The launch type to use for the cluster             | FARGATE    |
| REGION                    | true     | The AWS region to create the        |     |
|                           |          | VM in. E.g. us-west-1                 |     |
| TERRAFORM_PROJECT         | true     | The path or repo where the            |     |
|                           |          | terraform files are. E.g.             |     |
|                           |          | ./examples/terraform or               |     |
|                           |          | https://github.com/examples/terraform |     |

Options can either be set in `env` or using for example:

```sh
devpod provider set-options -o IMAGE_DISK=my-custom-ami
devpod provider set-options -o INSTANCE_TYPE=t2.micro
devpod provider set-options -o REGION=us-west-2
```
