package constants

import "time"

const (
	InitControllerName            = "virtual-cluster-init-controller"
	NodeControllerName            = "virtual-cluster-node-controller"
	KosmosJoinControllerName      = "kosmos-join-controller"
	KosmosNs                      = "kosmos-system"
	SystemNs                      = "kube-system"
	DefauleImageRepositoryEnv     = "IMAGE_REPOSITIRY"
	DefauleImageVersionEnv        = "IMAGE_VERSION"
	VirtualClusterFinalizerName   = "kosmos.io/virtual-cluster-finalizer"
	ServiceType                   = "NodePort"
	EtcdServiceType               = "ClusterIP"
	DisableCascadingDeletionLabel = "operator.virtualcluster.io/disable-cascading-deletion"
	ControllerFinalizerName       = "operator.virtualcluster.io/finalizer"
	DefaultKubeconfigPath         = "/etc/cluster-tree/cert"
	Label                         = "virtualCluster-app"
	ComponentBeReadyTimeout       = 300 * time.Second
	ComponentBeDeletedTimeout     = 300 * time.Second

	// CertificateBlockType is a possible value for pem.Block.Type.
	CertificateBlockType           = "CERTIFICATE"
	RsaKeySize                     = 2048
	KeyExtension                   = ".key"
	CertExtension                  = ".crt"
	CertificateValidity            = time.Hour * 24 * 365
	CaCertAndKeyName               = "ca"
	VirtualClusterCertAndKeyName   = "virtualCluster"
	VirtualClusterSystemNamespace  = "virtualCluster-system"
	ApiserverCertAndKeyName        = "apiserver"
	EtcdCaCertAndKeyName           = "etcd-ca"
	EtcdServerCertAndKeyName       = "etcd-server"
	EtcdClientCertAndKeyName       = "etcd-client"
	FrontProxyCaCertAndKeyName     = "front-proxy-ca"
	FrontProxyClientCertAndKeyName = "front-proxy-client"

	//controlplane apiserver
	ApiServer                     = "apiserver"
	ApiServerReplicas             = 2
	ApiServerServiceSubnet        = "10.237.6.18/29"
	ApiServerEtcdListenClientPort = 2379
	ApiServerServiceType          = "NodePort"
	// APICallRetryInterval defines how long kubeadm should wait before retrying a failed API operation
	ApiServerCallRetryInterval = 100 * time.Millisecond

	//controlplane etcd
	Etcd                 = "etcd"
	EtcdReplicas         = 3
	EtcdDataVolumeName   = "etcd-data"
	EtcdListenClientPort = 2379
	EtcdListenPeerPort   = 2380

	//controlplane kube-controller
	KubeControllerReplicas         = 2
	KubeControllerManagerComponent = "KubeControllerManager"
	KubeControllerManager          = "kube-controller-manager"

	//controlplane scheduler
	VirtualClusterSchedulerReplicas           = 2
	VirtualClusterSchedulerComponent          = "VirtualClusterScheduler"
	VirtualClusterSchedulerComponentConfigMap = "scheduler-config"
	VirtualClusterScheduler                   = "scheduler"
	VirtualClusterKubeProxyComponent          = "kube-proxy"

	//controlplane auth
	AdminConfig        = "admin-config"
	KubeConfig         = "kubeconfig"
	KubeProxyConfigmap = "kube-proxy"

	//controlplane upload
	VirtualClusterLabelKeyName = "app.kubernetes.io/managed-by"
	VirtualClusterController   = "virtual-cluster-controller"
	ClusterName                = "virtualCluster-apiserver"
	UserName                   = "virtualCluster-admin"

	// InitAction represents init virtual cluster instance
	InitAction Action = "init"
	// DeInitAction represents delete virtual cluster instance
	DeInitAction Action = "deInit"

	ManifestComponentsConfigmap = "components-manifest-cm"
	NodePoolConfigmap           = "node-pool"
	NodeShareState              = "share"
	NodeVirtualclusterState     = "virtualcluster"
	NodeFreeState               = "free"

	WaitAllPodsRunningTimeoutSeconds = 1800
)

type Action string
