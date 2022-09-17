# Task App

Guides to setup the app locally.

1. Create a database with a name `task_app_dev` in postgresql server. You can change the database name to whichever you want and just need to update the `DATABASE_NAME` value in `.env` file.
2. Eventually we will not commit the `.env` into the repo, this is just for demostration.
3. Update the `DATABASE_USER` and `DATABASE_PASS` for the database connection access in the `.env` file.
4. There is a config file named `dbconfig.yml` under `config` folder also needed to update the connection setting accordingly.  
5. Run the following command in your terminal to start the backend server: `make  start_api`
6. Open another terminal and run the following command to start the front-end server: `make start_fe`
7. Open `http://localhost:8080` in your browser and you will see the homepage of the app.

**Front-end**

I'm using Svelte as the front-end framework which I can learn quicky (first time) and start to develop a simple working prototype as I dont have any experience in ReactJS. The business logic are located in the files under `src` folder.

**Back-end**

The apis in backend is developed using [Fiber](https://docs.gofiber.io/), an Express inspired web framework built on top of Fasthttp, a HTTP engine for GO and [upper/db](https://upper.io/v4/) as the database access layer to postgresql server.

**Database migration**

[sql-migrate](https://github.com/rubenv/sql-migrate) is a SQL Schema migration tool for Go, all the migration SQL files are located under `migrations/postgresql` folder.

**Should Have user stories**

1. *Sort by due date or create date.*
   To achieve this goal, we will need to store the state to indicate the direction (ascending or descending) whenever user click on the sorting icon, update the value in the store and append the direction in the payload of the request to the API to get the result.

2. *Search based on task name.*
   Append the text that entered by user to a query payload in the request and send to the API which is able to update the SQL statement to query in the table using "ILIKE" operator with case-insensitive pattern matching format.

**Risks**

1. *There is a rick that a user may create more than 1000 tasks which might caused some server performance issue.*
   One of the way to resolve it is to paginate all those records into multiple pages. We may set a ideal number of records to show in the page depend on the UI design. Let's say the optimal number of tasks to show in a page is 50, 1000 / 50 should have 20 pages, user can click on the page number link or next page link in the pagination bar and FE will send the page/offset number in the payload to the API to get the result.

**Todo**

- The pagination in FE still need some work to get the page/offset number to send in the payload of the request.