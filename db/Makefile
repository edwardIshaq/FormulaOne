MYSQL_CONTAINER_NAME := mysqlContainer
MYSQL_PORT_MAPPING := 3307:3306
MYSQL_ROOT_PASSWORD := qwe123
MYSQL_NETWORK := my-bridge
MYSQL_VOLUME := ~/Workspace/mysql/containerData
MYSQL_PERSISTENT_VOLUME_YAML := mysql-pv.yaml
MYSQL_DEPLOYMENT_YAML := mysql-deployment.yaml

print:
	@echo MYSQL_CONTAINER_NAME=${MYSQL_CONTAINER_NAME}
	@echo MYSQL_PORT_MAPPING=${MYSQL_PORT_MAPPING}
	@echo MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD}
	@echo MYSQL_NETWORK=${MYSQL_NETWORK}
	@echo MYSQL_VOLUME=${MYSQL_VOLUME}


deploy: deploy-pvc deploy-mysql
deploy-pvc:
	kubectl apply -f ${MYSQL_PERSISTENT_VOLUME_YAML}
deploy-mysql:
	kubectl apply -f ${MYSQL_DEPLOYMENT_YAML}

destroy: destroy-pvc destroy-mysql
destroy-pvc:
	kubectl delete pvc f1-mysql-pv-claim
	kubectl delete pv f1-mysql-pv-volume
destroy-mysql:
	kubectl delete deployment,svc mysql

	# cp ergast-mrd/f1db.sql /Users/edward/Workspace/mysql/f1-db-mount/
load-data: POD=$(shell kubectl get pods -l app=mysql | grep Running | awk '{print $$1}')
load-data:
	kubectl exec ${POD} -- mysql -h mysql -p${MYSQL_ROOT_PASSWORD} < initialize.sql


connect:
	kubectl run -it --rm --image=mysql:5.6 --restart=Never mysql-client -- mysql -h mysql -p${MYSQL_ROOT_PASSWORD}

docker-run:
	docker run \
	--name ${MYSQL_CONTAINER_NAME} \
	--network ${MYSQL_NETWORK} \
	-p ${MYSQL_PORT_MAPPING} \
	-e MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} \
	-v ${MYSQL_VOLUME}:/var/lib/mysql \
	-d mysql

docker-attach:
	@docker exec -it ${MYSQL_CONTAINER_NAME} bash
