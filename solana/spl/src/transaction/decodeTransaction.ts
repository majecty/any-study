/*
{"jsonrpc":"2.0","method":"sendTransaction","params":
["AVAwaT1fct/LbKLB85GB0XgcsB9rVHEbnVQ8Dk79x8lZX5Rjn1B2DBLJz27sZ2aIPZ6oOi89C9c4sKpZr7WoLwIBAAEC+PutuZL5Q0Ny4XzfHTA98bSnenZBJIdRuIkcWY/OXnEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAONtpOpF75ETb1Ab/z06UE1NBPMEYEgvxEJtHrfaF+CIAQECAAAMAgAAAEBCDwAAAAAA",{"encoding":"base64"}],"id":4}

*/

import * as solanaWeb3 from "@solana/web3.js";

function main() {
  const encodedTx =
    "AVAwaT1fct/LbKLB85GB0XgcsB9rVHEbnVQ8Dk79x8lZX5Rjn1B2DBLJz27sZ2aIPZ6oOi89C9c4sKpZr7WoLwIBAAEC+PutuZL5Q0Ny4XzfHTA98bSnenZBJIdRuIkcWY/OXnEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAONtpOpF75ETb1Ab/z06UE1NBPMEYEgvxEJtHrfaF+CIAQECAAAMAgAAAEBCDwAAAAAA";
  const buffer = Buffer.from(encodedTx, "base64");

  const tx = solanaWeb3.Transaction.from(buffer);
  console.log("tx", tx);
}

main();
