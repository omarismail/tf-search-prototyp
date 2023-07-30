# Terraform Search (tfs)

`tfs` is a command-line tool that integrates with the AWS Resource Explorer to find and list AWS resources matching given criteria, such as tags and region. The goal is to assist in identifying resources that are part of your Terraform configuration or any other infrastructure-as-code (IaC) setup.

## Setup

Before using `tfs`, make sure you have set up the following environment variables:

```bash
export AWS_ACCESS_KEY_ID=<your-access-key>
export AWS_SECRET_ACCESS_KEY=<your-secret-key>
export AWS_REGION=<your-region>
```

Install the binary:

```
make get
make install
```

## AWS Resource Explorer

This tool makes use of the AWS Resource Explorer service to search for resources. You need to have this service enabled for your account. Follow the steps below to enable it:

1. Sign in to the AWS Management Console.
2. Open the service menu by clicking on the "Services" at the top of the console.
3. In the search box, type "Resource Explorer" and select it from the list.
4. If it's not already turned on, you will see an "Enable Resource Explorer" button. Click on it to enable the service.

For more details, refer to the official [AWS Resource Explorer documentation](https://docs.aws.amazon.com/).

## Running the Tool

To use `tfs`, you need to pass the search name to the command like so:

```bash
tfs search <search-name>
```

The `<search-name>` corresponds to a search block defined in your HCL configuration file.

## Sample HCL Configuration

Here's a simple example of what the HCL configuration (`main.hcl`) might look like for `tfs`:

```hcl
provider "aws" {
  access_key = "YOUR_ACCESS_KEY"
  secret_key = "YOUR_SECRET_KEY"
  region     = "us-east-1"
}

search "aws_vpc_foo" {
  provider = "aws"
  
  query "explorer" {
    region   = "us-east-1"
    tags     = ["env:prod"]
  }
}
```

In the above example, a `search` block is defined with the name `aws_vpc_foo` that uses the `aws` provider. Inside the `search` block, a `query` block is defined with the `explorer` type. The `query` block includes the `category`, `region`, and `tags` parameters. 

The `tags` parameter is a list of tags in the form `key:value`, and the tool will search for resources that match all of these tags. In this case, it will search for resources with the tag `env:prod`.

