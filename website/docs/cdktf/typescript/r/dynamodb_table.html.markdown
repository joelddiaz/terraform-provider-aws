---
subcategory: "DynamoDB"
layout: "aws"
page_title: "AWS: aws_dynamodb_table"
description: |-
  Provides a DynamoDB table resource
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_dynamodb_table

Provides a DynamoDB table resource.

~> **Note:** We recommend using `lifecycle` [`ignoreChanges`](https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html#ignore_changes) for `readCapacity` and/or `writeCapacity` if there's [autoscaling policy](/docs/providers/aws/r/appautoscaling_policy.html) attached to the table.

~> **Note:** When using [aws_dynamodb_table_replica](/docs/providers/aws/r/dynamodb_table_replica.html) with this resource, use `lifecycle` [`ignoreChanges`](https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html#ignore_changes) for `replica`, _e.g._, `lifecycle { ignore_changes = [replica] }`.

## DynamoDB Table attributes

Only define attributes on the table object that are going to be used as:

* Table hash key or range key
* LSI or GSI hash key or range key

The DynamoDB API expects attribute structure (name and type) to be passed along when creating or updating GSI/LSIs or creating the initial table. In these cases it expects the Hash / Range keys to be provided. Because these get re-used in numerous places (i.e the table's range key could be a part of one or more GSIs), they are stored on the table object to prevent duplication and increase consistency. If you add attributes here that are not used in these scenarios it can cause an infinite loop in planning.

## Example Usage

### Basic Example

The following dynamodb table description models the table and GSI shown in the [AWS SDK example documentation](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.html)

```typescript
// Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DynamodbTable } from "./.gen/providers/aws/dynamodb-table";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DynamodbTable(this, "basic-dynamodb-table", {
      attribute: [
        {
          name: "UserId",
          type: "S",
        },
        {
          name: "GameTitle",
          type: "S",
        },
        {
          name: "TopScore",
          type: "N",
        },
      ],
      billingMode: "PROVISIONED",
      globalSecondaryIndex: [
        {
          hashKey: "GameTitle",
          name: "GameTitleIndex",
          nonKeyAttributes: ["UserId"],
          projectionType: "INCLUDE",
          rangeKey: "TopScore",
          readCapacity: 10,
          writeCapacity: 10,
        },
      ],
      hashKey: "UserId",
      name: "GameScores",
      rangeKey: "GameTitle",
      readCapacity: 20,
      tags: {
        Environment: "production",
        Name: "dynamodb-table-1",
      },
      ttl: {
        attributeName: "TimeToExist",
        enabled: false,
      },
      writeCapacity: 20,
    });
  }
}

```

### Global Tables

This resource implements support for [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) via `replica` configuration blocks. For working with [DynamoDB Global Tables V1 (version 2017.11.29)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V1.html), see the [`awsDynamodbGlobalTable` resource](/docs/providers/aws/r/dynamodb_global_table.html).

~> **Note:** [aws_dynamodb_table_replica](/docs/providers/aws/r/dynamodb_table_replica.html) is an alternate way of configuring Global Tables. Do not use `replica` configuration blocks of `awsDynamodbTable` together with [aws_dynamodb_table_replica](/docs/providers/aws/r/dynamodb_table_replica.html).

```typescript
// Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DynamodbTable } from "./.gen/providers/aws/dynamodb-table";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DynamodbTable(this, "example", {
      attribute: [
        {
          name: "TestTableHashKey",
          type: "S",
        },
      ],
      billingMode: "PAY_PER_REQUEST",
      hashKey: "TestTableHashKey",
      name: "example",
      replica: [
        {
          regionName: "us-east-2",
        },
        {
          regionName: "us-west-2",
        },
      ],
      streamEnabled: true,
      streamViewType: "NEW_AND_OLD_IMAGES",
    });
  }
}

```

### Replica Tagging

You can manage global table replicas' tags in various ways. This example shows using `replica.*PropagateTags` for the first replica and the `awsDynamodbTag` resource for the other.

```typescript
// Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, Fn, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsRegion } from "./.gen/providers/aws/data-aws-region";
import { DynamodbTable } from "./.gen/providers/aws/dynamodb-table";
import { DynamodbTag } from "./.gen/providers/aws/dynamodb-tag";
import { AwsProvider } from "./.gen/providers/aws/provider";
import { AwsalternateProvider } from "./.gen/providers/awsalternate/provider";
import { AwsthirdProvider } from "./.gen/providers/awsthird/provider";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    /*The following providers are missing schema information and might need manual adjustments to synthesize correctly: awsalternate, awsthird.
    For a more precise conversion please use the --provider flag in convert.*/
    new AwsProvider(this, "aws", {
      region: "us-west-2",
    });
    new AwsalternateProvider(this, "awsalternate", {
      region: "us-east-1",
    });
    new AwsthirdProvider(this, "awsthird", {
      region: "us-east-2",
    });
    const alternate = new DataAwsRegion(this, "alternate", {
      provider: "awsalternate",
    });
    const current = new DataAwsRegion(this, "current", {});
    const third = new DataAwsRegion(this, "third", {
      provider: "awsthird",
    });
    const example = new DynamodbTable(this, "example", {
      attribute: [
        {
          name: "TestTableHashKey",
          type: "S",
        },
      ],
      billingMode: "PAY_PER_REQUEST",
      hashKey: "TestTableHashKey",
      name: "example-13281",
      replica: [
        {
          regionName: Token.asString(alternate.name),
        },
        {
          propagateTags: true,
          regionName: Token.asString(third.name),
        },
      ],
      streamEnabled: true,
      streamViewType: "NEW_AND_OLD_IMAGES",
      tags: {
        Architect: "Eleanor",
        Zone: "SW",
      },
    });
    const awsDynamodbTagExample = new DynamodbTag(this, "example_7", {
      key: "Architect",
      resourceArn: Token.asString(
        Fn.replace(
          example.arn,
          Token.asString(current.name),
          Token.asString(alternate.name)
        )
      ),
      value: "Gigi",
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsDynamodbTagExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

Required arguments:

* `attribute` - (Required) Set of nested attribute definitions. Only required for `hashKey` and `rangeKey` attributes. See below.
* `hashKey` - (Required, Forces new resource) Attribute to use as the hash (partition) key. Must also be defined as an `attribute`. See below.
* `name` - (Required) Unique within a region name of the table.

Optional arguments:

* `billingMode` - (Optional) Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `provisioned` and `payPerRequest`. Defaults to `provisioned`.
* `deletionProtectionEnabled` - (Optional) Enables deletion protection for table. Defaults to `false`.
* `globalSecondaryIndex` - (Optional) Describe a GSI for the table; subject to the normal limits on the number of GSIs, projected attributes, etc. See below.
* `localSecondaryIndex` - (Optional, Forces new resource) Describe an LSI on the table; these can only be allocated _at creation_ so you cannot change this definition after you have created the resource. See below.
* `pointInTimeRecovery` - (Optional) Enable point-in-time recovery options. See below.
* `rangeKey` - (Optional, Forces new resource) Attribute to use as the range (sort) key. Must also be defined as an `attribute`, see below.
* `readCapacity` - (Optional) Number of read units for this table. If the `billingMode` is `provisioned`, this field is required.
* `replica` - (Optional) Configuration block(s) with [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) replication configurations. See below.
* `restoreDateTime` - (Optional) Time of the point-in-time recovery point to restore.
* `restoreSourceName` - (Optional) Name of the table to restore. Must match the name of an existing table.
* `restoreToLatestTime` - (Optional) If set, restores table to the most recent point-in-time recovery point.
* `serverSideEncryption` - (Optional) Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS-owned Customer Master Key if this argument isn't specified. See below.
* `streamEnabled` - (Optional) Whether Streams are enabled.
* `streamViewType` - (Optional) When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `keysOnly`, `newImage`, `oldImage`, `newAndOldImages`.
* `tableClass` - (Optional) Storage class of the table.
  Valid values are `standard` and `standardInfrequentAccess`.
  Default value is `standard`.
* `tags` - (Optional) A map of tags to populate on the created table. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `ttl` - (Optional) Configuration block for TTL. See below.
* `writeCapacity` - (Optional) Number of write units for this table. If the `billingMode` is `provisioned`, this field is required.

### `attribute`

* `name` - (Required) Name of the attribute
* `type` - (Required) Attribute type. Valid values are `s` (string), `n` (number), `b` (binary).

#### `globalSecondaryIndex`

* `hashKey` - (Required) Name of the hash key in the index; must be defined as an attribute in the resource.
* `name` - (Required) Name of the index.
* `nonKeyAttributes` - (Optional) Only required with `include` as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
* `projectionType` - (Required) One of `all`, `include` or `keysOnly` where `all` projects every attribute into the index, `keysOnly` projects  into the index only the table and index hash_key and sort_key attributes ,  `include` projects into the index all of the attributes that are defined in `nonKeyAttributes` in addition to the attributes that that`keysOnly` project.
* `rangeKey` - (Optional) Name of the range key; must be defined
* `readCapacity` - (Optional) Number of read units for this index. Must be set if billing_mode is set to PROVISIONED.
* `writeCapacity` - (Optional) Number of write units for this index. Must be set if billing_mode is set to PROVISIONED.

### `localSecondaryIndex`

* `name` - (Required) Name of the index
* `nonKeyAttributes` - (Optional) Only required with `include` as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
* `projectionType` - (Required) One of `all`, `include` or `keysOnly` where `all` projects every attribute into the index, `keysOnly` projects  into the index only the table and index hash_key and sort_key attributes ,  `include` projects into the index all of the attributes that are defined in `nonKeyAttributes` in addition to the attributes that that`keysOnly` project.
* `rangeKey` - (Required) Name of the range key.

### `pointInTimeRecovery`

* `enabled` - (Required) Whether to enable point-in-time recovery. It can take 10 minutes to enable for new tables. If the `pointInTimeRecovery` block is not provided, this defaults to `false`.

### `replica`

* `kmsKeyArn` - (Optional, Forces new resource) ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, `alias/aws/dynamodb`. **Note:** This attribute will _not_ be populated with the ARN of _default_ keys.
* `pointInTimeRecovery` - (Optional) Whether to enable Point In Time Recovery for the replica. Default is `false`.
* `propagateTags` - (Optional) Whether to propagate the global table's tags to a replica. Default is `false`. Changes to tags only move in one direction: from global (source) to replica. In other words, tag drift on a replica will not trigger an update. Tag or replica changes on the global table, whether from drift or configuration changes, are propagated to replicas. Changing from `true` to `false` on a subsequent `apply` means replica tags are left as they were, unmanaged, not deleted.
* `regionName` - (Required) Region name of the replica.

### `serverSideEncryption`

* `enabled` - (Required) Whether or not to enable encryption at rest using an AWS managed KMS customer master key (CMK). If `enabled` is `false` then server-side encryption is set to AWS-_owned_ key (shown as `default` in the AWS console). Potentially confusingly, if `enabled` is `true` and no `kmsKeyArn` is specified then server-side encryption is set to the _default_ KMS-_managed_ key (shown as `kms` in the AWS console). The [AWS KMS documentation](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html) explains the difference between AWS-_owned_ and KMS-_managed_ keys.
* `kmsKeyArn` - (Optional) ARN of the CMK that should be used for the AWS KMS encryption. This argument should only be used if the key is different from the default KMS-managed DynamoDB key, `alias/aws/dynamodb`. **Note:** This attribute will _not_ be populated with the ARN of _default_ keys.

### `ttl`

* `enabled` - (Required) Whether TTL is enabled.
* `attributeName` - (Required) Name of the table attribute to store the TTL timestamp in.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - ARN of the table
* `id` - Name of the table
* `replica.*Arn` - ARN of the replica
* `replica.*StreamArn` - ARN of the replica Table Stream. Only available when `stream_enabled = true`.
* `replica.*StreamLabel` - Timestamp, in ISO 8601 format, for the replica stream. Note that this timestamp is not a unique identifier for the stream on its own. However, the combination of AWS customer ID, table name and this field is guaranteed to be unique. It can be used for creating CloudWatch Alarms. Only available when `stream_enabled = true`.
* `streamArn` - ARN of the Table Stream. Only available when `stream_enabled = true`
* `streamLabel` - Timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not a unique identifier for the stream on its own. However, the combination of AWS customer ID, table name and this field is guaranteed to be unique. It can be used for creating CloudWatch Alarms. Only available when `stream_enabled = true`.
* `tagsAll` - Map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

~> **Note:** There are a variety of default timeouts set internally. If you set a shorter custom timeout than one of the defaults, the custom timeout will not be respected as the longer of the custom or internal default will be used.

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30M`)
* `update` - (Default `60M`)
* `delete` - (Default `10M`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import DynamoDB tables using the `name`. For example:

```typescript
// Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
  }
}

```

Using `terraform import`, import DynamoDB tables using the `name`. For example:

```console
% terraform import aws_dynamodb_table.basic-dynamodb-table GameScores
```

<!-- cache-key: cdktf-0.17.1 input-e3925cbe8e971022eb782ebe4015491c74d9d9263b8511d0e7f093fb1e1c7b8f -->