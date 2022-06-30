console.clear();
const {
  AccountId,
  PrivateKey,
  Client,
  AccountBalanceQuery,
} = require("@hashgraph/sdk");

// Grab the OPERATOR_ID and OPERATOR_KEY from the .env file
const operatorId = AccountId.fromString(
  "0.0.46018668" || process.env.OPERATOR_ID,
);
const operatorKey = PrivateKey.fromString(
  "302e020100300506032b6570042204208c8e7386dd9a19356a2fb94718955f53ed30bb437e9afc1789426fa9bf90ab8fs" ||
    process.env.OPERATOR_KEY,
);

console.log(operatorId, operatorKey);

// Build Hedera testnet and mirror node client
const client = Client.forTestnet();

// Set the operator account ID and operator private key
client.setOperator(operatorId, operatorKey);

async function main() {
  // Create the account balance query
  const query = new AccountBalanceQuery().setAccountId(operatorId);
  //Submit the query to a Hedera network
  const accountBalance = await query.execute(client);
  //Print the balance of hbars
  console.log(
    "The hbar account balance for this account is " + accountBalance?.hbars,
  );
  // v2.0.7
}
main();
