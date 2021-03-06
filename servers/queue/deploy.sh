# deploy to aws server

./build.sh


ssh ec2-user@api.bfranzen.me "
    docker rm -f pyq;

    docker pull bfranzen1/queue &&
    docker run -d -t \
    --name pyq \
    --network apinet \
    -e ADDR='pyq' \
    -e RABBITMQ_HOST='rmq' \
    -e RABBITMQ_PORT=5672 \
    -e RABBITMQ_USER='guest' \
    -e RABBITMQ_PW='guest' \
    -e RMQUEUE='devices' \
    -e DM_USER='ericjwei@uw.edu' \
    -e DM_PW='NrvnFFjG' \
    -e STOMP_PORT=61612 \
    -e AMQ_BROKER='alert5.eew.shakealert.org' \
    -e TEST_BROKER='eew-test1.wr.usgs.gov' \
    -e MGO_HOST='mgo' \
    -e MGO_PORT=27017 \
    -e PYTHONUNBUFFERED=0 \
    bfranzen1/queue;

    exit
"