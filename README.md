# dba-agent

A lightweight app to execute various SQL queries against various SQL databases on various schedules.
Initial intent was to run various [MaintenanceSolution commands from Ola Hallengren](https://ola.hallengren.com/) when there's no SQL Agent available (like in Azure SQL Database) or when you want your scheduling across many servers to be centralized.

## Build and Execute

- `go build -o ./build/dba-agent .`
- `./build/dba-agent`

## Required Environment Variable Configuration

Database connection strings should be stored as Environment Variables.  Each named variable needs to be referenced in the `config.toml`

Here's an example:
`Localhost_DBCONN="server=localhost;user id=sa;password=TestPassword;port=1433;database=MyDb"`

## Config Adjustments

- This app currently is driven by the config.toml file. The included file just has placeholder values that should be modified to meet your needs.
- The main requirements are:
    1. `[[db_servers]]` section containing "name" and "conn_string_variable" values.
    1. `[[db_queries]]` section containing "name", "query", "schedule", and "server".
- The "server" value in db_queries needs to reference the "name" value from one of the db_servers.
