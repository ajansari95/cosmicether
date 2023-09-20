# CosmicEther: Ethereum Data Integration for Cosmos SDK

`CosmicEther` bridges Ethereum data into the Cosmos ecosystem, making it possible for Cosmos chains to fetch, store, and verify Ethereum state data seamlessly. Built around two core modules, `ethquery` and `ethstate`, this framework provides a robust, secure, and efficient way to bring Ethereum's rich state data to the Cosmos environment.

## Modules:

### 1. **EthQuery**

The `ethquery` module acts as a bridge between Cosmos SDK and Ethereum. It assists in creating an Ethereum query stub on the Cosmos side, which is then serviced by a designated relayer named `cosmic-relayer`.

- **Key Features**:
    - Supports two types of Ethereum queries: "eth_getStorageAt" and "eth_block_by_number".
    - Uses Ethereum's "eth_getProof" for data verification through Merkle Patricia Trie (MPT).

[Detailed EthQuery README](./x/ethquery/README.md)

### 2. **EthState**

The `ethstate` module manages the Ethereum contract slot's data within Cosmos. It fetches the data through the `ethquery` module, then verifies and maintains it, ensuring its authenticity and availability for Cosmos chain users.

- **Key Features**:
    - Provides a transaction to initiate data retrieval from Ethereum.
    - Features callbacks for data storage and periodic verification.
    - Offers user-friendly commands to query stored Ethereum data on Cosmos.

[Detailed EthState README](./x/ethstate/README.md)

## Getting Started:

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/ajansari95/cosmicether
    ```

2. **Build and Install**:
    ```bash
    cd cosmicether
    make build
    make install
    ```

3. **Verify Installation**:
   To verify that the installation was successful, you can check the version:
    ```bash
    cosmicetherd --version
    ```

Follow the module-specific READMEs for detailed usage instructions.
