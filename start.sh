#!/bin/bash
set -ex

# SET THE FOLLOWING VARIABLES
# docker hub username
USERNAME=yunabe
# image name
IMAGE=lgo
version=20180525

create_go_jupyter()
{
docker run --name go-jupyter -p 8888:8888 -p 80:5000 -v $(pwd)/material:/examples -d $USERNAME/$IMAGE:$version jupyter notebook --ip=0.0.0.0 --NotebookApp.token=''
docker exec go-jupyter cp -r /examples/* /home/gopher/
}

# 必須偵測container是否已經存在，進而做判斷
if [ `docker ps -a |grep go-jupyter | wc -l`  =  0  ]; then
  create_go_jupyter
else
  docker rm -f go-jupyter
  create_go_jupyter
fi
