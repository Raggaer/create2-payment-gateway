# create2-payment-gateway

Proof of concept for a payment gateway using the CREATE2 opcode.
Addresses are deterministically generated based on the sender address and a unique identifier.

The server backend checks for the generated address balance periodically and triggers the
drain contract deployment when the balance is greater than the threshold.

This method is useful for creating a payment gateway without the need to
handle private keys or seed phrases and only wasting resources when a payment is made (deployment gas cost).

## Compile deployer

Compile the deployer and generate the go bindings. Place the generated file in the `backend/deployer` folder.

```bash
abigen --sol ContractFactory.sol --pkg deployer --out deployer.go
```

## Setup

Sample .env file for the backend:

```bash
LISTEN_ADDR=":8080"

ETH_NODE="http://localhost:8545"

FACTORY_ADDR="0xCa35387AC0f318404B72Be39dFbfC373E0E1877f"
FACTORY_PK="0x00"

OWNER_ADDR="0x941b74b4943a982691E7f4A120E759E4d8eA72fa"

DB_USER="root"
DB_PASS="admin"
DB_NAME="create2_payment_gateway"
DB_HOST="localhost"
DB_PORT=3306
```

Deploy the factory contract and set the `FACTORY_ADDR` and `FACTORY_PK` in the .env file.
