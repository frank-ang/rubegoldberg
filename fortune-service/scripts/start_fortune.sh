export AWS_REGION="ap-southeast-1"
export DB_SECRET_NAME="service/db/mysql"
export REDIS_HOST="redis00.v8jh5m.ng.0001.apse1.cache.amazonaws.com"
export ES_HOST="https://vpc-sandbox00-7zz2l5gm2ihbkrgwsumu2bcpmu.ap-southeast-1.es.amazonaws.com"

# Testing on Mac:
# nohup ./fortune -mysql -redis -es > ./fortune.out 2>&1 </dev/null &

# Run background process:
nohup /opt/fortune/fortune-linux -mysql -redis -es > /opt/fortune/fortune.out 2>&1 </dev/null &
