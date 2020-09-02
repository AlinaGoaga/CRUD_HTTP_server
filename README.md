你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# CRUD_HTTP_server

Simple REST API to retrieve, add and delete books. Data is managed in the app via variables (rather than interacting with a database).

This requires the instalation of gorilla/mux: `go get -u github.com/gorilla/mux`

* Run the app: `go run main.go` 

* Run the tests: `go test -v`

* App image at: alinag1/mycrudapp

* Create deployment: `kubectl create -f deployment.yml`
* Create service: `kubectl create -f mycrudapp-svc.yml`
* See resources: `kubectl get all`

* Get access to one of the pods and open terminal: `kubectl exec -it pod/mycrudapp-7d65644c4-lkqbw -- /bin/sh`

* Check app: `curl http://mycrudapp-svc:5000/books` or in the browser by running

`kubectl port-forward deployment/mycrudapp 5000:5000` (visit `http://localhost:5000/`)

Expose the app to the outside world with ingress: 

* Create the ingress resource: `kubectl apply -f ingress.yml`

* Add the NGINX ingress controller on the cluster: `https://kind.sigs.k8s.io/docs/user/ingress/`
I am using kind locally. 

* The app should now be visible from outside the cluster (check with `curl http://localhost` or directly in the browser).

