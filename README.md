# noundo

### Redistributed 
interconnected servers, every server can be accessed by any other server.

### Nothing can disapear
peers permanently mutualy cache each others content.

### No action can be undone
the cached content doesn't expier, doesnt need to be refreshed.

The forums structure was based on reddit.
- history - a server
    - age - a subreddit
        - story - a post in an age
            - answer - a comment to a story, also answers to answers are possible
    - users - a history has users,


## Progress

### Done:
- User SignUp, SignIn,
- HomePage,
- BrowsingPeers (gRPC connection),
- Browsing Ages,
- Browsing Stories,
- Browsing Answers,
- Adding Stories,
- Adding Answers,

### ToDo:
- creating Ages (frontend),
- signIn to account created on History1, while browsing History2,
- caching content retrieved from peers,
- reactions (frontend, backend),
- a lot of refactoring,


## Interactions with other histories:

### Currently:
- visit: http://history1.com (this is an example link, not a real one),
- you sign up
- you sign in,
- lets say history1.com is peered with history2,
- while beeing on http://history1.com in you browser you may browse ages from history2 (and vice versa),
- you can go to a link: http://history1.com/a/history2.com/age2,
- and then create a story,
- from now on, this story is available on http://history2.com,


### Intended:
- you may add stories to other histories if the history you are browsing from is peered with that history 
  you want to post to.
- Example: on your browser you have opened history `brows.com`, your user is registered @ `reg.com`,  you 
  are browsing an *age* from `age.com`. You may post a *story* to that *age* if `brows.com <-> age.com` and `age.com <-> reg.com`, where `<->` meens a peer connection.
- if you browse `a.com` from `brows.com`, then every post, answer coming retrived from `a.com` will be saved at `brows.com`. Thanks to **no-undo** policy, those posts never expire, and will never be deleted from the internet.

## Running your own history

To run your own history you should configure the code in main.go, eg. specify the name etc. and create the admin
account.





  


