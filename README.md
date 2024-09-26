# Eth Fetcher Go App

This application is an Ethereum transaction fetcher and person info manager built with Go. It provides endpoints for fetching Ethereum transactions, managing user authentication, and interacting with a smart contract.

## Setup

### Prerequisites

- Go 1.22 or later
- PostgreSQL
- Homebrew (for macOS users)
- Recommended test net - Base Sepolia
- Example Simple Person Info Contract is deployed on Base Sepolia with address 
`0xf321e3770293Bbb920032C5501Cd9A64b223bB9c`

### Environment Variables

Create a `.env` file in the root directory following the example in the `.env.example` file:

```
API_PORT=4000
DB_CONNECTION_URL=postgres://username:password@localhost:5432/database_name?sslmode=disable
JWT_SECRET=your_jwt_secret

ETH_NODE_URL=https://your-ethereum-node-url
ETH_SOCKET_URL=wss://your-ethereum-websocket-url
PRIVATE_KEY=your_private_key
SIMPLE_PERSON_INFO_CONTRACT_ADDRESS=
```

Replace the placeholder values with your actual configuration.

### Database Setup

If you're using macOS and Homebrew, follow these steps to set up PostgreSQL:

1. Install PostgreSQL:

   ```
   brew install postgresql
   ```

2. Start the PostgreSQL service:

   ```
   brew services start postgresql
   ```

3. Create a new database:

   ```
   createdb eth_fetcher_go
   ```

4. Update the `DB_CONNECTION_URL` in your `.env` file with the correct credentials and database name.

## Running the Application

1. Clone the repository:

   ```
   git clone https://github.com/ddzhalevlime/Etherium-fetcher-Go.git
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Run the application:
   ```
   go run . //run from within cmd/web dir
   ```

## Notes:

- The server will start on the port specified in your `.env` file
- The server will automatically create the `personInfoEvents`, `transactions`, `users` tables upon launch
- The `users` table will be auto populated with 4 users with the following username/password pairs

- `alice`/ `alice`
- `bob`/ `bob`
- `carol` / `carol`
- `dave`/ `dave`

## Running Tests

To run the tests, use the following command in the project root:

```
go test ./...
```

## API Endpoints

### 1. Get Ethereum Transactions By Hash

- **GET** `/lime/eth`
- **Query Parameters**: `transactionHashes` (separated by `&transactionHashes=`)
- **Headers**: `AUTH_TOKEN: <token>` (OPTIONAL)
- **Example Request**:
  ```
  GET /lime/eth?transactionHashes=0x123&transactionHashes0x456
  ```
- **Example Response**:

  ```json
  {
    "transactions": [
      {
        "transactionHash": "0x123...",
        "transactionStatus": 1,
        "blockHash": "0xdd...",
        "blockNumber": 15746162,
        "from": "0xabc...",
        "to": "0xdef...",
        "contractAddress": "",
        "logsCount": 4,
        "input": "1234...",
        "value": "10000"
      }
    ]
  }
  ```

  ### 2. Get Ethereum Transactions By RLP encoded list of hashes

  example data on Base Sepolia

  Raw Data

  ```
  [“0x86716be20c708b1664e74f41e6e0dd2b880efb3bae242a9dc7307af8dcb62a19","0x7dc054db23ee31c836ae8f4ad7d2c1bbe1d4c43115bb645e006fde52b15a53ab"]
  ```

  RLP Hex Produced by Raw Data

  ```
  f888b842307838363731366265323063373038623136363465373466343165366530646432623838306566623362616532343261396463373330376166386463623632613139b842307837646330353464623233656533316338333661653866346164376432633162626531643463343331313562623634356530303666646535326231356135336162
  ```

- **GET** `/lime/eth/{rlphex}`
- **Path Parameters**: `rlphex`
- **Headers**: `AUTH_TOKEN: <token>` (OPTIONAL)
- **Example Request**:

```

GET /lime/eth/f888b84230783836373136626532306337303862....

```

- **Example Response**:

```json
{
  "transactions": [
    {
      "transactionHash": "0x123...",
      "transactionStatus": 1,
      "blockHash": "0xdd...",
      "blockNumber": 15746162,
      "from": "0xabc...",
      "to": "0xdef...",
      "contractAddress": "",
      "logsCount": 4,
      "input": "1234...",
      "value": "10000"
    }
  ]
}
```

### 3. Get All Transactions

- **GET** `/lime/all`
- **Example Response**:
  ```json
  {
    "transactions": [
      {
        "transactionHash": "0x123...",
        "transactionStatus": 1,
        "blockHash": "0xdd...",
        "blockNumber": 15746162,
        "from": "0xabc...",
        "to": "0xdef...",
        "contractAddress": "",
        "logsCount": 4,
        "input": "1234...",
        "value": "10000"
      }
    ]
  }
  ```

### 4. Authenticate User

Note: by default the application

- **POST** `/lime/authenticate`
- **Example Request**:
  ```json
  {
    "username": "alice",
    "password": "alice"
  }
  ```
- **Example Response**:
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
  ```

### 5. Get User's Transactions

- **GET** `/lime/my`
- **Headers**: `AUTH_TOKEN: <token>`
- **Example Response**:
  ```json
  {
    "transactions": [
      {
        "transactionHash": "0x123...",
        "transactionStatus": 1,
        "blockHash": "0xdd...",
        "blockNumber": 15746162,
        "from": "0xabc...",
        "to": "0xdef...",
        "contractAddress": "",
        "logsCount": 4,
        "input": "1234...",
        "value": "10000"
      }
    ]
  }
  ```

### 6. Save Person Info

Returns `txStatus` `true` for success and `false` for fail

- **POST** `/lime/savePerson`
- **Headers**: `AUTH_TOKEN: <token>`
- **Example Request**:
  ```json
  {
    "name": "John Doe",
    "age": 30
  }
  ```
- **Example Response**:
  ```json
  {
    "txHash": "0xabc...",
    "txStatus": "true"
  }
  ```

### 7. List Persons

- **GET** `/lime/listPersons`
- **Example Response**:
  ```json
  {
    "persons": [
      {
        "id": 1,
        "personIndex": 20,
        "personName": "Jon Doe",
        "personAge": 50,
        "TransactionHash": "0x123...."
      }
    ]
  }
  ```
