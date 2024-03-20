PDF Generator Project
=====

Quick Start
-----

To run this project we need to have `golang` and `nodejs` (`yarn`) installed.

Backend
-----

Under the rood directory of the project run the following commands:

1. `golang build -o ./api/app ./api`
2. `golang build -o ./daemon/app ./daemon`
3. `docker network create torrens`
4. `docker compose up -d db`
5. Check `docker ps` to ensure that `db` service is running
6. Connect to `mysql` at localhost:3306
7. Create a database `sep401`
8. Create tables from `db.sql` in the created database
9. `docker compose up -d`
10. Check `docker ps` to ensure that five containers are running (`app`, `daemon`, `db`, `chromiun`, `smtp`)
11. Check that http://localhost:8080/v1/user/balance returns message `not authorized`

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

