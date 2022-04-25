```bash
##############################################
$ kubebuilder init --domain "apache.org" --project-name rocketmq-operator --license apache2 --repo github.com/apache/rocketmq-operator
Writing kustomize manifests for you to edit...
Writing scaffold for you to edit...
Get controller runtime:
$ go get sigs.k8s.io/controller-runtime@v0.11.0
go: downloading sigs.k8s.io/controller-runtime v0.11.0
go: downloading k8s.io/client-go v0.23.0
go: downloading github.com/prometheus/client_golang v1.11.0
go: downloading k8s.io/component-base v0.23.0
go: downloading golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac
go: downloading github.com/evanphx/json-patch v4.12.0+incompatible
go: downloading gomodules.xyz/jsonpatch/v2 v2.2.0
go: downloading sigs.k8s.io/structured-merge-diff/v4 v4.2.0
go: downloading github.com/google/go-cmp v0.5.5
go: downloading k8s.io/kube-openapi v0.0.0-20211115234752-e816edb12b65
go: downloading github.com/google/uuid v1.1.2
go: downloading github.com/pkg/errors v0.9.1
go: downloading github.com/prometheus/client_model v0.2.0
go: downloading github.com/prometheus/common v0.28.0
go: downloading github.com/beorn7/perks v1.0.1
go: downloading github.com/cespare/xxhash/v2 v2.1.1
go: downloading github.com/golang/protobuf v1.5.2
go: downloading github.com/prometheus/procfs v0.6.0
go: downloading golang.org/x/sys v0.0.0-20211029165221-6e7872819dc8
go: downloading github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
go: downloading github.com/googleapis/gnostic v0.5.5
go: downloading github.com/fsnotify/fsnotify v1.5.1
go: downloading github.com/imdario/mergo v0.3.12
go: downloading golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b
go: downloading google.golang.org/protobuf v1.27.1
go: downloading golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
go: downloading github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369
go: downloading google.golang.org/appengine v1.6.7
Update dependencies:
$ go mod tidy
go: downloading github.com/go-logr/zapr v1.2.0
go: downloading go.uber.org/zap v1.19.1
go: downloading github.com/stretchr/testify v1.7.0
go: downloading github.com/onsi/ginkgo v1.16.5
go: downloading github.com/onsi/gomega v1.17.0
go: downloading github.com/Azure/go-autorest/autorest v0.11.18
go: downloading github.com/Azure/go-autorest/autorest/adal v0.9.13
go: downloading go.uber.org/goleak v1.1.12
go: downloading cloud.google.com/go v0.81.0
go: downloading github.com/Azure/go-autorest v14.2.0+incompatible
go: downloading github.com/pmezard/go-difflib v1.0.0
go: downloading gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f
go: downloading github.com/Azure/go-autorest/tracing v0.6.0
go: downloading github.com/Azure/go-autorest/autorest/mocks v0.4.1
go: downloading github.com/Azure/go-autorest/autorest/date v0.3.0
go: downloading github.com/Azure/go-autorest/logger v0.2.1
go: downloading github.com/form3tech-oss/jwt-go v3.2.3+incompatible
go: downloading golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
go: downloading go.uber.org/atomic v1.7.0
go: downloading go.uber.org/multierr v1.6.0
go: downloading github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e
go: downloading github.com/benbjohnson/clock v1.1.0
go: downloading github.com/nxadm/tail v1.4.8
go: downloading github.com/kr/text v0.2.0
go: downloading github.com/cespare/xxhash v1.1.0
Next: define a resource with:
$ kubebuilder create api

##############################################
kustomize v4.5.4
kubebuilder "3.3.0"
go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.8.0
go install sigs.k8s.io/controller-runtime/tools/setup-envtest v0.0.0-20220423154536-b1e1a4f79554

./bin/controller-gen
./bin/kustomize
./bin/setup-envtest

##############################################
$ kubebuilder create api --group rocketmq --version v1alpha1 --kind Broker
Create Resource [y/n]
y
Create Controller [y/n]
y
Writing kustomize manifests for you to edit...
Writing scaffold for you to edit...
api/v1alpha1/broker_types.go
controllers/broker_controller.go
Update dependencies:
$ go mod tidy
Running make:
$ make generate
/Users/huzhi/work/code/go_code/rocketmq/rocketmq-operator/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
Next: implement your new API and generate the manifests (e.g. CRDs,CRs) with:
$ make manifests

##############################################
$ kubebuilder create api --group rocketmq --version v1alpha1 --kind Console
Create Resource [y/n]
y
Create Controller [y/n]
y
Writing kustomize manifests for you to edit...
Writing scaffold for you to edit...
api/v1alpha1/console_types.go
controllers/console_controller.go
Update dependencies:
$ go mod tidy
Running make:
$ make generate
/Users/huzhi/work/code/go_code/rocketmq/rocketmq-operator/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
Next: implement your new API and generate the manifests (e.g. CRDs,CRs) with:
$ make manifests

##############################################
$ kubebuilder create api --group rocketmq --version v1alpha1 --kind NameService
Create Resource [y/n]
y
Create Controller [y/n]
y
Writing kustomize manifests for you to edit...
Writing scaffold for you to edit...
api/v1alpha1/nameservice_types.go
controllers/nameservice_controller.go
Update dependencies:
$ go mod tidy
Running make:
$ make generate
/Users/huzhi/work/code/go_code/rocketmq/rocketmq-operator/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
Next: implement your new API and generate the manifests (e.g. CRDs,CRs) with:
$ make manifests

##############################################
$ kubebuilder create api --group rocketmq --version v1alpha1 --kind TopicTransfer
Create Resource [y/n]
y
Create Controller [y/n]
y
Writing kustomize manifests for you to edit...
Writing scaffold for you to edit...
api/v1alpha1/topictransfer_types.go
controllers/topictransfer_controller.go
Update dependencies:
$ go mod tidy
Running make:
$ make generate
/Users/huzhi/work/code/go_code/rocketmq/rocketmq-operator/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
Next: implement your new API and generate the manifests (e.g. CRDs,CRs) with:
$ make manifests

##############################################
$ git status
On branch master
Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	deleted:    .gitignore
	modified:   PROJECT
	modified:   go.mod
	modified:   main.go

Untracked files:
  (use "git add <file>..." to include in what will be committed)
	.gitignore-bak
	README.md
	api/
	bin/
	config/crd/
	config/rbac/broker_editor_role.yaml
	config/rbac/broker_viewer_role.yaml
	config/rbac/console_editor_role.yaml
	config/rbac/console_viewer_role.yaml
	config/rbac/nameservice_editor_role.yaml
	config/rbac/nameservice_viewer_role.yaml
	config/rbac/topictransfer_editor_role.yaml
	config/rbac/topictransfer_viewer_role.yaml
	config/samples/
	controllers/

no changes added to commit (use "git add" and/or "git commit -a")

##############################################
diff --git a/PROJECT b/PROJECT
index c091bfa..e4ae10b 100644
--- a/PROJECT
+++ b/PROJECT
@@ -3,4 +3,41 @@ layout:
 - go.kubebuilder.io/v3
 projectName: rocketmq-operator
 repo: github.com/apache/rocketmq-operator
+resources:
+- api:
+    crdVersion: v1
+    namespaced: true
+  controller: true
+  domain: rocketmq.apache.org
+  group: rocketmq.apache.org
+  kind: Broker
+  path: github.com/apache/rocketmq-operator/api/v1alpha1
+  version: v1alpha1
+- api:
+    crdVersion: v1
+    namespaced: true
+  controller: true
+  domain: rocketmq.apache.org
+  group: rocketmq.apache.org
+  kind: Console
+  path: github.com/apache/rocketmq-operator/api/v1alpha1
+  version: v1alpha1
+- api:
+    crdVersion: v1
+    namespaced: true
+  controller: true
+  domain: rocketmq.apache.org
+  group: rocketmq.apache.org
+  kind: NameService
+  path: github.com/apache/rocketmq-operator/api/v1alpha1
+  version: v1alpha1
+- api:
+    crdVersion: v1
+    namespaced: true
+  controller: true
+  domain: rocketmq.apache.org
+  group: rocketmq.apache.org
+  kind: TopicTransfer
+  path: github.com/apache/rocketmq-operator/api/v1alpha1
+  version: v1alpha1
 version: "3"

##############################################
diff --git a/go.mod b/go.mod
index e686db6..a4c2748 100644
--- a/go.mod
+++ b/go.mod
@@ -3,6 +3,8 @@ module github.com/apache/rocketmq-operator
 go 1.17

 require (
+       github.com/onsi/ginkgo v1.16.5
+       github.com/onsi/gomega v1.17.0
        k8s.io/apimachinery v0.23.0
        k8s.io/client-go v0.23.0
        sigs.k8s.io/controller-runtime v0.11.0
@@ -36,6 +38,7 @@ require (
        github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
        github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
        github.com/modern-go/reflect2 v1.0.2 // indirect
+       github.com/nxadm/tail v1.4.8 // indirect
        github.com/pkg/errors v0.9.1 // indirect
        github.com/prometheus/client_golang v1.11.0 // indirect
        github.com/prometheus/client_model v0.2.0 // indirect
@@ -56,6 +59,7 @@ require (
        google.golang.org/appengine v1.6.7 // indirect
        google.golang.org/protobuf v1.27.1 // indirect
        gopkg.in/inf.v0 v0.9.1 // indirect
+       gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
        gopkg.in/yaml.v2 v2.4.0 // indirect
        gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
        k8s.io/api v0.23.0 // indirect

##############################################
diff --git a/main.go b/main.go
index 52cfa40..f8327ee 100644
--- a/main.go
+++ b/main.go
@@ -30,6 +30,9 @@ import (
        ctrl "sigs.k8s.io/controller-runtime"
        "sigs.k8s.io/controller-runtime/pkg/healthz"
        "sigs.k8s.io/controller-runtime/pkg/log/zap"
+
+       rocketmqapacheorgv1alpha1 "github.com/apache/rocketmq-operator/api/v1alpha1"
+       "github.com/apache/rocketmq-operator/controllers"
        //+kubebuilder:scaffold:imports
 )

@@ -41,6 +44,7 @@ var (
 func init() {
        utilruntime.Must(clientgoscheme.AddToScheme(scheme))

+       utilruntime.Must(rocketmqapacheorgv1alpha1.AddToScheme(scheme))
        //+kubebuilder:scaffold:scheme
 }

@@ -74,6 +78,34 @@ func main() {
                os.Exit(1)
        }

+       if err = (&controllers.BrokerReconciler{
+               Client: mgr.GetClient(),
+               Scheme: mgr.GetScheme(),
+       }).SetupWithManager(mgr); err != nil {
+               setupLog.Error(err, "unable to create controller", "controller", "Broker")
+               os.Exit(1)
+       }
+       if err = (&controllers.ConsoleReconciler{
+               Client: mgr.GetClient(),
+               Scheme: mgr.GetScheme(),
+       }).SetupWithManager(mgr); err != nil {
+               setupLog.Error(err, "unable to create controller", "controller", "Console")
+               os.Exit(1)
+       }
+       if err = (&controllers.NameServiceReconciler{
+               Client: mgr.GetClient(),
+               Scheme: mgr.GetScheme(),
+       }).SetupWithManager(mgr); err != nil {
+               setupLog.Error(err, "unable to create controller", "controller", "NameService")
+               os.Exit(1)
+       }
+       if err = (&controllers.TopicTransferReconciler{
+               Client: mgr.GetClient(),
+               Scheme: mgr.GetScheme(),
+       }).SetupWithManager(mgr); err != nil {
+               setupLog.Error(err, "unable to create controller", "controller", "TopicTransfer")
+               os.Exit(1)
+       }
        //+kubebuilder:scaffold:builder

        if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {

##############################################
##############################################
##############################################
##############################################
##############################################
##############################################





```
