// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0 <0.9.0;

import "./callee.sol";

contract Caller {
    function someAction(address addr) public pure returns (uint) {
        Callee c = Callee(addr);
        return c.getValue(100);
    }

    function storeAction(address addr) public returns (uint) {
        Callee c = Callee(addr);
        c.storeValue(100);
        return c.getValues();
    }

    function someUnsafeAction(address addr) public pure {
        keccak256(abi.encode("storeValue(uint256)", 100, addr));
        //addr.call(bytes4(keccak256("storeValue(uint256)", 100)));
    }
}

abstract contract CalleeAbstract {
    function getValue(uint initialValue) public virtual returns (uint);

    function getValues() public  virtual returns (uint);

    function storeValue(uint value) public virtual;

}

