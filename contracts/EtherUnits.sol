// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract EtherUnits {
    uint public oneWei = 1 wei;
    bool public isOneWei = 1 wei == 1;

    uint public oneEther = 1 ether;
    bool public oneEther = 1 ether == 1e18;
}
