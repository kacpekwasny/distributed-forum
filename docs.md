# noundo - documentation

`HistoryIface` - interface to a **history**, access its contents, is used to connect to local database as well as to connect to a remote server via gRPC,
It has methods for reading, writing, creating, contents.

- `pkg/peer/` - holds `historyServer.proto` file, that generates other gRPC code for golang, also some methods additionaly implemented in file `methods.go` (this file is not overwriten during compilation of .proto file),

- `pkg/noundo/` - holds rest of the code (needs refactoring into smaller parts),

- `pkg/noundu/templates_pages` - `html` files for go Templating engine.
  - `base.go.html` - holds header, navigation and footer - things that usualy dont change, when we change the view. But sometimes they do have to change. Base defines scripts, that can be accessed globaly by any other template. These scripts are used to control the state of ex. navbar.

  - `page_*.go.html` - a 'view', that extends `base.go.html` template, when requested by Htmx it is returned without base rendered,

  - `comp_*.go.html` - componenets, that are always returned standalone, always without base,

- `NoUndo` - struture holding:

  - `Universe` - structure that contains:
    - self - our history,
    - peersNexus - interface to other histories,

  - everything required by `http server`,

  - everything required by `grpc server`,

- `noundo_handle_ABC.go` - files with http handler methods defined for `NoUndo` - thus they have access to everything like: `Self() HistoryIface`, access to peers `HistoryIface`. Those methods use `ExecTemplHtmxSensitive` - this function executes templates with respect to the method of request. If the request was initialized by Htmx - then the template is returned without adding base. If the request comes from a refresh of page, then the whole page must be returned with the `base`.

- `peer_*.go` files for making a peer connection, the server, the client, the wrappers for grpc server and client, wrappers for the grpc structures,

- `auth_*.go` files for authenticator storage iface, authenticator middleware (it checks for a valid JWT in incoming request, and if so it sets a `JWTFields` in `request.Context()`),






