Project README
Project Name: Gofr-api

Book Store Management System with CRUD Operations:
  This user-friendly system streamlines book store operations by managing book information and transactions efficiently.

Technologies:

-Golang
-Gofr web framework
-CompileDaemon for automated recompilation
-MySQL database
-Docker for containerization and deployment

Installation:

Clone the repository.
-Build the Docker image: docker-compose build.
-Run the application: docker-compose up.

Configuration:

-Environment variables can be used to configure the application and database connection settings.
Refer to the entrypoint.sh script for configuration options.
Running the application:

The application exposes port 8080 by default.
You can access the application at http://localhost:8080.

API Endpoints
  GET /books: Retrieve all movies.
  GET /books/{id}: Retrieve details of a specific movie.
  POST /books: Create a new movie.
  PUT /books/{id}: Update details of a specific movie.
  DELETE /books/{id}: Delete a movie.


Testing:

  Unit tests are included in the project.
  You can run the tests with: go test -v ./...
Deployment:

  The Docker image can be deployed to any platform that supports Docker containers.
  Consider using container orchestration platforms like Kubernetes for advanced deployment scenarios.
  Contributing:


Postman collection for trying out the APIs:
  <img width="1227" alt="Screenshot 2023-12-18 at 2 01 09 AM" src="https://github.com/ishaan5470/gofr-api/assets/124041853/6590cc39-2dbf-40f0-90ac-923ecf47cc61">


  UML Diagram:
    Project Structure:
      <img width="789" alt="Screenshot 2023-12-18 at 2 29 50 AM" src="https://github.com/ishaan5470/gofr-api/assets/124041853/cd0f335e-c706-41b6-9110-dc0acb1c1955">
      


  Sequence Diagram: (Considering one example only)
    +-------------------+          +---------------------+          +---------------------+          +-----------------------+
|      User         |          |      HTTP Server    |          |    UpdateBook()     |          |       Database        |
+-------------------+          +---------------------+          +---------------------+          +-----------------------+
       |                             |                               |                               |
       | 1. HTTP PUT /books/{id}      |                               |                               |
       |---------------------------> |                               |                               |
       |                             |                               |                               |
       |                             |                               |                               |
       |                             |                               |                               |
       |                             |                               |                               |
       |                             |                               |                               |
       |                             |                               |                               |
       |                             |                               |                               |
       |                             |                               |                               |
       |                             | 2. Parse request body         |                               |
       |                             |---------------------------> |                               |
       |                             |                               |                               |
       |                             |                               |                               |
       |                             | 3. Retrieve book details     |                               |
       |                             |    by ID                     |                               |
       |                             |---------------------------> |                               |
       |                             |                               |                               |
       |                             |                               |                               |
       |                             | 4. Update book details       |                               |
       |                             |    based on request          |                               |
       |                             |    and save to database      |                               |
       |                             |---------------------------> |                               |
       |                             |                               |                               |
       |                             |                               |                               |
       |                             | 5. Marshal updated book      |                               |
       |                             |    details to JSON           |                               |
       |                             |<--------------------------- |                               |
       |                             |                               |                               |
       |                             | 6. Set HTTP headers and      |                               |
       |                             |    respond with JSON         |                               |
       |                             |---------------------------> |                               |
       |                             |                               |                               |
       |                             |                               |                               |
       |                             |                               |                               |
       |                             |                               |                               |





















  

We welcome contributions! Pull requests are encouraged.
Contact: ishaan8248@gmail.com




