# POST
$postParams = @{name='name';host='host';port='123';user='user';pass='pass';duarble='dur';usessl='true'}
Invoke-WebRequest -Uri http://localhost/api/v1/create_connection -Method POST -Body $postParams


# DELETE
Invoke-WebRequest -Uri http://localhost/api/v1/delete_connection/1 -Method DELETE
