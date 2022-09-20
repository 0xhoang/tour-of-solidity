// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/Todos.sol";

contract TestTodos {
    string public _rs;
    bool public _isCompleted;

    function testCreate() public {
        Todos todos = Todos(DeployedAddresses.Todos());
        todos.create("example");
        (_rs, _isCompleted) = todos.get(0);
        Assert.equal(_rs, "example", "equal");
        Assert.equal(_isCompleted, true, "equal");
    }
}
