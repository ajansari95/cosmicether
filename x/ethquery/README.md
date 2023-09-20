# EthQuery Module 
The `ethquery` module for the Cosmos SDK acts as a bridge between Ethereum and Cosmos chains. It offers an interface for users to post Ethereum-related queries to the Cosmos chain, which are subsequently catered to by a dedicated relayer termed as `cosmic-relayer`.

## Features

1. **Ethereum Query Types:** The module currently facilitates the following queries:
   - `eth_getStorageAt`
   - `eth_block_by_number`

2. **Data Verification:** Associated with the `eth_getStorageAt` query, supplementary data from `eth_getProof` can be appended to ascertain data's validity. Two checks are meticulously carried out:
   - Cross-checking the given `StorageProof` with the `storageHash`.
   - Counter-checking the stated Account state against the `rootHash` of the assigned block.

   The purpose of `eth_block_by_number` is to glean and archive block header information, which aids the second verification step aforementioned.

## Protocol Buffers

The underpinnings of the module are rooted in Protocol Buffers (protobuf) which dictates message definitions.

- **Query Protobuf (`query.proto`):** This blueprint maps out the service to pull a list of queries from the module.

- **Transaction Protobuf (`messages.proto`):** This design defines both the service and the messages to introduce a query response back into the module.

## Usage

1. **Submitting a Query:** You can relay queries related to Ethereum via the `ethquery` module. The queries' nature and their format should remain congruent with the protobuf messages delineated.

2. **Gathering Responses:** Post the submission of a query, the `cosmic-relayer` procures the requisite details from Ethereum, subsequently reverting the response to the `ethquery` module. The relayer stands as a guarantor for the data's authenticity and completeness.

## Querying Pending Queries

To monitor and access pending queries, follow the steps below:

- **Using the SDK CLI:**
    ```bash
    cosmicetherd q ethquery eth-queries
    ```

- **Using HTTP Endpoint:** If you prefer to use an HTTP method to fetch pending queries, you can simply `curl` the provided server endpoint:
    ```bash
    curl <SERVER_URL>:1317/cosmicether/ethquery/queries
    ```

**Note**: Replace `<SERVER_URL>` with the actual server's URL or IP address where the node is hosted.