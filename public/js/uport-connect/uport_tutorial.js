// Setup
const SimpleSigner = window.uportconnect.SimpleSigner
const Connect = window.uportconnect.Connect
const appName = 'Persper foundation'
// Setup the objective Dapp, including clientID and Signer(important)
const connect = new Connect("Persper", {
clientId: "2okgFPYSjSjkVnXM6rFLG3e7gcfF7n6xisF",
signer: SimpleSigner("88a8736d7bf7e7f1fa3dea8db75b6caea79bb0d337b3482633b5ec040b04414e"),
network: 'rinkeby'
})
const web3 = connect.getWeb3()

// Setup the simple Status contract - allows you to set and read a status string
const abi = [{"constant":false,"inputs":[{"name":"status","type":"string"}],"name":"updateStatus","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"addr","type":"address"}],"name":"getStatus","outputs":[{"name":"","type":"string"}],"payable":false,"type":"function"}]

const StatusContract = web3.eth.contract(abi);
const statusInstance = StatusContract.at('0x70A804cCE17149deB6030039798701a38667ca3B')

// uPort connect
const uportConnect = function () {
  connect.requestCredentials({
    requested: ['name', 'phone', 'country', 'email', 'avatar', 'identity_no'],
    notifications: true // We want this if we want to recieve credentials
  })
  .then((credentials) => {
  // get the characters
    console.log("Credentials:", credentials);
    globalState.uportId = credentials.address;
    globalState.userName = credentials.name;
    globalState.country = credentials.country;
    globalState.identity_no = credentials.identity_no;
    globalState.email = credentials.email;
    globalState.uportPhone = credentials.phone;
    globalState.avatar = credentials.avatar;
    render();
    }, (err) => {
        console.log("Error:", err);
    })
}

// Send ether
const sendEther = () => {
  const value = parseFloat(globalState.sendToVal) * 1.0e18

  web3.eth.sendTransaction(
    {
      to: globalState.sendToAddr,
      value: value
    },
    (error, txHash) => {
      if (error) { throw error }
      globalState.txHashSentEth = txHash
      render()
    }
  )
}

// Set Status
const setStatus = () => {
  const newStatusMessage = globalState.statusInput
  statusInstance.updateStatus(newStatusMessage,
    (error, txHash) => {
      if (error) { throw error }
      globalState.txHashSetStatus = txHash
      render()
      })
}