# POST
$postParams = @{name='name';host='localhost';port='5672';user='guest';pass='guest';duarble='true';usessl='true'}
Invoke-WebRequest -Uri http://localhost/api/v1/create_connection -Method POST -Body $postParams

$msgParam = @{id=1;body='my text message'}
Invoke-WebRequest -Uri http://localhost/api/v1/send_message -Method POST -Body $msgParam

# DELETE
Invoke-WebRequest -Uri http://localhost/api/v1/delete_connection/1 -Method DELETE
