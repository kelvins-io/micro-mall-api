#! /bin/bash

export COMPOSE_HTTP_TIMEOUT=500
export DOCKER_CLIENT_TIMEOUT=500
export COMPOSE_PARALLEL_LIMIT=1024
echo y | docker network prune
docker network create mall
docker-compose -f docker-compose.yml up -d

echo "配置rabbitmq，下面注释的命令手动执行"
#docker exec -it $(docker ps -aqf "name=rabbitmq") /bin/bash
#rabbitmqctl add_vhost micro-mall
#rabbitmqctl set_permissions -p micro-mall root ".*" ".*" ".*"
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare exchange name=user_register_notice type=direct
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare exchange name=user_state_notice type=direct
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare exchange name=trade_order_notice type=direct
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare exchange name=trade_order_pay_callback type=direct
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare exchange name=trade_pay_notice type=direct
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare exchange name=shop_info_search_notice type=direct
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare exchange name=sku_inventory_search_notice type=direct
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare exchange name=user_info_search_notice type=direct
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare exchange name=trade_order_info_search_notice type=direct
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare queue name=user_register_notice
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare queue name=user_state_notice
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare queue name=trade_order_notice
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare queue name=trade_order_pay_callback
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare queue name=trade_pay_notice
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare queue name=shop_info_search_notice
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare queue name=sku_inventory_search_notice
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare queue name=user_info_search_notice
#rabbitmqadmin -u root -p micro-mall -V micro-mall declare queue name=trade_order_info_search_notice
#exit

echo "配置elasticsearch，下面注释的命令手动执行"
#docker exec -it $(docker ps -aqf "name=elasticsearch") /bin/bash
#echo y |elasticsearch-plugin install https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v7.5.2/elasticsearch-analysis-ik-7.5.2.zip
#curl -X PUT "localhost:9200/micro-mall-user-info?pretty"
#curl -X PUT "localhost:9200/micro-mall-merchants-material-info?pretty"
#curl -X PUT "localhost:9200/micro-mall-shop?pretty"
#curl -X PUT "localhost:9200/micro-mall-trade-order?pretty"
#curl -X PUT "localhost:9200/micro-mall-sku-inventory?pretty"
#curl -X GET "localhost:9200/_cat/indices?v"
#exit

docker ps