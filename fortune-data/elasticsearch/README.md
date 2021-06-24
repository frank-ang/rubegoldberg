# Redis ElasticSearch Readme.

## Create ElasticSearch Domain.

on AWS Console.

## Test connectivity

From a connected bastion host.

```bash
ES_HOST=https://vpc-sandbox00-7zz2l5gm2ihbkrgwsumu2bcpmu.ap-southeast-1.es.amazonaws.com

curl $ES_HOST
curl $ES_HOST/_plugin/kibana/
```

## Load test data.

* Example loading simple docs

    ```bash
    curl -k -XPUT $ES_HOST/movies/_doc/1 -d '{"director": "Burton, Tim", "genre": ["Comedy","Sci-Fi"], "year": 1996, "actor": ["Jack Nicholson","Pierce Brosnan","Sarah Jessica Parker"], "title": "Mars Attacks!"}' -H 'Content-Type: application/json'

    curl -k -XGET "$ES_HOST/movies/_search?q=mars&pretty=true"

    curl -k -XPUT $ES_HOST/quotes/_doc/0 -d '{"quote":"Age is an issue of mind over matter. If you don&apos;t mind, it doesn&apos;t matter.", "author":"Mark Twain", "genre":"age"}' -H 'Content-Type: application/json'

    curl -k -XGET "$ES_HOST/quotes/_doc/0"

    curl -k -XGET "$ES_HOST/quotes/_search?q=age&pretty=true"
    ```

* Load quotes test data.

    Converted file already provided. SCP to a host and import.

    ```bash
    # create index
    curl -k -XPUT $ES_HOST/quotes
    # split files into smaller chunks, this has already been done (see zip file).
    head -k -50000 quotes.json > quotes0-50k.json
    cat quotes.json | head -100000 | tail -50000 > quotes50k-100k.json
    tail -k -51932 quotes.json > quotes100k-151932.json
    head -k -50000 quotes.json > quotes0-50k.json

    # import!
    curl -k -XPOST $ES_HOST/_bulk --data-binary @quotes0-50k.json -H 'Content-Type: application/json'
    curl -k -XPOST $ES_HOST/_bulk --data-binary @quotes50k-100k.json -H 'Content-Type: application/json'
    curl -k -XPOST $ES_HOST/_bulk --data-binary @quotes100k-151932.json -H 'Content-Type: application/json'

    # verify
    curl -k -XGET "$ES_HOST/quotes/_doc/1"
    curl -k -XGET "$ES_HOST/quotes/_doc/75965"
    curl -k -XGET "$ES_HOST/quotes/_search?q=age&pretty=true"
    ```

## More queries.

```bash
# the following queriy forms are equivalent:
curl -k -X GET "$ES_HOST/quotes/_search?pretty=true&q=genre:age&size=1" -H 'Content-Type: application/json'

curl -k -X GET "$ES_HOST/quotes/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "size": 1,
    "query": {
        "term": {
            "genre": {
                "value": "age"
            }
        }
    }
}
'

# Get random doc
curl -k -X GET "$ES_HOST/quotes/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "size": 1,
    "query": {
        "function_score": {
            "query": { "match_all": {} },
            "functions": [ 
                {
                    "random_score": {} 
                }
            ]
        }
    }
}
'

# Get random doc, within a genre.
curl -k -X GET "$ES_HOST/quotes/_search?pretty" -H 'Content-Type: application/json' -d'
{
    "size": 1,
    "query": {
        "function_score": {
            "query": { "match": { "genre": "age" } },
            "functions": [ 
                {
                    "random_score": {} 
                }
            ]
        }
    }
}
'

```

## TODO: microservice


## TODO: setup LB to Kibana and Cognito (ABORTED)
