// Setup
const SimpleSigner = window.uportconnect.SimpleSigner
const Connect = window.uportconnect.Connect
const QRUtil = window.uportconnect.QRUtil
const appName = 'Persper foundation'
// Setup the objective Dapp, including clientID and Signer(important)
const connect = new Connect("Persper", {
clientId: "2okgFPYSjSjkVnXM6rFLG3e7gcfF7n6xisF",
signer: SimpleSigner("88a8736d7bf7e7f1fa3dea8db75b6caea79bb0d337b3482633b5ec040b04414e"),
network: 'rinkeby'
})
const web3 = connect.getWeb3();
console.log("Test Contract Start");
// Setup the simple Status contract - allows you to set and read a status string
// const abi = [{"constant":false,"inputs":[{"name":"status","type":"string"}],"name":"updateStatus","outputs":[],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"addr","type":"address"}],"name":"getStatus","outputs":[{"name":"","type":"string"}],"payable":false,"type":"function"}]
// const StatusContract = web3.eth.contract(abi);
// const getValueInstance = StatusContract.at('0x70A804cCE17149deB6030039798701a38667ca3B')

// const abi = [{"constant":true,"inputs":[],"name":"getHello2","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"getSender2","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"getHello1","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"getThis1","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getSender3","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"getSender1","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getThis2","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"a","type":"uint256"}],"name":"multiply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"sayHello","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"anonymous":false,"inputs":[{"indexed":false,"name":"_value","type":"string"}],"name":"say","type":"event"}]
// const StatusContract = web3.eth.contract(abi);
// // const StatusContract = connect.contract(abi);
// const getValueInstance = StatusContract.at('0xffe459d507ea419a78ca55446b2281177061102f')

// const abi = [{"constant":false,"inputs":[],"name":"testFunc","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"anonymous":false,"inputs":[{"indexed":false,"name":"caller","type":"address"},{"indexed":false,"name":"message","type":"bytes"}],"name":"Becalled","type":"event"}]
// const StatusContract = web3.eth.contract(abi);
// // const getValueInstance = StatusContract.at('0x95783e8aa5fb9a6481be3fa0387b084c01fc278e')
// const getValueInstance = StatusContract.at('0x8fb8204f79a03fb01cb00d6cb906fed15bc55887')


