## ENVIROMENT FILE - CONFIG.ENV

**TYPE** : *[sqlite3 / postgres / mysql]* Indica se o banco a ser usado é o SQLite3 (local, dentro do diretório raiz), ou postgres/mysql em nuvem ou servidor dedicado.

**FILEDATABASE** : Nome do banco de dados sqlite3 (deve incluir a extensão)

**CONNECTIONSTRING** : String necessária para conexão com banco de dados postgres/mysql em servidor ou em nuvem.

- **_MYSQL_** : {user}:{password}@tcp({address_or_servername}:{port_of_mysql})/{dbname}?charset=utf8mb4&parseTime=True&loc=Local

- **_POSTGRES_** : host={address_or_servername} user={username} password={password} dbname={database} port={port_of_postgres} sslmode={enable/disable} TimeZone={Country/Zone}

        