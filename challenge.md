# Backend Developer Assignment

## Task details

- The expected duration for completion of the task is **from 3 to 5 days**
- Please deploy your code on Netlify or Vercel
- Programming Language to use: Golang or Rust. The task is also to get you familiar with Golang if you have not previously worked with it.

---

### Minimum Requirements

**Data Acquisition:** Use an existing API like CoinGecko or CryptoCompare to get exchange rates for various cryptocurrencies (Bitcoin, Ethereum, Litecoin, etc.) against a selection of fiat currencies (USD, EUR, GBP, etc.). You should update this data periodically (for example, every 5 minutes).

**Database:** Implement a database to store the exchange rate data. This could be an SQL or NoSQL database, depending on your preference. You should design the database schema to efficiently store and retrieve the exchange rate data.

**API Endpoints:** Develop the following RESTful API endpoints:

API Endpoints:

- GET /rates/{cryptocurrency}/{fiat}: Returns the current exchange rate between the specified cryptocurrency and fiat currency.
- GET /rates/{cryptocurrency}: Returns the current exchange rates between the specified cryptocurrency and all supported fiat currencies.
- GET /rates: Returns the current exchange rates for all supported cryptocurrency-fiat pairs.
- GET /rates/history/{cryptocurrency}/{fiat}: Returns the exchange rate history between the specified cryptocurrency and fiat currency for the past 24 hours.

Web3 Integration:
Use the web3 library to retrieve the current balance of a specific Ethereum address. Expose this functionality through an API endpoint such as GET /balance/{address}. Note that no smart contract creation is necessary; you only need to interact with the Ethereum blockchain.

**Error Handling and Validation:** Your API should validate request data and return appropriate error messages in case of invalid requests. Additionally, it should handle potential errors or exceptions gracefully.

**Testing:** Write unit and integration tests to verify the correctness of your API.

### **Expected Output**

Source code for the API, stored in a public or private repository on a version control system like Git.
Documentation on how to run and use your API. This should include instructions for setting up the database and any necessary environment variables, as well as how to start the server.
Postman collection or equivalent for testing all the developed APIs.

### Challenges to Overcome

Gathering data from external APIs in an efficient manner.
Designing a database schema for efficient storage and retrieval of exchange rate data.
Building a robust API with proper validation and error handling.
Interacting with the Ethereum blockchain to get an address's balance.
Writing comprehensive tests for your application.
