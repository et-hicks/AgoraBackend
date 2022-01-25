### Welcome to Agora Backend

There is a lot to do here so lets get started

#### The Completed
1. We have the first basic types down
   1. threads
   2. comments
   3. users
   4. reports

### Talking with Andrew
1. Follow a thread without signing in 
   1. Prompt them then to make an account for max ability 
   2. Follow a thread via email 
   3. Still no sign in process 
   4. Email/notify a post stat at the End of the day 
      1. Stats 
      2. Threads 
      3. Summary

#### Notes
1. JSON marshalling Go
2. Eventually we will need to build this and everything it comes with, 
but that comes later.


#### The TODOs

1. Need to write the SQL that will create the tables in the database
2. Need to transform the types into protos
   1. thread
   2. report (this can wait for a while)
   3. user
   4. comment
3. Need to write the functions that will put these types into the table
   1. Go Struct to Json
   2. Construct the Oracle-expecting json
   3. Call the Oracle REST endpoint
4. Need to write the functions that will get the data from the tables
   1. Call the oracle REST endpoint
   2. Marshal the Json to Go Struct
   3. return the struct from the fucntion
      1. return as a data pointer
5. Need to define the Event struct
   1. the event type
      1. comment
      2. thread created
      3. reported
   2. the event data
      1. 1 through 3 above
6. Need to create the event processor
   1. go source code and this might be the first actual logic worth thinking about
   2. needs to post to the firebase real time database to all the required microservices categories
7. need to create the users mention MS
   1. does the comment or post mention anyone?
8. need to create the Users involved with the thread MS
   1. MS to notify them in:
      1. their browser
      2. their email
      3. their summary
      4. etc
   2. and which type
9. need to create the event aggregation MS
   1. put the event in the DB
      1. store the comment
      2. store the post
10. Need to create the Data Collection MS
    1. likely many microservices here
    2. if a user was mentioned who mentioned them?
    3. if a user posted to a topic what was the topic?
    4. if people are talking,
       1. whos talking to who?
       2. who started the convo?
       3. whats the stats on this
       4. have these people talked with each other before?
       5. what are they talking about?
          1. keywords that keep popping up?
       6. etc.
    5. this is the users graph and this is big boy money $$$ that we need for funding and ad $$
11. Notify the people that are invited to contribute to a thread MS
12. Post Hashtags - later
    1. Need to make the hashtags not be virtue signaling
       1. nor dog whistling
    2. Make the hashtag a tagging fucntion
       1. referencing a thread
    3. hashtags as a niche in threads
    4. hashtag data agg
    5. hashtags are too broad and too varied to enable the feature of following hashtags
13. Post topic(s) - High priority
    1. users following the topic
    2. topic data aggregation
14. If a user posts a thread, notify the people that are following that user
    1. notify via:
       1. email?
       2. in browser notifications?
15. Create the code testing framework
    1. Every commit should have tests that it passes
    2. Need to run all tests before submitting PR
16. PR github actions
    1. need to set up dummy events and databases that can act as test passing
    2. need to create and think about a full testing suite before we merge to main
17. need to create the firebase infrastructure to auto-deploy microservices and updates to their firebase function
    1. terraform might help here
    2. github actions might also help here
18. Oracle
    1. need to create the oracle rest endpoints that actually run the SQL inserts, fetches, etc
    2. TODO: test that firebase functions can actually call the oracle rest endpoints 
       1. (super hope they can)
       2. (they 100% can cause its literally just REST)
19. 

