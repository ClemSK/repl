# Repl

CLI app in Go that fetches data from an API and displays it in the terminal.

### Learning objectives

- Parsing JSON in Go from a public API
- Making HTTP requests in Go using the net/http standard library package
- Building a CLI tool that makes interacting with a back-end server easier
- Caching and how to use it to improve performance

### API data

The project uses the Locations and Pokemon endpoints of the [PokeApi](https://pokeapi.co/docs/v2#pokemon-section)

### Repl commands

| Commands                | Description                                           |
| ----------------------- | ----------------------------------------------------- |
| map                     | Lists the next page of location areas                 |
| mapb                    | Lists the previous page of location areas             |
| explore {location_area} | Lists the pokemons in a location areas                |
| catch {pokemon_name}    | Attempt to catch a pokemon and add it to your pokedex |
| inspect {pokemon_name}  | View info about a caught pokemon                      |
| pokedex                 | View all the pokemon in your pokedex                  |
| help                    | Displays a help message and lists available commands  |
| exit                    | Exit the pokedex                                      |

### To run

1. Install Go 1.21 or later for your machine (latest stable version)
2. Clone the repo locally
3. Run: go build && ./your-local-directory
4. Fun!
