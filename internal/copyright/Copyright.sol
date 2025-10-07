// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import "@openzeppelin/contracts/access/Ownable.sol";

contract ReblogCopyright is Ownable {
    string _siteName;

    constructor(string memory siteName) Ownable(msg.sender) {
        _siteName = siteName;
    }

    struct Post {
        string title;
        string ipfsURL;
    }

    mapping (string => Post) private slugToPost;
    
    event PostAddedOrUpdated(string slug, string title, string ipfsURL, address author);
    event SiteNameUpdated(string newName);

    function setSiteName(string memory siteName) public onlyOwner {
        _siteName = siteName;
        emit SiteNameUpdated(siteName);
    }

    function getSiteName() public view returns(string memory) {
        return _siteName;
    }

    function addOrUpdatePost(string memory slug, string memory title, string memory ipfsURL) public onlyOwner {
        require(bytes(slugToPost[slug].ipfsURL).length == 0);
        slugToPost[slug] = Post(title, ipfsURL);

        emit PostAddedOrUpdated(slug, title, ipfsURL, msg.sender);
    }

    function getPost(string memory slug) public view returns(Post memory) {
        return slugToPost[slug];
    }
}
