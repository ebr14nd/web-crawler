.PHONY=build-wcserver,build-wcclient,push-wcserver,push-wcclient,deploy-wcserver,deploy-wcclient

build-wcserver:
	  docker image build . -f Dockerfile.wcserver -t ebr41nd/wc-server

build-wcclient:
	  docker image build . -f Dockerfile.wcclient -t ebr41nd/wc-client

push-wcserver:
	  docker image push ebr41nd/wc-server

push-wcclient:
	  docker image push ebr41nd/wc-client

deploy-wcserver:
	  kubectl apply -f k8s/wc-sevrer

deploy-wcclient:
	  kubectl apply -f k8s/wc-client