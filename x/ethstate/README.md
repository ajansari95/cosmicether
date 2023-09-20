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


---
