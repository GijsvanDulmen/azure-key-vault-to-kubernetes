---
title: "Requirements"
description: "Requirements for installing akv2k8s"
---

* Kubernetes version >= 1.15.11
* A dedicated kubernetes namespace (e.g. akv2k8s)
* Enabled admission controllers: MutatingAdmissionWebhook and ValidatingAdmissionWebhook
* RBAC enabled
* Default [authentication](../security/authentication) requires Azure AKS - use [custom authentication](../security/authentication) if running outside Azure AKS.

## Dedicated namespace for akv2k8s

Akv2k8s should be installed in a dedicated Kubernetes namespace **NOT** label with `azure-key-vault-env-injection: enabled`.

**If the namespace where the akv2k8s components is installed has the injector enabled (`azure-key-vault-env-injection: enabled`), the Env Injector will most likely not be able to start.** This is because the Env Injector mutating webhook will trigger for every pod about to start in namespaces where enabled, and in the home namepsace of the Env Injector, it will effectively point to itself, which does not exist yet.

**The simple rule to avoid any issues related to this, is to just install akv2k8s components in its own dedicated namespace.**