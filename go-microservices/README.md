# go-microservices

#### Concepts learned in this project:
- Reflection 
- Middleware functions
- Writing custom validations with Govalidator 
- Go routines 
- Documenting with Swagger:
    - Creating a spec as swagger.yml
        - http://goswagger.io/use/spec/ 
            - Comment the handlers using swagger syntax 
            - Initiliaze the spec: swagger init spec 
            - Generate the spec: swagger generate spec -o ./swagger.yml --scan-models
    - Serving the spec in the API path /docs with Redoc
    ![image](https://user-images.githubusercontent.com/109003970/231711286-1ea9c184-c342-4183-8cd4-314acbbfaa80.png)
- Generating a client with Swagger (see folder sdk):
    - swagger generate client -f ../swagger.yml -A product-api 
- CORS: Allow calls to be made by origins other than the default origin.
- Multipart requests 
- Gzip compression for HTTP responses (reduce data size vs a little CPU consumption)
