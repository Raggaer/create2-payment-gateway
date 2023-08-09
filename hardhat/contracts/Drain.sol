// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract Drain {
  address payable owner;

  constructor(address payable _owner) {
    owner = _owner;
    if (address(this).balance > 0) {
      (bool success, ) = owner.call{value: address(this).balance}("");
      require(success, "Drain: Failed to send Ether");
    }
  }

  receive() external payable {
    owner.transfer(msg.value);
  }
}
