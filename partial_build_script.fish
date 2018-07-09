#
# the main issue here is that keytransparency/docker-compose.yml references images
# we can't read from. This is an attempt to set up parent images so that
# keytransparency can build correctly
#

set PROJECT identity-188116
set ZONE us-central1-a
set CLUSTER key-transparency
gcloud config set project $PROJECT
gcloud config set compute/zone $ZONE
gcloud container clusters get-credentials $CLUSTER
gcloud --quiet auth configure-docker

#TODO: I ran a bunch of ed commands that found/replaced us.gcr.io/keystrasparency
# and substituted $PROJECT

cd $GOPATH/src/github.com/google/keytransparency
sudo ./scripts/prepare_server.sh -f

# build prometheus image
cd $GOPATH/src/github.com/google/keytransparency
sudo docker build --tag us.gcr.io/$PROJECT/db -f deploy/prometheus/Dockerfile .
sudo docker push us.gcr.io/$PROJECT/prometheus

# build trillian images
cd $GOPATH/src/github.com/google/trillian
sudo docker build --tag us.gcr.io/$PROJECT/db -f examples/deployment/docker/db_server/Dockerfile .
sudo docker build --tag us.gcr.io/$PROJECT/log-server -f examples/deployment/docker/log_server/Dockerfile .
sudo docker build --tag us.gcr.io/$PROJECT/map-server -f examples/deployment/docker/map_server/Dockerfile .
sudo docker build --tag us.gcr.io/$PROJECT/log-signer -f examples/deployment/docker/log_signer/Dockerfile .
sudo docker push us.gcr.io/$PROJECT/db
sudo docker push us.gcr.io/$PROJECT/log-server
sudo docker push us.gcr.io/$PROJECT/log-signer
sudo docker push us.gcr.io/$PROJECT/map-server

# build keytransparence images
cd $GOPATH/src/github.com/google

sudo docker build --tag us.gcr.io/$PROJECT/keytransparency-server -f ./keytransparency/cmd/keytransparency-server/Dockerfile
sudo docker build --tag us.gcr.io/$PROJECT/keytransparency-sequencer -f ./keytransparency/cmd/keytransparency-sequencer/Dockerfile
sudo docker build --tag us.gcr.io/$PROJECT/keytransparency-monitor -f ./keytransparency/cmd/keytransparency-monitor/Dockerfile
sudo docker push us.gcr.io/$PROJECT/keytransparency-monitor
sudo docker push us.gcr.io/$PROJECT/keytransparency-sequencer
sudo docker push us.gcr.io/$PROJECT/keytransparency-server


# build init image
cd $GOPATH/src/github.com/google/keytransparency
sudo docker build --tag us.gcr.io/$PROJECT/init -f deploy/docker/init/Dockerfile .
sudo docker push us.gcr.io/$PROJECT/init

# rekompose kubernetes deployments
cd $GOPATH/src/gitub.com/google/keytransparency/deploy/kubernetes
curl -L https://github.com/kubernetes/kompose/releases/download/v1.15.0/kompose-linux-amd64 -o kompose
chmod +x kompose
sudo mv ./kompose /usr/loca/bin/kompose

# TODO: I have low confidence in this part.
export TRAVIS_COMMIT=latest

# Now that parent images exist on a server, we should be able to do a build
docker-compose build









# attempt to build project
sudo -E docker-compose up -d
