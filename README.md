# CRUD_HTTP_server - WIP 

Simple REST API to retrieve, add and delete books. Data is managed in the app via variables (rather than interacting with a database).

This requires the instalation of gorilla/mux: `go get -u github.com/gorilla/mux`

Run the app: `go run main.go` 

Run the tests: `go test -v`

App image at: alinag1/mycrudapp

Create deployment: `kubectl create -f deployment.yml`
Create service: `kubectl create mycrudapp-svc.yml`
See resources: `kubectl get all`

Create ingress: `kubectl create -f ingress.yml`
See ingress: `kubectl get ingress mycrudapp`

Get access to one of the pods and open terminal: `kubectl exec -it pod/mycrudapp-7d65644c4-lkqbw -- /bin/sh`

`curl http://mycrudapp-svc:5000/books` or check app in the browser: 

`kubectl port-forward deployment/mycrudapp 5000:5000` (visit `http://localhost:5000/`)