## Project Structure:

**bin** - contains compiled binaries

**cmd/api** - app specific code for api app
running server, reading and writing http requests and manage auth
imports packages in the internal directory

**internal** - various auxilary packages. contains code for interacting with db,
validation, sending emails, etc.

    internal has a special meaning and behavior in Go. any packages
    which live under directory can only be imported by code inside
    the parent of the internal directory. nice.

**migrations**: sql migrations

**remote** - configuation files and setup scripts
