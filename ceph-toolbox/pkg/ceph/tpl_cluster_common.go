package ceph

const SplitString = "---"
const ClusterCommon = `apiVersion: v1
kind: Namespace
metadata:
  name: rook-ceph
# OLM: BEGIN CEPH CRD
# The CRD declarations
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: cephclusters.ceph.rook.io
spec:
  group: ceph.rook.io
  names:
    kind: CephCluster
    listKind: CephClusterList
    plural: cephclusters
    singular: cephcluster
  scope: Namespaced
  version: v1
  validation:
    openAPIV3Schema:
      properties:
        spec:
          properties:
            annotations: {}
            cephVersion:
              properties:
                allowUnsupported:
                  type: boolean
                image:
                  type: string
            dashboard:
              properties:
                enabled:
                  type: boolean
                urlPrefix:
                  type: string
                port:
                  type: integer
                  minimum: 0
                  maximum: 65535
                ssl:
                  type: boolean
            dataDirHostPath:
              pattern: ^/(\S+)
              type: string
            mon:
              properties:
                allowMultiplePerNode:
                  type: boolean
                count:
                  maximum: 9
                  minimum: 0
                  type: integer
            network:
              properties:
                hostNetwork:
                  type: boolean
            storage:
              properties:
                useAllNodes:
                  type: boolean
                nodes:
                  items:
                    properties:
                      name:
                        type: string
                      config:
                        properties:
                          metadataDevice:
                            type: string
                          storeType:
                            type: string
                            pattern: ^(filestore|bluestore)$
                          databaseSizeMB:
                            type: string
                          walSizeMB:
                            type: string
                          journalSizeMB:
                            type: string
                          osdsPerDevice:
                            type: string
                          encryptedDevice:
                            type: string
                            pattern: ^(true|false)$
                      useAllDevices:
                        type: boolean
                      deviceFilter: {}
                      directories:
                        type: array
                        items:
                          properties:
                            path:
                              type: string
                      devices:
                        type: array
                        items:
                          properties:
                            name:
                              type: string
                            config: {}
                      location: {}
                      resources: {}
                  type: array
                useAllDevices:
                  type: boolean
                deviceFilter: {}
                location: {}
                directories:
                  type: array
                  items:
                    properties:
                      path:
                        type: string
                config: {}
                topologyAware:
                  type: boolean
            monitoring:
              properties:
                enabled:
                  type: boolean
                rulesNamespace:
                  type: string
            rbdMirroring:
              properties:
                workers:
                  type: integer
            placement: {}
            resources: {}
            configOverrides:
              items:
                properties:
                  who:
                    type: string
                  option:
                    type: string
                  value:
                    type: string
  additionalPrinterColumns:
    - name: DataDirHostPath
      type: string
      description: Directory used on the K8s nodes
      JSONPath: .spec.dataDirHostPath
    - name: MonCount
      type: string
      description: Number of MONs
      JSONPath: .spec.mon.count
    - name: Age
      type: date
      JSONPath: .metadata.creationTimestamp
    - name: State
      type: string
      description: Current State
      JSONPath: .status.state
    - name: Health
      type: string
      description: Ceph Health
      JSONPath: .status.ceph.health
# OLM: END CEPH CRD
# OLM: BEGIN CEPH FS CRD
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: cephfilesystems.ceph.rook.io
spec:
  group: ceph.rook.io
  names:
    kind: CephFilesystem
    listKind: CephFilesystemList
    plural: cephfilesystems
    singular: cephfilesystem
  scope: Namespaced
  version: v1
  validation:
    openAPIV3Schema:
      properties:
        spec:
          properties:
            metadataServer:
              properties:
                activeCount:
                  type: integer
                activeStandby:
                  type: boolean
                annotations: {}
                placement: {}
                resources: {}
            metadataPool:
              properties:
                failureDomain:
                  type: string
                replicated:
                  properties:
                    size:
                      type: integer
                erasureCoded:
                  properties:
                    dataChunks:
                      type: integer
                    codingChunks:
                      type: integer
            dataPools:
              type: array
              items:
                properties:
                  failureDomain:
                    type: string
                  replicated:
                    properties:
                      size:
                        type: integer
                  erasureCoded:
                    properties:
                      dataChunks:
                        type: integer
                      codingChunks:
                        type: integer
  additionalPrinterColumns:
    - name: MdsCount
      type: string
      description: Number of MDSs
      JSONPath: .spec.metadataServer.activeCount
    - name: Age
      type: date
      JSONPath: .metadata.creationTimestamp
# OLM: END CEPH FS CRD
# OLM: BEGIN CEPH NFS CRD
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: cephnfses.ceph.rook.io
spec:
  group: ceph.rook.io
  names:
    kind: CephNFS
    listKind: CephNFSList
    plural: cephnfses
    singular: cephnfs
    shortNames:
    - nfs
  scope: Namespaced
  version: v1
  validation:
    openAPIV3Schema:
      properties:
        spec:
          properties:
            rados:
              properties:
                pool:
                  type: string
                namespace:
                  type: string
            server:
              properties:
                active:
                  type: integer
                annotations: {}
                placement: {}
                resources: {}

# OLM: END CEPH NFS CRD
# OLM: BEGIN CEPH OBJECT STORE CRD
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: cephobjectstores.ceph.rook.io
spec:
  group: ceph.rook.io
  names:
    kind: CephObjectStore
    listKind: CephObjectStoreList
    plural: cephobjectstores
    singular: cephobjectstore
  scope: Namespaced
  version: v1
  validation:
    openAPIV3Schema:
      properties:
        spec:
          properties:
            gateway:
              properties:
                type:
                  type: string
                sslCertificateRef: {}
                port:
                  type: integer
                securePort: {}
                instances:
                  type: integer
                annotations: {}
                placement: {}
                resources: {}
            metadataPool:
              properties:
                failureDomain:
                  type: string
                replicated:
                  properties:
                    size:
                      type: integer
                erasureCoded:
                  properties:
                    dataChunks:
                      type: integer
                    codingChunks:
                      type: integer
            dataPool:
              properties:
                failureDomain:
                  type: string
                replicated:
                  properties:
                    size:
                      type: integer
                erasureCoded:
                  properties:
                    dataChunks:
                      type: integer
                    codingChunks:
                      type: integer
# OLM: END CEPH OBJECT STORE CRD
# OLM: BEGIN CEPH OBJECT STORE USERS CRD
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: cephobjectstoreusers.ceph.rook.io
spec:
  group: ceph.rook.io
  names:
    kind: CephObjectStoreUser
    listKind: CephObjectStoreUserList
    plural: cephobjectstoreusers
    singular: cephobjectstoreuser
  scope: Namespaced
  version: v1
# OLM: END CEPH OBJECT STORE USERS CRD
# OLM: BEGIN CEPH BLOCK POOL CRD
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: cephblockpools.ceph.rook.io
spec:
  group: ceph.rook.io
  names:
    kind: CephBlockPool
    listKind: CephBlockPoolList
    plural: cephblockpools
    singular: cephblockpool
  scope: Namespaced
  version: v1
# OLM: END CEPH BLOCK POOL CRD
# OLM: BEGIN CEPH VOLUME POOL CRD
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: volumes.rook.io
spec:
  group: rook.io
  names:
    kind: Volume
    listKind: VolumeList
    plural: volumes
    singular: volume
    shortNames:
    - rv
  scope: Namespaced
  version: v1alpha2
# OLM: END CEPH VOLUME POOL CRD
# OLM: BEGIN OBJECTBUCKET CRD
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: objectbuckets.objectbucket.io
spec:
  group: objectbucket.io
  versions:
    - name: v1alpha1
      served: true
      storage: true
  names:
    kind: ObjectBucket
    listKind: ObjectBucketList
    plural: objectbuckets
    singular: objectbucket
    shortNames:
      - ob
      - obs
  scope: Cluster
  subresources:
    status: {}
# OLM: END OBJECTBUCKET CRD
# OLM: BEGIN OBJECTBUCKETCLAIM CRD
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: objectbucketclaims.objectbucket.io
spec:
  versions:
    - name: v1alpha1
      served: true
      storage: true
  group: objectbucket.io
  names:
    kind: ObjectBucketClaim
    listKind: ObjectBucketClaimList
    plural: objectbucketclaims
    singular: objectbucketclaim
    shortNames:
      - obc
      - obcs
  scope: Namespaced
  subresources:
    status: {}
# OLM: END OBJECTBUCKETCLAIM CRD
# OLM: BEGIN OBJECTBUCKET ROLEBINDING
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: rook-ceph-object-bucket
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: rook-ceph-object-bucket
subjects:
  - kind: ServiceAccount
    name: rook-ceph-system
    namespace: rook-ceph
# OLM: END OBJECTBUCKET ROLEBINDING
# OLM: BEGIN OPERATOR ROLE
---
# The cluster role for managing all the cluster-specific resources in a namespace
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: rook-ceph-cluster-mgmt
  labels:
    operator: rook
    storage-backend: ceph
aggregationRule:
  clusterRoleSelectors:
  - matchLabels:
      rbac.ceph.rook.io/aggregate-to-rook-ceph-cluster-mgmt: "true"
rules: []
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: rook-ceph-cluster-mgmt-rules
  labels:
    operator: rook
    storage-backend: ceph
    rbac.ceph.rook.io/aggregate-to-rook-ceph-cluster-mgmt: "true"
rules:
- apiGroups:
  - ""
  resources:
  - secrets
  - pods
  - pods/log
  - services
  - configmaps
  verbs:
  - get
  - list
  - watch
  - patch
  - create
  - update
  - delete
- apiGroups:
  - apps
  resources:
  - deployments
  - daemonsets
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - delete
---
# The role for the operator to manage resources in its own namespace
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: Role
metadata:
  name: rook-ceph-system
  namespace: rook-ceph
  labels:
    operator: rook
    storage-backend: ceph
rules:
- apiGroups:
  - ""
  resources:
  - pods
  - configmaps
  - services
  verbs:
  - get
  - list
  - watch
  - patch
  - create
  - update
  - delete
- apiGroups:
  - apps
  resources:
  - daemonsets
  - statefulsets
  - deployments
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - delete
---
# The cluster role for managing the Rook CRDs
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: rook-ceph-global
  labels:
    operator: rook
    storage-backend: ceph
aggregationRule:
  clusterRoleSelectors:
  - matchLabels:
      rbac.ceph.rook.io/aggregate-to-rook-ceph-global: "true"
rules: []
---
apiVersion: rbac.authorization.k8s.io/v1beta1
kind: ClusterRole
metadata:
  name: rook-ceph-global-rules
  labels:
    operator: rook
    storage-backend: ceph
    rbac.ceph.rook.io/aggregate-to-rook-ceph-global: "true"
rules:
- apiGroups:
  - ""
  resources:
  # Pod access is needed for fencing
  - pods
  # Node access is needed for determining nodes where mons should run
  - nodes
  - nodes/proxy
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - events
    # PVs and PVCs are managed by the Rook provisioner
  - persistentvolumes
  - persistentvolumeclaims
  - endpoints
  verbs:
  - get
  - list
  - watch
  - patch
  - create
  - update
  - delete
- apiGroups:
  - storage.k8s.io
  resources:
  - storageclasses
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - batch
  resources:
  - jobs
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - delete
- apiGroups:
  - ceph.rook.io
  resources:
  - "*"
  verbs:
  - "*"
- apiGroups:
  - rook.io
  resources:
  - "*"
  verbs:
  - "*"
- apiGroups:
  - policy
  - apps
  resources:
  #this is for the clusterdisruption controller
  - poddisruptionbudgets
  #this is for both clusterdisruption and nodedrain controllers
  - deployments
  verbs:
  - "*"

---
# Aspects of ceph-mgr that require cluster-wide access
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: rook-ceph-mgr-cluster
  labels:
    operator: rook
    storage-backend: ceph
aggregationRule:
  clusterRoleSelectors:
  - matchLabels:
      rbac.ceph.rook.io/aggregate-to-rook-ceph-mgr-cluster: "true"
rules: []
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: rook-ceph-mgr-cluster-rules
  labels:
    operator: rook
    storage-backend: ceph
    rbac.ceph.rook.io/aggregate-to-rook-ceph-mgr-cluster: "true"
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  - nodes
  - nodes/proxy
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - create
  - patch
  - list
  - get
  - watch
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: rook-ceph-object-bucket
  labels:
    operator: rook
    storage-backend: ceph
    rbac.ceph.rook.io/aggregate-to-rook-ceph-mgr-cluster: "true"
rules:
- apiGroups:
  - ""
  verbs:
  - "*"
  resources:
  - secrets
  - configmaps
- apiGroups:
    - storage.k8s.io
  resources:
    - storageclasses
  verbs:
    - get
    - list
    - watch
- apiGroups:
  - "objectbucket.io"
  verbs:
  - "*"
  resources:
  - "*"
# OLM: END OPERATOR ROLE
# OLM: BEGIN SERVICE ACCOUNT SYSTEM
---
# The rook system service account used by the operator, agent, and discovery pods
apiVersion: v1
kind: ServiceAccount
metadata:
  name: rook-ceph-system
  namespace: rook-ceph
  labels:
    operator: rook
    storage-backend: ceph
# imagePullSecrets:
# - name: my-registry-secret

# OLM: END SERVICE ACCOUNT SYSTEM
# OLM: BEGIN OPERATOR ROLEBINDING
---
# Grant the operator, agent, and discovery agents access to resources in the namespace
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: rook-ceph-system
  namespace: rook-ceph
  labels:
    operator: rook
    storage-backend: ceph
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: rook-ceph-system
subjects:
- kind: ServiceAccount
  name: rook-ceph-system
  namespace: rook-ceph
---
# Grant the rook system daemons cluster-wide access to manage the Rook CRDs, PVCs, and storage classes
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: rook-ceph-global
  namespace: rook-ceph
  labels:
    operator: rook
    storage-backend: ceph
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: rook-ceph-global
subjects:
- kind: ServiceAccount
  name: rook-ceph-system
  namespace: rook-ceph
# OLM: END OPERATOR ROLEBINDING
#################################################################################################################
# Beginning of cluster-specific resources. The example will assume the cluster will be created in the "rook-ceph"
# namespace. If you want to create the cluster in a different namespace, you will need to modify these roles
# and bindings accordingly.
#################################################################################################################
# Service account for the Ceph OSDs. Must exist and cannot be renamed.
# OLM: BEGIN SERVICE ACCOUNT OSD
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: rook-ceph-osd
  namespace: rook-ceph
# imagePullSecrets:
# - name: my-registry-secret

# OLM: END SERVICE ACCOUNT OSD
# OLM: BEGIN SERVICE ACCOUNT MGR
---
# Service account for the Ceph Mgr. Must exist and cannot be renamed.
apiVersion: v1
kind: ServiceAccount
metadata:
  name: rook-ceph-mgr
  namespace: rook-ceph
# imagePullSecrets:
# - name: my-registry-secret

# OLM: END SERVICE ACCOUNT MGR
# OLM: BEGIN CMD REPORTER SERVICE ACCOUNT
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: rook-ceph-cmd-reporter
  namespace: rook-ceph
# OLM: END CMD REPORTER SERVICE ACCOUNT
# OLM: BEGIN CLUSTER ROLE
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: rook-ceph-osd
  namespace: rook-ceph
rules:
- apiGroups: [""]
  resources: ["configmaps"]
  verbs: [ "get", "list", "watch", "create", "update", "delete" ]
- apiGroups: ["ceph.rook.io"]
  resources: ["cephclusters", "cephclusters/finalizers"]
  verbs: [ "get", "list", "create", "update", "delete" ]
---
# Aspects of ceph-mgr that require access to the system namespace
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: rook-ceph-mgr-system
  namespace: rook-ceph
aggregationRule:
  clusterRoleSelectors:
  - matchLabels:
      rbac.ceph.rook.io/aggregate-to-rook-ceph-mgr-system: "true"
rules: []
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: rook-ceph-mgr-system-rules
  namespace: rook-ceph
  labels:
      rbac.ceph.rook.io/aggregate-to-rook-ceph-mgr-system: "true"
rules:
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
  - list
  - watch
---
# Aspects of ceph-mgr that operate within the cluster's namespace
kind: Role
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: rook-ceph-mgr
  namespace: rook-ceph
rules:
- apiGroups:
  - ""
  resources:
  - pods
  - services
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - batch
  resources:
  - jobs
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - delete
- apiGroups:
  - ceph.rook.io
  resources:
  - "*"
  verbs:
  - "*"
# OLM: END CLUSTER ROLE
# OLM: BEGIN CMD REPORTER ROLE
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: rook-ceph-cmd-reporter
  namespace: rook-ceph
rules:
- apiGroups:
  - ""
  resources:
  - pods
  - configmaps
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - delete
# OLM: END CMD REPORTER ROLE
# OLM: BEGIN CLUSTER ROLEBINDING
---
# Allow the operator to create resources in this cluster's namespace
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: rook-ceph-cluster-mgmt
  namespace: rook-ceph
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: rook-ceph-cluster-mgmt
subjects:
- kind: ServiceAccount
  name: rook-ceph-system
  namespace: rook-ceph
---
# Allow the osd pods in this namespace to work with configmaps
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: rook-ceph-osd
  namespace: rook-ceph
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: rook-ceph-osd
subjects:
- kind: ServiceAccount
  name: rook-ceph-osd
  namespace: rook-ceph
---
# Allow the ceph mgr to access the cluster-specific resources necessary for the mgr modules
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: rook-ceph-mgr
  namespace: rook-ceph
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: rook-ceph-mgr
subjects:
- kind: ServiceAccount
  name: rook-ceph-mgr
  namespace: rook-ceph
---
# Allow the ceph mgr to access the rook system resources necessary for the mgr modules
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: rook-ceph-mgr-system
  namespace: rook-ceph
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: rook-ceph-mgr-system
subjects:
- kind: ServiceAccount
  name: rook-ceph-mgr
  namespace: rook-ceph
---
# Allow the ceph mgr to access cluster-wide resources necessary for the mgr modules
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: rook-ceph-mgr-cluster
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: rook-ceph-mgr-cluster
subjects:
- kind: ServiceAccount
  name: rook-ceph-mgr
  namespace: rook-ceph
# OLM: END CLUSTER ROLEBINDING
# OLM: BEGIN CMD REPORTER ROLEBINDING
---
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: rook-ceph-cmd-reporter
  namespace: rook-ceph
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: rook-ceph-cmd-reporter
subjects:
- kind: ServiceAccount
  name: rook-ceph-cmd-reporter
  namespace: rook-ceph
# OLM: END CMD REPORTER ROLEBINDING
#################################################################################################################
# Beginning of pod security policy resources. The example will assume the cluster will be created in the
# "rook-ceph" namespace. If you want to create the cluster in a different namespace, you will need to modify
# the roles and bindings accordingly.
#################################################################################################################
# OLM: BEGIN CLUSTER POD SECURITY POLICY
---
apiVersion: policy/v1beta1
kind: PodSecurityPolicy
metadata:
  name: rook-privileged
spec:
  privileged: true
  allowedCapabilities:
    # required by CSI
    - SYS_ADMIN
  # fsGroup - the flexVolume agent has fsGroup capabilities and could potentially be any group
  fsGroup:
    rule: RunAsAny
  # runAsUser, supplementalGroups - Rook needs to run some pods as root
  # Ceph pods could be run as the Ceph user, but that user isn't always known ahead of time
  runAsUser:
    rule: RunAsAny
  supplementalGroups:
    rule: RunAsAny
  # seLinux - seLinux context is unknown ahead of time; set if this is well-known
  seLinux:
    rule: RunAsAny
  volumes:
    # recommended minimum set
    - configMap
    - downwardAPI
    - emptyDir
    - persistentVolumeClaim
    - secret
    - projected
    # required for Rook
    - hostPath
    - flexVolume
  # allowedHostPaths can be set to Rook's known host volume mount points when they are fully-known
  # directory-based OSDs make this hard to nail down
  # allowedHostPaths:
  #   - pathPrefix: "/run/udev"  # for OSD prep
  #     readOnly: false
  #   - pathPrefix: "/dev"  # for OSD prep
  #     readOnly: false
  #   - pathPrefix: "/var/lib/rook"  # or whatever the dataDirHostPath value is set to
  #     readOnly: false
  # Ceph requires host IPC for setting up encrypted devices
  hostIPC: true
  # Ceph OSDs need to share the same PID namespace
  hostPID: true
  # hostNetwork can be set to 'false' if host networking isn't used
  hostNetwork: true
  hostPorts:
    # Ceph messenger protocol v1
    - min: 6789
      max: 6790 # <- support old default port
    # Ceph messenger protocol v2
    - min: 3300
      max: 3300
    # Ceph RADOS ports for OSDs, MDSes
    - min: 6800
      max: 7300
    # # Ceph dashboard port HTTP (not recommended)
    # - min: 7000
    #   max: 7000
    # Ceph dashboard port HTTPS
    - min: 8443
      max: 8443
    # Ceph mgr Prometheus Metrics
    - min: 9283
      max: 9283
# OLM: END CLUSTER POD SECURITY POLICY
# OLM: BEGIN POD SECURITY POLICY BINDINGS
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: 'psp:rook'
rules:
  - apiGroups:
      - policy
    resources:
      - podsecuritypolicies
    resourceNames:
      - rook-privileged
    verbs:
      - use
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: rook-ceph-system-psp
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: 'psp:rook'
subjects:
  - kind: ServiceAccount
    name: rook-ceph-system
    namespace: rook-ceph
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: rook-ceph-default-psp
  namespace: rook-ceph
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: psp:rook
subjects:
- kind: ServiceAccount
  name: default
  namespace: rook-ceph
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: rook-ceph-osd-psp
  namespace: rook-ceph
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: psp:rook
subjects:
- kind: ServiceAccount
  name: rook-ceph-osd
  namespace: rook-ceph
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: rook-ceph-mgr-psp
  namespace: rook-ceph
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: psp:rook
subjects:
- kind: ServiceAccount
  name: rook-ceph-mgr
  namespace: rook-ceph
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: rook-ceph-cmd-reporter-psp
  namespace: rook-ceph
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: psp:rook
subjects:
- kind: ServiceAccount
  name: rook-ceph-cmd-reporter
  namespace: rook-ceph
# OLM: END CLUSTER POD SECURITY POLICY BINDINGS
# OLM: BEGIN CSI CEPHFS SERVICE ACCOUNT
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: rook-csi-cephfs-plugin-sa
  namespace: rook-ceph
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: rook-csi-cephfs-provisioner-sa
  namespace: rook-ceph
# OLM: END CSI CEPHFS SERVICE ACCOUNT
# OLM: BEGIN CSI CEPHFS ROLE
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  namespace: rook-ceph
  name: cephfs-external-provisioner-cfg
rules:
  - apiGroups: [""]
    resources: ["endpoints"]
    verbs: ["get", "watch", "list", "delete", "update", "create"]
  - apiGroups: [""]
    resources: ["configmaps"]
    verbs: ["get", "list", "create", "delete"]
  - apiGroups: ["coordination.k8s.io"]
    resources: ["leases"]
    verbs: ["get", "watch", "list", "delete", "update", "create"]
# OLM: END CSI CEPHFS ROLE
# OLM: BEGIN CSI CEPHFS ROLEBINDING
---
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: cephfs-csi-provisioner-role-cfg
  namespace: rook-ceph
subjects:
  - kind: ServiceAccount
    name: rook-csi-cephfs-provisioner-sa
    namespace: rook-ceph
roleRef:
  kind: Role
  name: cephfs-external-provisioner-cfg
  apiGroup: rbac.authorization.k8s.io
# OLM: END CSI CEPHFS ROLEBINDING
# OLM: BEGIN CSI CEPHFS CLUSTER ROLE
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: cephfs-csi-nodeplugin
aggregationRule:
  clusterRoleSelectors:
  - matchLabels:
      rbac.ceph.rook.io/aggregate-to-cephfs-csi-nodeplugin: "true"
rules: []
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: cephfs-csi-nodeplugin-rules
  labels:
    rbac.ceph.rook.io/aggregate-to-cephfs-csi-nodeplugin: "true"
rules:
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "update"]
  - apiGroups: [""]
    resources: ["namespaces"]
    verbs: ["get", "list"]
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["volumeattachments"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: [""]
    resources: ["configmaps"]
    verbs: ["get", "list"]
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: cephfs-external-provisioner-runner
aggregationRule:
  clusterRoleSelectors:
  - matchLabels:
      rbac.ceph.rook.io/aggregate-to-cephfs-external-provisioner-runner: "true"
rules: []
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: cephfs-external-provisioner-runner-rules
  labels:
    rbac.ceph.rook.io/aggregate-to-cephfs-external-provisioner-runner: "true"
rules:
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get", "list"]
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "create", "delete", "update"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["storageclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["events"]
    verbs: ["list", "watch", "create", "update", "patch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["volumeattachments"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "watch"]
# OLM: END CSI CEPHFS CLUSTER ROLE
# OLM: BEGIN CSI CEPHFS CLUSTER ROLEBINDING
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: rook-csi-cephfs-plugin-sa-psp
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: 'psp:rook'
subjects:
  - kind: ServiceAccount
    name: rook-csi-cephfs-plugin-sa
    namespace: rook-ceph
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: rook-csi-cephfs-provisioner-sa-psp
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: 'psp:rook'
subjects:
  - kind: ServiceAccount
    name: rook-csi-cephfs-provisioner-sa
    namespace: rook-ceph
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: cephfs-csi-nodeplugin
subjects:
  - kind: ServiceAccount
    name: rook-csi-cephfs-plugin-sa
    namespace: rook-ceph
roleRef:
  kind: ClusterRole
  name: cephfs-csi-nodeplugin
  apiGroup: rbac.authorization.k8s.io

---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: cephfs-csi-provisioner-role
subjects:
  - kind: ServiceAccount
    name: rook-csi-cephfs-provisioner-sa
    namespace: rook-ceph
roleRef:
  kind: ClusterRole
  name: cephfs-external-provisioner-runner
  apiGroup: rbac.authorization.k8s.io
# OLM: END CSI CEPHFS CLUSTER ROLEBINDING
# OLM: BEGIN CSI RBD SERVICE ACCOUNT
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: rook-csi-rbd-plugin-sa
  namespace: rook-ceph
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: rook-csi-rbd-provisioner-sa
  namespace: rook-ceph
# OLM: END CSI RBD SERVICE ACCOUNT
# OLM: BEGIN CSI RBD ROLE
---
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  namespace: rook-ceph
  name: rbd-external-provisioner-cfg
rules:
  - apiGroups: [""]
    resources: ["endpoints"]
    verbs: ["get", "watch", "list", "delete", "update", "create"]
  - apiGroups: [""]
    resources: ["configmaps"]
    verbs: ["get", "list", "watch", "create", "delete"]
  - apiGroups: ["coordination.k8s.io"]
    resources: ["leases"]
    verbs: ["get", "watch", "list", "delete", "update", "create"]
# OLM: END CSI RBD ROLE
# OLM: BEGIN CSI RBD ROLEBINDING
---
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rbd-csi-provisioner-role-cfg
  namespace: rook-ceph
subjects:
  - kind: ServiceAccount
    name: rook-csi-rbd-provisioner-sa
    namespace: rook-ceph
roleRef:
  kind: Role
  name: rbd-external-provisioner-cfg
  apiGroup: rbac.authorization.k8s.io
# OLM: END CSI RBD ROLEBINDING
# OLM: BEGIN CSI RBD CLUSTER ROLE
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rbd-csi-nodeplugin
aggregationRule:
  clusterRoleSelectors:
  - matchLabels:
      rbac.ceph.rook.io/aggregate-to-rbd-csi-nodeplugin: "true"
rules: []
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rbd-csi-nodeplugin-rules
  labels:
    rbac.ceph.rook.io/aggregate-to-rbd-csi-nodeplugin: "true"
rules:
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get", "list"]
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "update"]
  - apiGroups: [""]
    resources: ["namespaces"]
    verbs: ["get", "list"]
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["volumeattachments"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: [""]
    resources: ["configmaps"]
    verbs: ["get", "list"]
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rbd-external-provisioner-runner
aggregationRule:
  clusterRoleSelectors:
  - matchLabels:
      rbac.ceph.rook.io/aggregate-to-rbd-external-provisioner-runner: "true"
rules: []
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rbd-external-provisioner-runner-rules
  labels:
    rbac.ceph.rook.io/aggregate-to-rbd-external-provisioner-runner: "true"
rules:
  - apiGroups: [""]
    resources: ["secrets"]
    verbs: ["get", "list"]
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "create", "delete", "update"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["volumeattachments"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["storageclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["events"]
    verbs: ["list", "watch", "create", "update", "patch"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshots"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshotcontents"]
    verbs: ["create", "get", "list", "watch", "update", "delete"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshotclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["apiextensions.k8s.io"]
    resources: ["customresourcedefinitions"]
    verbs: ["create", "list", "watch", "delete", "get", "update"]
  - apiGroups: ["snapshot.storage.k8s.io"]
    resources: ["volumesnapshots/status"]
    verbs: ["update"]
# OLM: END CSI RBD CLUSTER ROLE
# OLM: BEGIN CSI RBD CLUSTER ROLEBINDING
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: rook-csi-rbd-plugin-sa-psp
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: 'psp:rook'
subjects:
  - kind: ServiceAccount
    name: rook-csi-rbd-plugin-sa
    namespace: rook-ceph
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: rook-csi-rbd-provisioner-sa-psp
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: 'psp:rook'
subjects:
  - kind: ServiceAccount
    name: rook-csi-rbd-provisioner-sa
    namespace: rook-ceph
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rbd-csi-nodeplugin
subjects:
  - kind: ServiceAccount
    name: rook-csi-rbd-plugin-sa
    namespace: rook-ceph
roleRef:
  kind: ClusterRole
  name: rbd-csi-nodeplugin
  apiGroup: rbac.authorization.k8s.io
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: rbd-csi-provisioner-role
subjects:
  - kind: ServiceAccount
    name: rook-csi-rbd-provisioner-sa
    namespace: rook-ceph
roleRef:
  kind: ClusterRole
  name: rbd-external-provisioner-runner
  apiGroup: rbac.authorization.k8s.io
# OLM: END CSI RBD CLUSTER ROLEBINDING`
