// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract ContractFactory {
  address private _owner;

  event Deployed(address addr);

  modifier onlyOwner() {
    require(msg.sender == _owner, "Only owner can call this function");
    _;
  }

  constructor(address owner) {
    _owner = owner;
  }

  function deploy(bytes memory code, bytes32 salt) onlyOwner external returns (address) {
    address ret;

    assembly {
      ret := create2(
        0, // Value sent
        add(code, 0x20), // Skip the length field to get where the bytecode starts
        mload(code), // First 32 bytes are the code length
        salt
      )
    }

    require(ret != address(0), "Failed to deploy contract");

    emit Deployed(ret);

    return ret;
  }
}
