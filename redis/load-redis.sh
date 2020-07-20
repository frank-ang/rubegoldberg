#!/bin/bash

# Import fortune test data into Redis. 

FORTUNE_CSV=$1
REDIS_HOST=$2

if [ -z "$FORTUNE_CSV" ] | [ -z "$REDIS_HOST" ]; then
    echo "Usage $0 [FORTUNE_CSV] [REDIS_HOST]" 1>&2
    exit 1
fi

load-to-redis() {

	echo "Loading: $FORTUNE_CSV , to Redis endpoint: $REDIS_HOST"
	{
        read # fields are QUOTE;AUTHOR;GENRE. Skip header
        INDEX_NUM=0
        while IFS=";" read -r QUOTE AUTHOR GENRE
        do 
            INDEX_NUM=$((INDEX_NUM+1))
            # escape single quote '\'' and double quote \", lets try html escaping.
            QUOTE_ESC=`echo $QUOTE | sed 's/'\''/\&apos;/g' | sed 's/'\"'/\&quot;/g'`
            echo "$INDEX_NUM,  $GENRE, $AUTHOR, $QUOTE_ESC"
            redis-cli -h $REDIS_HOST HSET quote:$INDEX_NUM genre "$GENRE" author "$AUTHOR" quote "$QUOTE_ESC"
		done
	} < $FORTUNE_CSV
    # Show number of items.
    redis-cli -h $REDIS_HOST info keyspace
    redis-cli -h $REDIS_HOST dbsize
}

load-to-redis