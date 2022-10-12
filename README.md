# DDD PATTERN

A simple pattern for golang development

## Description

Domain-driven design (DDD) is a major software design approach, focusing on modelling software to match a domain according to input from that domain's experts

## Folder management

```
 >config    #configuration folder
    app.go  #app config
    mysql.go    #database config
    validator.go    #other config
 >src   #folder that load about this pattern
    >domain #domain folder, load an entities/model and repo
        >model  #save that entities of product/feature
            >book #feature / product entities
                book.go #model
                bookRequest.go  #entities for request
                bookResponse.go #entities for response
        >repository
            book.go #contract interfaces about your func
    >infrastructure #function that handle from another bussines logic/server
        >crm-core   #example another service
        >repository #repository code
            >mocks  #mocking for repository
                bookRepository.go   #example mocks
            bookRepository.go   #example repository
        >utilities  #data needle for common that call globally
            commonResponse.go   #example
    >interfaces gateway to bussiness
        >command    gateway from command
            >migration  #ecample migrate
            >worker #example for worker
        >http   #http request
            >handler    #handler gateway
                >v1 #versioning folder
                    bookHandler.go  #handler 
                >v2
        >routes #routes folder
            v1.go
            v2.go
    >services
        bookServices.go
.gitignore
config.json
main.go
README.md
go.mod
go.sum        
```

## Authors

Contributors names and contact info

- Ahmad Yusuf Ardabilli  
[@email](billya749@gmail.com)
