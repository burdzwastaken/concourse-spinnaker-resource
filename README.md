# Concourse Spinnaker Resource

*Deprecated*: Please use the [pivotal resource](https://github.com/pivotal-cf/spinnaker-resource) which while based off this resource has active development. 

A [Concourse](https://concourse.ci/) resource that allows jobs to trigger [Spinnaker](https://spinnaker.io/) pipelines. 

## Source Configuration
### Required

* `spinnaker_api`: 
* `x509_cert`:
* `x509_key`:
* `spinnaker_application`: The Spinnaker application you would like to trigger.
* `spinnaker_pipeline`: The Spinnaker pipeline you would like to trigger.

## Behaviour

### `check`, `in`

Currently this resource only supports the put phase of a job plan, so these are effectively no-ops. This will change in the future.

### `out`: Triggers a pipeline

Triggers a Spinnaker pipeline.

#### Parameters
**NOTE:** Any [metadata](http://concourse.ci/implementing-resources.html#resource-metadata) in the parameters will be evaluated prior to triggering the pipeline. 

##### Optional:
* `trigger_params`: build information to send to Spinnaker pipeline execution which can be consumed by the [pipeline expressions](https://www.spinnaker.io/guides/user/pipeline-expressions/). can be any key/value pair that you would like to consume. 

## Example Pipeline

```yml
---
resource_types:
- name: spinnaker
  type: docker-image
  source:
    repository: burdz/concourse-spinnaker-resource:0.1

resources:
  - name: spinnaker
    type: spinnaker
    source:
      spinnaker_api: ((spinnaker-api))
      spinnaker_application: samplespinnakerapp
      spinnaker_pipeline: samplespinnakerapp
      spinnaker_x509_cert: ((spinnaker-x509-cert))
      spinnaker_x509_key: ((spinnaker-x509-key))

jobs:
- name: trigger-pipeline
  plan:
  - put: spinnaker
    params:
      trigger_params:
        build_id: (build ${BUILD_ID})
```

**NOTE**: this changed in `burdz/concourse-spinnaker-resource:0.1`
