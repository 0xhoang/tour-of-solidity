// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.8.0 <0.9.0;

contract Caller {
    function someAction(address addr) public pure returns (uint) {
        CalleeAbs c = CalleeAbs(addr);
        return c.getValue(100);
    }

    function storeAction(address addr) public returns (uint) {
        CalleeAbs c = CalleeAbs(addr);
        c.storeValue(100);
        return c.getValues();
    }

    function someUnsafeAction(address addr) public returns (bool) {
        //(bool success,) = addr.call(abi.encodeWithSelector(bytes4(keccak256("storeValue(uint256)")), 100));
        (bool success,) = addr.call(abi.encodeWithSignature("storeValue(uint256)", 100));

        require(success, "storeValue failed.");
        return success;
    }
}

abstract contract CalleeAbs {
    function getValue(uint initialValue) public pure virtual returns (uint);
    function getValues() public virtual returns (uint);
    function storeValue(uint value) public virtual;
}