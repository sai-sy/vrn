# ideas
## Schema design questions
### How should I store tags
1. A table with a tag column, a type column (person, address... whatever you are applying the tag to), an ID column
2. model schema has a tag column
    - a list
        - makes sense and simple
        - not resizable
    - a comma seperated string
        - simple data type
        - schemaless could be a problem
    - a dictionary/hash/objects/struct

### How to store memberships information
I think we would want to store it as key: value pairs with the key being the thing they are a member of and the value being the last year that we know they signed up/were a member

### How to design membership information scheme
1. Each Person can have inididual collums for each type of membership (YLC, LPC, CPC.. ETC)
    - obviouss solution
    - migrating to include more or less clubs is stupid
2. Each Person can have a column that holds a json
    - standardizing a shemaless data form is stupid
3. Membership Table where we have club names and year and PersonID
    - a whole other table for this seems excessive
    - still might be cool. 
    - could use this to populate idea #2
