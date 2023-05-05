# KubeFox Platform Services

The KubeFox Platform is driven by three core components: 

<table cellpadding="0">
  <tr style="padding: 0">
    <td valign="center">
      <p>API Server</p>
    </td>
    <td valign="center"><a href="https://goreportcard.com/report/github.com/xigxog/kubefox/components/api-server"><img src="https://goreportcard.com/badge/github.com/xigxog/kubefox/components/api-server"/></a>
    </td>
  </tr>
  <tr style="padding: 0">
    <td valign="center">
      <p>Operator</p>
    </td>
    <td valign="center"><a href="https://goreportcard.com/report/github.com/xigxog/kubefox/components/operator"><img src="https://goreportcard.com/badge/github.com/xigxog/kubefox/components/operator"/></a>
    </td>
  </tr>
  <tr style="padding: 0">
    <td valign="center">
      <p>Runtime Server</p>
    </td>
    <td valign="center"><a href="https://goreportcard.com/report/github.com/xigxog/kubefox/components/runtime-server"><img src="https://goreportcard.com/badge/github.com/xigxog/kubefox/components/runtime-server"/></a>
    </td>
  </tr>
</table>

## API Server

The API Server implements the KubeFox API, a resource-based (RESTful)
programmatic interface provided via HTTP. Several types of primary resources are
defined by the API. Additional subresources are defined for each primary
resource type.

Most KubeFox API resource types are objects: they represent a concrete instance
of a concept, like a config or a system. A smaller number of API resource types
are virtual in that they represent operations on objects, rather than objects
themselves, such as a deployment or release.

Every primary resource includes a metadata subresource. The metadata subresource
contains attributes that apply to all instances (objects) of that resource.

## Ref Subresources

Resources of type config, environment, and system include a special set of
subresources known as refs which allow retrieval of a specific resource instance
(object). Three types of ref subresources exist: id, tag, and branch. The branch
subresource is only included in the system resource. Performing a get on a ref
subresource will return the specific resource instance (object) it references.

An id subresource is automatically created when a resource of type config,
environment, or system is created. The URI of the subresource is generated by
the server and is guaranteed to be unique to that instance of the resource.

A tag subresource can be created by sending the id of the resource instance it
should reference. Tag subresources can be created or deleted, but not updated.
If a tag subresource is being utilized by a deployment or release it cannot be
deleted.

For system resources a branch subresource can be created by sending the id of
the resource instance it should reference. Branch subresources can be updated to
reference a different id at any time.

## API Verbs

API verbs are supported by most primary resources and subresources.

| API Verb | HTTP Verb | Description                                                                                                          |
| -------- | --------- | -------------------------------------------------------------------------------------------------------------------- |
| List     | GET       | Lists available resources of specified type.                                                                         |
| Create   | POST      | Creates resource of specified type. For resources of type object, the object should be provided in the request body. |
| Get      | GET       | Retrieves the latest resource instance and status of specified type and name.                                        |
| Delete   | DELETE    | Deletes resource of specified type and name. Often the object for the resource is also deleted.                      |
| Update   | PUT       | Updates resource of specified type and name. If the resource does not exist it is created.                           |
| Patch    | PATCH     | Patches resource of specified type and name. Can only be performed on existing resources.                            |

### Config

A config contains secrets and component configuration required for a component
to process an event.

Supported refs:

- id
- tag

#### Endpoints

All endpoints are prefixed with the path `api/kubefox/v0/`.

| API Verb | HTTP Verb | Path                      |
| -------- | --------- | ------------------------- |
| List     | GET       | `config`                  |
| Create   | POST      | `config/{name}`           |
| Get      | GET       | `config/{name}`           |
| Delete   | DELETE    | `config/{name}`           |
| Update   | PUT       | `config/{name}/metadata`  |
| Get      | GET       | `config/{name}/metadata`  |
| List     | GET       | `config/{name}/id`        |
| Get      | GET       | `config/{name}/id/{id}`   |
| List     | GET       | `config/{name}/tag`       |
| Update   | PUT       | `config/{name}/tag/{tag}` |
| Delete   | DELETE    | `config/{name}/tag/{tag}` |

### Environment

An environment is a named set of variables. These variables can be accessed by
components during runtime or used in route templates. An environment must
reference a config so that component configuration can be looked up when
processing events assigned to that environment.

Supported refs:

- id
- tag

#### Endpoints

