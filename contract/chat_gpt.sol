// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Owner {
    address private owner;

    constructor() {
        owner = msg.sender;
    }

    struct WhiteData {
        address wallet;
        bool whitelist;
    }

    enum Status {
        Created,
        Approved,
        Closed
    }

    mapping(string => address) public logins;
    mapping(address => WhiteData) public whiteList;
    Product[] public products;

    modifier checkOfWhiteLists(address adr) {
        require(!whiteList[adr].whitelist, "Address is already whitelisted");
        _;
    }

    modifier checkStatusProduct(uint product_id, Status status) {
        require(
            products[product_id].status == status,
            "Incorrect product status"
        );
        _;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    function createUser(string memory _login) public {
        require(logins[_login] == address(0), "Login already exists");
        logins[_login] = address(0);
    }

    function updateWhiteList(address wallet) public onlyOwner {
        require(!whiteList[wallet].whitelist, "Address is already whitelisted");
        whiteList[wallet] = WhiteData(wallet, whiteList[wallet].whitelist);
    }

    function createProduct(
        string memory name,
        uint value,
        string memory info
    ) public checkOfWhiteLists(msg.sender) {
        products.push(
            Product(
                products.length,
                name,
                msg.sender,
                value,
                Status.Created,
                info
            )
        );
    }

    function approveProduct(
        uint product_id
    ) public checkStatusProduct(product_id, Status.Created) onlyOwner {
        products[product_id].status = Status.Approved;
    }

    function buyProduct(
        uint product_id
    ) internal checkStatusProduct(product_id, Status.Approved) {
        products[product_id].status = Status.Closed;
        uint fee = products[product_id].value / 100;
        uint newValue = products[product_id].value - fee;
        payable(products[product_id].owner).transfer(newValue);
        products[product_id].owner = msg.sender;
    }

    function sellProduct(
        uint product_id
    ) public checkStatusProduct(product_id, Status.Closed) {
        products[product_id].status = Status.Approved;
    }

    function withdrawFee() public onlyOwner {
        payable(owner).transfer(address(this).balance);
    }

    struct Product {
        uint product_id;
        string name;
        address owner;
        uint value;
        Status status;
        string info;
    }
}
