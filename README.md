# create2-payment-gateway

Proof of concept for a payment gateway using the CREATE2 opcode.
Addresses are deterministically generated based on the sender address and a unique identifier.

The server backend checks for the generated address balance periodically and triggers the
drain contract deployment when the balance is greater than the threshold.

This method is useful for creating a payment gateway without the need to
handle private keys or seed phrases and only wasting resources when a payment is made.
