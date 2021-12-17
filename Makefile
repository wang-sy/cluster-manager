test-pre:
	mv ~/.kube ~/.kubeback
	minikube start --cni="flannel"
	kubectl apply -f ./deployment/namespace.yaml
	kubectl apply -f ./deployment/local-test-mysql-service.yaml

test-post:
	minikube delete
	mv ~/.kubeback ~/.kube