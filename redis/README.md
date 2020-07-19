# Elasticache Redis README

## Create managed Redis.

Using AWS Console.


## Test using CLI
From Bastion,

1. Test connectivity.
    ```
    REDIS_HOST=redis00.v8jh5m.ng.0001.apse1.cache.amazonaws.com
    nc -vz $REDIS_HOST 6379
    ```

2. Setup redis-cli on Amazon Linux 1
    ```
    sudo yum install gcc 
    wget http://download.redis.io/redis-stable.tar.gz && tar xvzf redis-stable.tar.gz && cd redis-stable && make
    sudo yum install clang
    CC=clang make
    sudo make install
    sudo cp src/redis-cli /usr/local/bin/
    redis-cli -h $REDIS_HOST -p 6379
    rm -rf redis-stable*
    ```

    Try a few commands:
    ```
    set a "hello"
    get a
    del a
    set b "Good-bye" EX 5 
    get b
    quit
    ```

    Ref: https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/GettingStarted.ConnectToCacheNode.html

## TODO: microservice

## TODO: load test data.