// SPDX-License-Identifier: MIT
pragma solidity ^0.8.18;

contract Sample {
  uint256 private _value;

  constructor(uint256 value) {
    _value = value;
  }

  function set(uint256 value) external {
    _value = value;
  }

  function get() external view returns (uint256) {
    return _value;
  }
}
