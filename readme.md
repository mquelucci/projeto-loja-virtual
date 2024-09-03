## ENVIROMENT FILE - CONFIG.ENV

**TYPE** : *[sqlite3 / postgres / mysql]* Indica se o banco a ser usado é o SQLite3 (local, dentro do diretório raiz), ou postgres/mysql em nuvem ou servidor dedicado.

**FILEDATABASE** : Nome do banco de dados sqlite3 (deve incluir a extensão)

**CONNECTIONSTRING** : String necessária para conexão com banco de dados postgres/mysql em servidor ou em nuvem.

## API

### AUTENTICAÇÃO

- **[ POST ]** - **Autenticar** - Realiza a autenticação do usuário para interface administrativa
    - _Endpoint_: `/admin/autenticar`
    - **Parâmetros**:
        - **Tipo:** Multipart Form Data
        - Usuário (_string_)
        - Senha (_string_)
        