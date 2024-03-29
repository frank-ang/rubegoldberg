lookup-config-db() {
	echo "Looking up database config from secretsmanager..."
	SECRET_JSON=`aws secretsmanager get-secret-value --secret-id ${DB_SECRET_NAME} --query 'SecretString' --output text`
	DB_USERNAME=`echo $SECRET_JSON | jq -r '.username'`
	DB_PASSWORD=`echo $SECRET_JSON | jq -r '.password'`
	DB_HOST=`echo $SECRET_JSON | jq -r '.host'`
}