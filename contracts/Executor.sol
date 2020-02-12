pragma solidity >=0.5.0 <0.7.0;

contract Executor {

    event BoolLog(string name, bool result);
    event AddressLog(string name, address addr);
    
    /**
     * executeFunction is used to execute function provided in bytes code from a given address.
     * @param target is target's address that is used to call function.
     * @param code is encoded from function and its params.
     * @return result of call process.
     * @notice this is a payable function in case users want to trigger a payable function in target.
     */
    function executeFunction(address payable target, bytes memory code) public payable returns (bool success, bytes memory result) {
        emit AddressLog("start executing function", target);
        (success, result) = target.call.value(msg.value)(code);

        emit BoolLog("finish call", success);
        return (success, result);
    }
}
