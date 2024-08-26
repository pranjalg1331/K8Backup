<h2>Pre-requisites of using this project<h2>
<p>Kube config file-The .kube/config file is essential for connecting your local machine to a Kubernetes cluster. This file contains the necessary configuration details like cluster information, credentials, and context. Below are the steps to obtain and set up the .kube/config file<p>

<h2>Commands Used</h2>
<p>K8Backup pod backup --name "name of pod" --namespace "namespace of pod" --path "kube config file path"</p>
<p>K8Backup pod restore --object "file path name as in backup object" --path "kube config file path" --name "name of new resource"</p>
<p>K8Backup deployment backup --name "name of deployemnt" --namespace "namespace of deployment" --path "kube config file path"</p>
<p>K8Backup deployment restore --object "file path name as in backup object" --path "kube config file path" --name "name of new resource"</p>
<p>K8Backup volume backup --name "name of pvc" --namespace "namespace of pvc" --path "kube config file path"</p>
<p>K8Backup volume restore --object "name of colume snapshot" --path "kube config file path" --name "name of new pvc"</p>
<p>K8Backup list</p>
<p>K8Backup delete --file "filename of backup object"</p>


<h4>For volume snapshot and pvc</h4>
<p>Folloing are the commands witt minikube</p>
<p>minikube addons enable csi-hostpath-driver</p>
<p>minikube addons enable volumesnapshots</p>