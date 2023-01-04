# Development Guide on Ubuntu 20.04

## Environment Preparation

### Docker

### Golang
- Install the golang from the official website [[REF]](https://go.dev/doc/install)
- Add the installed golang sdk to the PATH variable
    ```bash
    # for bash/zsh, to persist the change, add the following line to .bashrc or .zshrc file
    $ export PATH=$PATH:/usr/local/go/bin
    # for fish, to persist the change, add the following line to .bashrc or omf.fish file
    $ set PATH $PATH /usr/local/go/bin

    # then check the installation
    $ go version
    ```

### Database
- Install the PostgreSQL [[REF]](https://www.digitalocean.com/community/tutorials/how-to-install-postgresql-on-ubuntu-20-04-quickstart)
- Configure the user password of your postgres [[REF]](https://stackoverflow.com/questions/12720967/how-can-i-change-a-postgresql-user-password)
- Check the connectivity
    ```bash
    $ cd rigel_ledger_server
    # Create the env file of the project with the following variables:
    # - DB_HOST=localhost
    # - DB_PORT=5432
    # - DB_SCHEMA=public
    # - DB_NAME=rigel_ledger
    # - DB_USER=postgres
    # - DB_PW=
    $ touch .env.development
    $ bash shell/migrate.sh info
    ```


## Other References
- [Postgresql - Linux downloads (Ubuntu)](https://www.postgresql.org/download/linux/ubuntu/)
