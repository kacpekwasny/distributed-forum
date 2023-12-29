- what are you tinking abut?
- im tinking abaut desajn




Założenia:
- dystrybuowane,
- obciążenie dzieli
- distributed, no single server is responsible for the failure of all

Important points:
- UNDO is not possible, no deleting posts, no removing reactions, no editing comments, no nothing, be tough and think twice,
- copies on other servers,
- distributed
  - no single server has all data,
  - traffic may handled by many servers (if one server is an owner of a very popular Age, clients may browse it through many )




To Do:
- UniverseIface(historyUri string)
- PeerManager
- Hot Posts, Top Posts (ale to wsm jak w order dziala)
- Write the NoUndo Iface and struct for dependency injection
- try out the gh.com/saber/mo?
- 


- notifications! It is pretty important to be able to follow up some threads on you comments.
  - since we want notification, we have to keep track of last seen state, and the change,
  - my proposed solution is:

  Merge all postable content into a single table, stories and comments:
  ```
  Postable
  -------------
  postable_id           // merge stories and posts into a single table, and the difference would be posts dont have a parent
  ...
  interactions  int     // this much easier, that you cannot unclick reaction, cannot unpost comment,
  ```

  a table storing postable content, that a user is following, with number of interactions since last clicked on this notification and current number of interactions,
  interactions would be comments and reactions,
  ```
  Followed Postable
  ------------------------------
  user_id                   int
  postable_id               int
  last_seen_interactions    int
  ```


- notifications:
  - for every postable in notifications get info on current state of interactions
    - on parent server:
      - if postable on self -> check self
      - if on remote

- i was thinking about connections to peeres through http2, considering the authentiaction, posting and restricted acces for reading from API - well, reading of some information has to be restricted, but especialy posting.
A good idea might be creating an interface, and then changing out the inner workings.