var abi = [{"constant":false,"inputs":[{"name":"holder","type":"address"},{"name":"data","type":"bytes32[]"}],"name":"construct","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"holder","type":"address"},{"name":"data","type":"bytes32[]"}],"name":"construct2","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"destroy","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"vacant","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"retrieve","outputs":[{"name":"holders","type":"address[]"},{"name":"data","type":"bytes32[]"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"holders","type":"address[]"}],"name":"addHolders","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"holders","type":"address[]"}],"name":"addHolders2","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"holders","type":"address[]"}],"name":"removeHolders","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"holders","type":"address[]"}],"name":"removeHolders2","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"holder","type":"address"}],"name":"hasPermission","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"data","type":"bytes32[]"}],"name":"resetData","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"indexes","type":"uint256[]"},{"name":"items","type":"bytes32[]"}],"name":"updateData","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"holder","type":"address"}],"name":"removeTT","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"countTT","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"holder","type":"address"}],"name":"addTT","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"listTT","outputs":[{"name":"","type":"address[]"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"addTTSelf","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[],"name":"clearTT","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]
var getPermission = web3.eth.contract(abi);
var getValueInstance = getPermission.at("0xd86c1d8dd43e3d251189da0aedc5a2ca41afa515")


// const abi = [{"constant":false,"inputs":[],"name":"testFunc","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"anonymous":false,"inputs":[{"indexed":false,"name":"caller","type":"address"},{"indexed":false,"name":"message","type":"bytes"}],"name":"Becalled","type":"event"}]
// const StatusContract = web3.eth.contract(abi);
// // const getValueInstance = StatusContract.at('0x95783e8aa5fb9a6481be3fa0387b084c01fc278e')
// const getValueInstance = StatusContract.at('0x83e602e0a73c7bc570b06cf67898d6d2616e0f17')


// // uPort connect
// const uportConnect = function () {
//   connect.requestCredentials({
//     requested: ['name', 'phone', 'country', 'email', 'avatar', 'identity_no'],
//     notifications: true // We want this if we want to recieve credentials
//   })
//   .then((credentials) => {
//   // get the characters
//     console.log("Credentials:", credentials);
//     globalState.uportId = credentials.address;
//     globalState.userName = credentials.name;
//     globalState.country = credentials.country;
//     globalState.identity_no = credentials.identity_no;
//     globalState.email = credentials.email;
//     globalState.uportPhone = credentials.phone;
//     globalState.avatar = credentials.avatar;
//     render();
//     }, (err) => {
//         console.log("Error:", err);
//     })
// }

// // Send ether
// const sendEther = () => {
//   const value = parseFloat(globalState.sendToVal) * 1.0e18

//   web3.eth.sendTransaction(
//     {
//       to: globalState.sendToAddr,
//       value: value
//     },
//     (error, txHash) => {
//       if (error) { throw error }
//       globalState.txHashSentEth = txHash
//       render()
//     }
//   )
// }

// uPort connect
const uportConnect = function () {
  web3.eth.getCoinbase((error, address) => {
    if (error) { throw error }
    globalState.ethAddress = address
    // This one is for display purposes - MNID encoding includes network 
    globalState.uportId = window.uportconnect.MNID.encode({network: '0x4', address: address})
    getValueInstance.getStatus.call(globalState.ethAddress, (err, st) => {
      globalState.currentStatus = st
      web3.eth.getBalance(globalState.ethAddress, (err, bal) => {
        globalState.ethBalance = web3.fromWei(bal)
        render()
      })
    })
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
// const setStatus = () => {
//   // const newStatusMessage = globalState.statusInput
//   globalState.txHashSetStatus = getValueInstance.getSender2()
//   globalState.ethAddress = getValueInstance.getSender3()
//   globalState.sendToAddr = getValueInstance.getThis2()

// }
const setStatus = () => {
    const uriHandler = (uri) => {
      const qrCode = QRUtil.getQRDataURI(uri)
    }
    console.log("Test Contract Start 0000 ");
    console.log("Current default: " + web3.eth.defaultAccount);
    // console.log("Current default: " + web3.eth.getTransactionCount);
    // console.log("Current default: " + web3.eth.getStorageAt);
    // console.log("Current default: " + web3.eth.call);
    getValueInstance.addTTSelf.sendTransaction( (error, txHash) => {
        if (error) {
          actions.getPointsERROR(error)
          throw error
        }
        console.log("getTxHash",txHash);
        render()
      // return posPoints
  })
    // var event = getValueInstance.say();
    // event.watch(function (error, result) {  
    //   if(!error){ console.log(JSON.stringify(result)); }  
    // });
    // getValueInstance
    // window.addEventListener('load', function() {
    // var testEvent = getValueInstance.Becalled();
    // testEvent.watch(function (error, result) {  
    //   // globalState.txHashSetStatus = "result"; 
    //   if(!error){
    //     console.log(result);
    //     // window.alert(result);
    //     // globalState.txHashSetStatus = result;
    //     render()
    //   } else {
    //     // window.alert("error");
    //     // globalState.txHashSetStatus = "error"; 
    //     render()
    //   }
    // }); 
      // if(!error){  globalState.txHashSetStatus=JSON.stringify(result); console.log(JSON.stringify(result)); }  });
    // getValueInstance.addTTSelf().sendTransaction( (error, txHash) => {
    //   if(error) {
    //     actions.getPointsERROR(error)
    //     throw error
    //   }
    //   console.log("getTxHash",txHash);
    //   render()
    // });
  // const newStatusMessage = globalState.statusInput

  console.log("Test Contract Start 333");
}