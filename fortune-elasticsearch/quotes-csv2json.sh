#!/bin/bash
# Convert Quotes CSV to JSON for Elasticsearch import. 
# Runs faster on EC2 than Macbook.

FORTUNE_CSV=./quotes.csv
#FORTUNE_CSV=./quotes10.csv
FORTUNE_JSON=./quotes.json
ES_INDEX_NAME=quotes

# convert csv into 
convert-csv-to-json() {

	echo "Converting $FORTUNE_CSV to json..."
	{
		read # fields are QUOTE;AUTHOR;GENRE. Skip header
		INDEX_NUM=0
		while IFS=";" read -r QUOTE AUTHOR GENRE
		do 
			INDEX_NUM=$((INDEX_NUM+1))
			echo '{ "index" : { "_index": "'$ES_INDEX_NAME'", "_id" : "'$INDEX_NUM'" } }'
			# escape single quote and double quote chars with html entity, escape the ampersand.
			QUOTE_ESC=`echo $QUOTE | sed 's/'\''/\\\&apos;/g' | sed 's/'\"'/\\\&quot;/g'`
			echo '{"quote":"'$QUOTE_ESC'", "author":"'$AUTHOR'", "genre":"'$GENRE'"}'
		done
	} < $FORTUNE_CSV >> ${FORTUNE_JSON}
	
}

convert-csv-to-json