| API Verb | HTTP Verb | Path                           |
| -------- | --------- | ------------------------------ |
| List     | GET       | `environment`                  |
| Create   | POST      | `environment/{name}`           |
| Get      | GET       | `environment/{name}`           |
| Delete   | DELETE    | `environment/{name}`           |
| Update   | PUT       | `environment/{name}/metadata`  |
| Get      | GET       | `environment/{name}/metadata`  |
| List     | GET       | `environment/{name}/id`        |
| Get      | GET       | `environment/{name}/id/{id}`   |
| List     | GET       | `environment/{name}/tag`       |
| Update   | PUT       | `environment/{name}/tag/{tag}` |
| Delete   | DELETE    | `environment/{name}/tag/{tag}` |

### System

A system is a collection of applications, components, and routes. They are
generated from a KubeFox Git repository tree. The id of a system is the hash of
the Git commit for that tree.

Supported refs:

- Id
- Branch
- Tag

#### Endpoints

| API Verb | HTTP Verb | Path                            |
| -------- | --------- | ------------------------------- |
| List     | GET       | `system`                        |
| Create   | POST      | `system/{name}`                 |
| Get      | GET       | `system/{name}`                 |
| Delete   | DELETE    | `system/{name}`                 |
| Update   | PUT       | `system/{name}/metadata`        |
| Get      | GET       | `system/{name}/metadata`        |
| List     | GET       | `system/{name}/id`              |
| Get      | GET       | `system/{name}/id/{id}`         |
| List     | GET       | `system/{name}/branch`          |
| Update   | PUT       | `system/{name}/branch/{branch}` |
| Delete   | DELETE    | `system/{name}/branch/{branch}` |
| List     | GET       | `system/{name}/tag`             |
| Create   | POST      | `system/{name}/tag/{tag}`       |
| Delete   | DELETE    | `system/{name}/tag/{tag}`       |

### Platform

Subresources:

- Deployment
- Release

#### Endpoints

| API Verb | HTTP Verb | Path                                                  |
| -------- | --------- | ----------------------------------------------------- |
| List     | GET       | `platform`                                            |
| Update   | PUT       | `platform/{name}`                                     |
| Update   | PATCH     | `platform/{name}`                                     |
| Get      | GET       | `platform/{name}`                                     |
| Update   | PUT       | `platform/{name}/metadata`                            |
| Get      | GET       | `platform/{name}/metadata`                            |
| Create   | POST      | `platform/{name}/deployment`                          |
| List     | GET       | `platform/{name}/deployment`                          |
| Get      | GET       | `platform/{name}/deployment/{system}/{refType}/{ref}` |
| Delete   | DELETE    | `platform/{name}/deployment/{system}/{refType}/{ref}` |
| Create   | POST      | `platform/{name}/release`                             |
| List     | GET       | `platform/{name}/release`                             |
| Get      | GET       | `platform/{name}/release/{system}/{environment}`      |
| Delete   | DELETE    | `platform/{name}/release/{system}/{environment}`      |

## Local Development

It is possible to start a Kubefox Platform component locally for development.
The follow examples apply to starting the `runtime-server` but can be used to
start any component with modifications to the commands. To start port forward to
vault and nats:

```bash
kubectl port-forward -n kubefox-system service/kfp-vault 8200
kubectl port-forward -n kubefox-system service/kfp-nats 4222
```

Then start the broker and component on your workstation. This can be done using
the `launch.json` in VS Code. First launch the broker using the launch
configuration named `runtime-broker`. Then launch the runtime server using the
`runtime` launch configuration.

Now copy the following file to `./workstation/runtime-srv-service-local.yaml`.
Replace `<YOUR IP ADDRESS>` with the workstation's IP address. This will
redirect requests from components running in Kubernetes to the instance running
locally on the workstation:

```yaml
---
apiVersion: v1
kind: Service
metadata:
  name: kfp-runtime-server
  labels:
    app.kubernetes.io/component: runtime-server
    app.kubernetes.io/instance: kubefox
    app.kubernetes.io/name: kubefox
spec:
  clusterIP: None
  ports:
    - name: grpc
      port: 6060
      protocol: TCP
      targetPort: 6060
---
apiVersion: v1
kind: Endpoints
metadata:
  name: kfp-runtime-server
  labels:
    app.kubernetes.io/component: runtime-server
    app.kubernetes.io/instance: kubefox
    app.kubernetes.io/name: kubefox
subsets:
  - addresses:
      - ip: <YOUR IP ADDRESS>
    ports:
      - name: grpc
        port: 6060
        protocol: TCP
```

Finally, scale down the component running on Kubernetes and redirect the service
to the component running locally on the workstation:

```bash
kubectl scale -n kubefox-system deployment/kubefox-runtime-server --replicas 0
kubectl delete -n kubefox-system service/kfp-runtime-server
kubectl apply -n kubefox-system -f ./workstation/runtime-srv-service-local.yaml
```
