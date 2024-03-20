PDF Generator Project
=====

Quick Start
-----

To run this project we need to have `golang` and `nodejs` (`yarn 1`) installed.

Backend
-----

Under the rood directory of the project run the following commands:

1. `go build -o ./api/app ./api`
2. `go build -o ./daemon/app ./daemon`
3. `docker network create torrens`
4. `docker compose up -d`
5. Check `docker ps` to ensure that five containers are up and running (`app`, `daemon`, `db`, `chromiun`, `smtp`)
6. `docker exec -it db sh` and run `mysql -u $MYSQL_USER -p"$MYSQL_PASSWORD" -D $MYSQL_DATABASE < /opt/db.sql` inside the container
7. Check that http://localhost:8080/v1/user/balance returns message `not authorized`

Frontend
-----

Under the rood directory of the project run the following commands:

1. `cd frontend`
2. `yarn install`
3. `yarn serve`
4. Navigate to http://localhost:8081 to ensure that frontend is working

Shutdown

1. Terminate `yarn serve` by pressing `Ctrl+C
2. Under the rood directory of the project run `docker-compose down`

