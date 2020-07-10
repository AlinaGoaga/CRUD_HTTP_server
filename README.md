# CRUD_HTTP_server - WIP 

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

