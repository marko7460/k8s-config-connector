// Copyright 2022 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// GENERATED BY gen_go_data.go
// gen_go_data -package storage -var YAML_bucket blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/storage/bucket.yaml

package storage

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/storage/bucket.yaml
var YAML_bucket = []byte("info:\n  title: Storage/Bucket\n  description: The Storage Bucket resource\n  x-dcl-struct-name: Bucket\n  x-dcl-has-iam: true\npaths:\n  get:\n    description: The function used to get information about a Bucket\n    parameters:\n    - name: bucket\n      required: true\n      description: A full instance of a Bucket\n  apply:\n    description: The function used to apply information about a Bucket\n    parameters:\n    - name: bucket\n      required: true\n      description: A full instance of a Bucket\n  delete:\n    description: The function used to delete a Bucket\n    parameters:\n    - name: bucket\n      required: true\n      description: A full instance of a Bucket\n  deleteAll:\n    description: The function used to delete all Bucket\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Bucket\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Bucket:\n      title: Bucket\n      x-dcl-id: b/{{name}}?userProject={{project}}\n      x-dcl-parent-container: project\n      x-dcl-has-create: true\n      x-dcl-has-iam: true\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - project\n      - location\n      - name\n      properties:\n        cors:\n          type: array\n          x-dcl-go-name: Cors\n          description: 'The bucket''s Cross-Origin Resource Sharing (CORS) configuration. '\n          x-dcl-send-empty: true\n          x-dcl-list-type: list\n          items:\n            type: object\n            x-dcl-go-type: BucketCors\n            properties:\n              maxAgeSeconds:\n                type: integer\n                format: int64\n                x-dcl-go-name: MaxAgeSeconds\n                description: 'The value, in seconds, to return in the Access-Control-Max-Age\n                  header used in preflight responses. '\n              method:\n                type: array\n                x-dcl-go-name: Method\n                description: 'The list of HTTP methods on which to include CORS response\n                  headers, (GET, OPTIONS, POST, etc) Note: \"*\" is permitted in the\n                  list of methods, and means \"any method\". '\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: string\n                  x-dcl-go-type: string\n              origin:\n                type: array\n                x-dcl-go-name: Origin\n                description: 'The list of Origins eligible to receive CORS response\n                  headers. Note: \"*\" is permitted in the list of origins, and means\n                  \"any Origin\". '\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: string\n                  x-dcl-go-type: string\n              responseHeader:\n                type: array\n                x-dcl-go-name: ResponseHeader\n                description: 'The list of HTTP headers other than the simple response\n                  headers to give permission for the user-agent to share across domains. '\n                x-dcl-send-empty: true\n                x-dcl-list-type: list\n                items:\n                  type: string\n                  x-dcl-go-type: string\n        lifecycle:\n          type: object\n          x-dcl-go-name: Lifecycle\n          x-dcl-go-type: BucketLifecycle\n          description: 'The bucket''s lifecycle configuration.  See https://developers.google.com/storage/docs/lifecycle\n            for more information. '\n          properties:\n            rule:\n              type: array\n              x-dcl-go-name: Rule\n              description: 'A lifecycle management rule, which is made of an action\n                to take and the condition(s) under which the action will be taken. '\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: BucketLifecycleRule\n                properties:\n                  action:\n                    type: object\n                    x-dcl-go-name: Action\n                    x-dcl-go-type: BucketLifecycleRuleAction\n                    description: The action to take.\n                    properties:\n                      storageClass:\n                        type: string\n                        x-dcl-go-name: StorageClass\n                        description: 'Target storage class. Required if the type of\n                          the action is SetStorageClass. '\n                      type:\n                        type: string\n                        x-dcl-go-name: Type\n                        x-dcl-go-type: BucketLifecycleRuleActionTypeEnum\n                        description: 'Type of the action. Currently, only Delete and\n                          SetStorageClass are supported. '\n                        enum:\n                        - Delete\n                        - SetStorageClass\n                  condition:\n                    type: object\n                    x-dcl-go-name: Condition\n                    x-dcl-go-type: BucketLifecycleRuleCondition\n                    description: 'The condition(s) under which the action will be\n                      taken. '\n                    properties:\n                      age:\n                        type: integer\n                        format: int64\n                        x-dcl-go-name: Age\n                        description: 'Age of an object (in days). This condition is\n                          satisfied when an object reaches the specified age. '\n                      createdBefore:\n                        type: string\n                        format: date-time\n                        x-dcl-go-name: CreatedBefore\n                        description: 'A date in RFC 3339 format with only the date\n                          part (for instance, \"2013-01-15\"). This condition is satisfied\n                          when an object is created before midnight of the specified\n                          date in UTC. '\n                      matchesStorageClass:\n                        type: array\n                        x-dcl-go-name: MatchesStorageClass\n                        description: 'Objects having any of the storage classes specified\n                          by this condition will be matched. Values include MULTI_REGIONAL,\n                          REGIONAL, NEARLINE, COLDLINE, STANDARD, and DURABLE_REDUCED_AVAILABILITY. '\n                        x-dcl-send-empty: true\n                        x-dcl-list-type: list\n                        items:\n                          type: string\n                          x-dcl-go-type: string\n                      numNewerVersions:\n                        type: integer\n                        format: int64\n                        x-dcl-go-name: NumNewerVersions\n                        description: 'Relevant only for versioned objects. If the\n                          value is N, this condition is satisfied when there are at\n                          least N versions (including the live version) newer than\n                          this version of the object. '\n                      withState:\n                        type: string\n                        x-dcl-go-name: WithState\n                        x-dcl-go-type: BucketLifecycleRuleConditionWithStateEnum\n                        description: 'Match to live and/or archived objects. Unversioned\n                          buckets have only live objects. Supported values include:\n                          ''LIVE'', ''ARCHIVED'', ''ANY''.'\n                        enum:\n                        - LIVE\n                        - ARCHIVED\n                        - ANY\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: 'The location of the bucket. Object data for objects in the\n            bucket resides in physical storage within this region. Defaults to `US`. '\n          x-kubernetes-immutable: true\n        logging:\n          type: object\n          x-dcl-go-name: Logging\n          x-dcl-go-type: BucketLogging\n          description: 'The bucket''s logging configuration, which defines the destination\n            bucket and optional name prefix for the current bucket''s logs. '\n          properties:\n            logBucket:\n              type: string\n              x-dcl-go-name: LogBucket\n              description: 'The destination bucket where the current bucket''s logs\n                should be placed. '\n            logObjectPrefix:\n              type: string\n              x-dcl-go-name: LogObjectPrefix\n              description: The object prefix for log objects. If it's not provided,\n                it defaults to the bucket's name.\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: 'The name of the bucket. '\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project id of the resource.\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        storageClass:\n          type: string\n          x-dcl-go-name: StorageClass\n          x-dcl-go-type: BucketStorageClassEnum\n          description: 'The bucket''s default storage class, used whenever no storageClass\n            is specified for a newly-created object. This defines how objects in the\n            bucket are stored and determines the SLA and the cost of storage. Values\n            include MULTI_REGIONAL, REGIONAL, STANDARD, NEARLINE, COLDLINE, ARCHIVE,\n            and DURABLE_REDUCED_AVAILABILITY. If this value is not specified when\n            the bucket is created, it will default to STANDARD. For more information,\n            see storage classes. '\n          enum:\n          - MULTI_REGIONAL\n          - REGIONAL\n          - STANDARD\n          - NEARLINE\n          - COLDLINE\n          - ARCHIVE\n          - DURABLE_REDUCED_AVAILABILITY\n        versioning:\n          type: object\n          x-dcl-go-name: Versioning\n          x-dcl-go-type: BucketVersioning\n          description: The bucket's versioning configuration.\n          properties:\n            enabled:\n              type: boolean\n              x-dcl-go-name: Enabled\n              description: 'While set to true, versioning is fully enabled for this\n                bucket. '\n        website:\n          type: object\n          x-dcl-go-name: Website\n          x-dcl-go-type: BucketWebsite\n          description: 'The bucket''s website configuration, controlling how the service\n            behaves when accessing bucket contents as a web site. See the Static Website\n            Examples for more information. '\n          properties:\n            mainPageSuffix:\n              type: string\n              x-dcl-go-name: MainPageSuffix\n              description: 'If the requested object path is missing, the service will\n                ensure the path has a trailing ''/'', append this suffix, and attempt\n                to retrieve the resulting object. This allows the creation of index.html\n                objects to represent directory pages. '\n            notFoundPage:\n              type: string\n              x-dcl-go-name: NotFoundPage\n              description: 'If the requested object path is missing, and any mainPageSuffix\n                object is missing, if applicable, the service will return the named\n                object from this bucket as the content for a 404 Not Found result. '\n")

// 11727 bytes
// MD5: a16310ceb8d340ad0e74595ad738b1bd
