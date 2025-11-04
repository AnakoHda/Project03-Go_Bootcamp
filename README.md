- Task 0. Project Initialization
You need to create a project that will be used for all subsequent tasks.
- Task 1. Project Structure
Each layer should be its own module.
The project structure must follow the Standard Go Project Layout.
Separate the API contracts layer for client interaction.
Separate the application layer for business logic implementation.
Separate the infrastructure layer for handling data operations (e.g., with a database).
Separate the DI (dependency injection) layer, where configurations for dependency injection are described (e.g., using the uber/fx library).
- Task 2. Implement the Domain Layer
Define a game board model as an integer matrix.
Define a current game model, which should include a UUID and the game board.
Define a service interface with the following methods:
A method to compute the next move of the current game using the Minimax algorithm.
A method to validate the game board of the current game (check that previous moves haven't been altered).
A method to check for game completion.
Models, interfaces, and implementations must be placed in separate files.
- Task 3. Implement the Datasource Layer
Implement a storage structure to keep track of ongoing games.
Use thread-safe collections (e.g., sync.Map?) for storing data.
Define models for the current game’s game board.
Implement mappers between domain and datasource layers (domain <-> datasource).
Implement a repository to interact with the storage structure. It must provide the following methods:
A method to save the current game.
A method to retrieve the current game.
Create a structure that implements the service interface and accepts a repository interface as a parameter to work with the storage structure.
Models, interfaces, and implementations must be placed in separate files.
- Task 4. Implement the Web Layer
Define models for the current game’s game board.
Implement mappers between the domain and web layers (domain <-> web).
Implement a handler using net/http, with a method:
POST /game/{current_game_UUID} — sends the current game with the user’s updated game board and returns the current game with the computer’s updated game board.
If an invalid game with an incorrect updated board is sent, an error with a description must be returned.
The application must support multiple games simultaneously.
Models, interfaces, and implementations must be placed in separate files.
- Task 5. Implement the DI Layer
Use the uber/fx library to define a dependency injection graph.
The graph must include at least the following components:
The storage structure, registered as a singleton;
The repository for working with the storage structure;
The service for working with the repository.