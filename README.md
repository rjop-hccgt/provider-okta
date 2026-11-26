# Provider Okta

`provider-okta` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Okta API.

The provider ships every managed resource in two scopes:

| Scope | API group | Configured by |
| --- | --- | --- |
| Cluster (legacy) | `*.okta.upbound.io` | `ProviderConfig.okta.upbound.io` |
| Namespaced (Crossplane v2) | `*.okta.m.upbound.io` | `ClusterProviderConfig.okta.m.upbound.io` or a namespaced `ProviderConfig.okta.m.upbound.io` |

## Getting Started

Install the provider, changing the tag to the latest release:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-okta
spec:
  package: xpkg.upbound.io/crossplane-contrib/provider-okta:v0.1.0
EOF
```

Create the credentials secret from
[`examples/cluster/providerconfig/secret.yaml.tmpl`](examples/cluster/providerconfig/secret.yaml.tmpl), then
configure the provider for the scope you use:

- Cluster-scoped resources: [`examples/cluster/providerconfig/providerconfig.yaml`](examples/cluster/providerconfig/providerconfig.yaml)
- Namespaced resources: [`examples/namespaced/providerconfig/clusterproviderconfig.yaml`](examples/namespaced/providerconfig/clusterproviderconfig.yaml).
  A namespaced managed resource with no `providerConfigRef` uses the
  `ClusterProviderConfig` named `default`. A namespaced `ProviderConfig`
  ([example](examples/namespaced/providerconfig/providerconfig.yaml)) reads its
  secret from the managed resource's own namespace.

Example managed resources are in [`examples/cluster`](examples/cluster) and
[`examples/namespaced`](examples/namespaced).

You can see the API reference [here](https://doc.crds.dev/github.com/crossplane-contrib/provider-okta).

## Upgrading from releases built with Upjet v0.x

This release is built with Upjet v2 and crossplane-runtime v2.

- Existing cluster-scoped resources (`*.okta.upbound.io`) and their
  `ProviderConfig` keep working on Crossplane v1 and v2.
- Namespaced resources (`*.okta.m.upbound.io`) can only be composed on
  Crossplane v2.
- The `StoreConfig` API and the `--enable-external-secret-stores` flag were
  removed. As a result, `spec.publishConnectionDetailsTo` no longer exists on
  any managed resource. Remove it from your manifests before upgrading; use
  `spec.writeConnectionSecretToRef` instead.
- Management policies are enabled by default
  (`--enable-management-policies=true`).
- The provider declares the Crossplane `SafeStart` capability. It starts each
  controller only after its CRD is installed.

## Developing

Run the code-generation pipeline (needs Go and `goimports` on `PATH`):
```console
make generate
```

Run unit tests:
```console
make test
```

Run against a Kubernetes cluster:

```console
make run
```

Build binary and package:

```console
make build
```

Deploy to a local kind cluster with Crossplane v2:

```console
make local-deploy
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/crossplane-contrib/provider-okta/issues).
