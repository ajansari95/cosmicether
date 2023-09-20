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


---

## Testing the POC

In order to verify the functionality of the `CosmicEther` bridge, follow the steps outlined below:

### Step 1: Installation
- Install both `cosmicether` and `cosmic-eth-relayer`.
    - `cosmic-eth-relayer` repository link: [cosmic-eth-relayer](https://github.com/ajansari95/cosmic-eth-relayer)

### Step 2: Configuration
- Configure the `cosmic-relayer` by creating or updating the `config.yaml` file in the relayer's HOME directory. Populate the file as follows:
```yaml
querier_chain: test-1
eth_rpc: https://mainnet.infura.io/v3/<API-KEY>
chains:
    test-1:
      key: default
      chain-id: test-1
      rpc-addr: http://localhost:26657
      grpc-addr: http://localhost:9090
      account-prefix: cosmos
      keyring-backend: test
      gas-adjustment: 1.5
      gas-prices: 0.0001stake
      min-gas-amount: 0
      key-directory: ./cosmicrelayer/keys
      debug: true
      timeout: 20s
      block-timeout: ""
      output-format: json
      sign-mode: direct
```
- Remember to add the key to the relayer.

### Step 3: Node Setup
- Execute the `testnode.sh` script:
```bash
sh testnode.sh
```

### Step 4: Fund Transfer
- Transfer funds to your relayer's wallet:
```bash
cosmicetherd tx bank send [myKey] [relayer-address] 9000stake --chain-id test-1 --keyring-backend test 
```

### Step 5: Start the Relayer
- Initiate the relayer using the command:
```bash
cosmic-relayer --home RELAYER_HOME_DIR run
```

### Step 6: Create SlotData Requests
- Create a few `get-slot-data-from-eth` requests:
```bash
cosmicetherd tx ethstate get-slot-data-from-eth 0x314159265dD8dbb310642f98f50C066173C1259b 0x02 18077318 --from myKey --keyring-backend test --chain-id test-1
cosmicetherd tx ethstate get-slot-data-from-eth 0x314159265dD8dbb310642f98f50C066173C1259b 0x03 18077318 --from myKey --keyring-backend test --chain-id test-1
cosmicetherd tx ethstate get-slot-data-from-eth 0xC18360217D8F7Ab5e7c516566761Ea12Ce7F9D72 0x03 18077319 --from myKey --keyring-backend test --chain-id test-1

```

### Step 7: Relaying Process
- Note that the relaying process might take some seconds. Be patient and wait for the completion.

### Step 8: Query Data
- You can query the data you've just requested:
    - For querying a single slot:
    ```bash
    cosmicetherd q ethstate slot-data 0x314159265dD8dbb310642f98f50C066173C1259b 0x02
    ```
    - For querying the entire contract:
    ```bash
    cosmicetherd q ethstate contract-data 0xC18360217D8F7Ab5e7c516566761Ea12Ce7F9D72 
    ```

### Step 9: Verification Logic
- Every 10 blocks (this interval is configurable), the verification logic is initiated. Until verification is completed, the data's state will display `verified: false`.

---

## Future Improvements

- The current framework, works as a POC in its initial offering, has various potential areas for expansion and refinement. Here are some areas we've identified for future enhancement:

1. **Performance Optimization:** The current implementation of data retrieval and validation can be further optimized. The usage of more efficient algorithms or introducing parallel processing can expedite the query-response cycle especially the relayer part.

2. **Broadened Query Support:** The addition of more Ethereum query types will expand the capabilities of the framework, catering to a wider range of use cases and wide range of callback support depending on usecase.

3. **Enhanced Security Measures:** While currently i have used MPT and proves to verify data. There can be other ways that will need some research.

4. **Interactivity with Other Blockchains:** Broadening the framework's scope beyond Ethereum, other EVM chains will widen the applicability of `CosmicEther`.

5. **Modular Architecture:** Restructuring the codebase to further modularize components can enhance scalability, facilitate easier updates and migration if needed. for eg. proto versioning.

8. **Advanced Verification Mechanisms:** Beyond the current verification logic that kicks in periodically. for eg. Baking verification in the consensus layer i.e tendermint.

9. **Increasing Usability of Data Stored:** The current implementation of the `ethstate` module stores the data as is. We can just view it, if schema for a contract is known we can parse it and store it in a more usable format, but getting access to ABI is only way i can look at it - will research more on it to see if there is a better way.
