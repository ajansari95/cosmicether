# EthState Module 

The `ethstate` module provides an interface to store and verify the state data of an Ethereum contract slot within the Cosmos chain. By collaborating with the `ethquery` module, it fetches the Ethereum data, verifies its authenticity, and makes it available for querying within the Cosmos chain.

## How It Works:

### 1. Data Retrieval:

Users can initiate the retrieval process by executing the `get-slot-data-from-eth` command.

#### Command:
```bash
cosmicetherd tx ethstate get-slot-data-from-eth CONTRACT_ADDRESS SLOT_NUMBER BLOCK_HEIGHT
```

#### Example:
```bash
cosmicetherd tx ethstate get-slot-data-from-eth 0x314159265dD8dbb310642f98f50C066173C1259b 0x02 18077318
```

By running the command, users specify:
- **CONTRACT_ADDRESS**: Ethereum contract address of interest.
- **SLOT_NUMBER**: Slot of interest within the contract.
- **BLOCK_HEIGHT**: Height of the Ethereum block to consider.

On execution:
- A user submits this command which results in a Cosmos transaction.
- This action creates two records in the `ethquery` module: one for fetching the storage data (`getStorage`) and another for fetching the block by its number (`getBlockByNumber`).

### 2. Cosmic-Relayer:

- The `cosmic-relayer` identifies the unserviced `ethquery` requests.
- It then retrieves the respective data from Ethereum using RPC calls in batches.
- The fetched data, while unverified, is submitted back to the Cosmos chain by the relayer.

### 3. Data Storage & Verification:

- Upon the receipt of the Ethereum data, designated callbacks (`GetStorageCallback` and `GetBlockCallback`) are invoked, storing the data appropriately.
- Periodically, the stored data undergoes verification checks using the Merkle Patricia Trie (MPT) and associated proofs.

## User Queries:

Users can retrieve the stored and verified data with the following commands:

1. **Slot Data Query**:
```bash
cosmicetherd q ethstate slot-data CONTRACT_ADDRESS SLOT_NUMBER
```

2. **Contract Data Query**:
```bash
cosmicetherd q ethstate contract-data CONTRACT_ADDRESS
```

3. **Block Header Query**:
```bash
cosmicetherd q ethstate eth-block BLOCK_NUMBER
```

## Verification Functions:

The verification of Ethereum data involves several functions:

- **VerifyAccountProof**: Confirms the authenticity of an Ethereum account's state by matching the provided state against the root hash using the MPT.
- **VerifyProof**: Verifies whether the given value, at the provided key, matches the state in the root hash using the MPT.
- **VerifyEthStorageProof**: Uses `VerifyProof` to confirm the legitimacy of the provided storage data against its hash.
- **VerifyEIP1186**: Validates storage proofs in compliance with the EIP-1186 standard.
- **VerifyEthAccountProof**: Confirms both the storage and account state by comparing them against the provided proofs and root hash.
- **VerifyMPTForSlotData**: Comprehensive function that verifies the data for a particular slot using the MPT, ensuring both storage and account data are legitimate.

These verification functions leverage Ethereum's Merkle Patricia Trie to ensure the correctness of the data. They use cryptographic proofs, hashes, and the MPT structure to validate that a particular piece of data, given its hash and accompanying proof, genuinely belongs to a particular state or block.


Sure, here's a structured section on how to test the POC:

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
# ... [Your testnode.sh script contents]
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

End users can follow these steps sequentially to test the functionality of the system.