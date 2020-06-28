# Test API

Run Project in local machine
 
  ```./run.sh```



## Endpoints

1. Fetch all Persons 

    **URI**: ``/api/v1/persons``, **METHOD**: ``GET``
1. Fetch a Person 
    
      **URI**: ``/api/v1/persons/<id>``, **METHOD**: ``GET``
1. Delete a Person
 
    **URI**: ``/api/v1/persons/<id>``, **METHOD**: ``DELETE``
1. Update a Person 
        
      **URI**: ``/api/v1/persons/<id>``, **METHOD**: ``PUT``
1. Create a Person
    
    **URI**: ``/api/v1/persons``, **METHOD**: ``POST``

## Note

1. Run the SQL migration under db->schema folder
1. Used docker-compose for local development