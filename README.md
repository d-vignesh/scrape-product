Product scraping service built using golang. The repo consists of three services namely scraper, store and database.

Installations needed to test the service locally:
    1. git
    2. docker
    3. docker-compse

Steps to build and test the service:
    1. clone the repo from the github(https://github.com/d-vignesh/scrape-product)
    2. navigate to the project root directory where docker-compose.yml file exist (scrape-product in this case)
    3. Bring up all the services using the command
            docker-compose up
       (Note: the scraper, store and database services are configured to run on ports 9001, 9002 and 5432 repectively. ensure these ports are free, or change the ports according in the docker-compose environment variables)
    4. click on the button to get all the endpoints in postman <br/>[![Run in Postman](https://run.pstmn.io/button.svg)](https://god.postman.co/run-collection/11a39d0cbb511338e62b)<br/> 
        there are three endpoint,
            /scrape-product?url - to scrape the product from url
            /store-product - to store the given product to db
            /get-products - to get the list of all products in db

Modules used:
    * go-colly (for scraping)
    * go-hclog (for logging)
    * gorilla/mux (for routing)
    * sqlx (for postgres db)