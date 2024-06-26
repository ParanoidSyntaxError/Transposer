// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

interface ITransmuterAdmin {
    function chainSelector() external view returns (uint64);
    function ccipRouter() external view returns (address);

    function getCcipRouter(uint64 chain) external view returns (address);
    function getTransmuter(uint64 chain) external view returns (address);

    function setCcipRouters(
        uint64[] memory chains,
        address[] memory routers
    ) external;

    function setTransmuters(
        uint64[] memory chains,
        address[] memory transmuters
    ) external;
}
