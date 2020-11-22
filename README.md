Product scraping service built using golang. The repo consists of three services namely scraper, store and database.<br/>

Installations needed to test the service locally:<br/>
    1. git<br/>
    2. docker<br/>
    3. docker-compse<br/>

Steps to build and test the service:<br/>
    1. clone the repo from the github(https://github.com/d-vignesh/scrape-product)<br/>
    2. navigate to the project root directory where docker-compose.yml file exist (scrape-product in this case)<br/>
    3. Bring up all the services using the command,<br/>
            docker-compose up<br/>
       (Note: the scraper, store and database services are configured to run on ports 9001, 9002 and 5432 repectively. ensure these ports are free, or change the ports according in the docker-compose environment variables)<br/>
    4. click on the button to get all the endpoints in postman <br/>[![Run in Postman](https://run.pstmn.io/button.svg)](https://god.postman.co/run-collection/11a39d0cbb511338e62b)<br/> 
        there are three endpoint,<br/>
            /scrape-product?url - to scrape the product from url<br/>
            /store-product - to store the given product to db<br/>
            /get-products - to get the list of all products in db<br/>

Modules used:<br/>
    * go-colly (for scraping)<br/>
    * go-hclog (for logging)<br/>
    * gorilla/mux (for routing)<br/>
    * sqlx (for postgres db)<br/>