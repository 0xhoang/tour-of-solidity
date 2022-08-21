// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.0 <0.9.0;

contract Callee {
    uint [] public values;
    function getValue(uint initialValue) public pure returns (uint) {
        return initialValue + 150;
    }

    function getValues() public returns (uint) {
        return values.length;
    }

    function storeValue(uint value) public {
        values.push(value);
    }
}

