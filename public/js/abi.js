const JNSABI = [{"constant":true,"inputs":[{"name":"_address","type":"address"}],"name":"getNickname","outputs":[{"name":"_nickname","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_nickname","type":"string"},{"name":"_address","type":"address"}],"name":"registerAddress","outputs":[{"name":"","type":"bool"}],"payable":true,"stateMutability":"payable","type":"function"},{"constant":true,"inputs":[{"name":"_nickname","type":"string"}],"name":"getAddress","outputs":[{"name":"_address","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"_wallet","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"}];


var contractAddress = "";
var contract;
var web3;
var eth;
var newAddress = "";

function setContract(error, networkId) {
    if (error == null) {
        if (networkId !== "1") {
            contractAddress = "0x4ead784006cb920965cb1b45a9f6923330a3ccf0";
        } else {
            contractAddress = "0x023fa9e2a97799b3D87B3fa35674b50B8b5c9f4E";
        }
        contract = web3.eth.contract(JNSABI).at(contractAddress);
        console.log(networkId);
        console.log(contractAddress);
    } else {
        console.error(error);
    }
}

function checkAvailability() {
    var frm = document.frmUrl;
    //console.log(frm.url.value);
    contract.getAddress(frm.url.value, function (err, result) {
        if (!err) {
            console.log(result);
            if (result == "0x0000000000000000000000000000000000000000") {
                $("#availabilityStatus").removeClass("glyphicon glyphicon-remove");
                $("#availabilityStatus").addClass("glyphicon glyphicon-ok");
                $("#url").css("background-color", "white");
            } else {
                $("#availabilityStatus").removeClass("glyphicon glyphicon-ok");
                $("#availabilityStatus").addClass("glyphicon glyphicon-remove");
                $("#url").css("background-color", "tomato");
            }
        } else {
            console.error(err);
        }
    });    
}

function iAmLucky() {
    var urlshort = '', character; 
    while (4 > urlshort.length) {
        if (urlshort.indexOf(character = String.fromCharCode(Math.floor(Math.random() * 90) + 48), Math.floor(urlshort.length / 90) * 90) < 0) {
            if ( (character.charCodeAt(0) >=48 && character.charCodeAt(0) <=57) || (character.charCodeAt(0) >=65 && character.charCodeAt(0) <=90) ) {
                urlshort += character;
            }
        }
    }
    urlshort = urlshort.toLowerCase(); 
    contract.getAddress(urlshort, function (err, result) {
        if (!err) {
            console.log(result);
            if (result == "0x0000000000000000000000000000000000000000") {
                $("#availabilityStatus").removeClass("glyphicon glyphicon-remove");
                $("#availabilityStatus").addClass("glyphicon glyphicon-ok");
                $("#url").css("background-color", "white");
            } else {
                $("#availabilityStatus").removeClass("glyphicon glyphicon-ok");
                $("#availabilityStatus").addClass("glyphicon glyphicon-remove");
                $("#url").css("background-color", "tomato");
            }
            $("#url").val(urlshort);
        } else {
            console.error(err);
        }
    });    
}

function shorten() {
    var frm = document.frmUrl;
    if (frm.url.value.trim().length < 2) {
        alert("Invalid URL");
        return;
    }

    if (frm.address.value.trim().length < 30) {
        alert("Invalid Ethereum or Bitcoin address");
        return;
    }
    var nickname = frm.url.value.trim().toLowerCase();
    var addressTo = frm.address.value.trim().toLowerCase();
    //console.log(nickname);
    //console.log(addressTo);
    contract.registerAddress(nickname, addressTo, {from: web3.eth.accounts[0], gas: 3000000, value: 1000000000000000}, function (err, result) {
        if (!err) {
            $("#divTrxStatus").css("background-color", "yellow");
            $("#trxStatus").text("Transaction sent. Wait until it is mined. Transaction hash: " + result);
            alert("Transaction sent. Wait until it is mined.");
            newAddress = "https://etd.bz/" + nickname;
            waitForTxToBeMined(result);
        } else {
            console.error(err);
        }
    });
    $("#btnShorten").hide();
    $("#btnStartOver").show();
}

async function waitForTxToBeMined (txHash) {
    let txReceipt
    while (!txReceipt) {
      try {
        txReceipt = await eth.getTransactionReceipt(txHash)
      } catch (err) {
        return console.error(err);
      }
    }
    $("#divTrxStatus").css("background-color", "LawnGreen");
    $("#trxStatus").html("The address is ready at <a href='" + newAddress + "'>" + newAddress + "</a>");
    $("#btnStartOver").show();
    //window.location.href = newAddress;
}

function startOver() {
    var frm = document.frmUrl;
    frm.url.value = "";
    frm.address.value = ""
    $("#btnStartOver").hide();
    $("#divTrxStatus").css("background-color", "white");
    $("#trxStatus").text("");
    $("#btnShorten").show();
}

if (typeof web3 !== 'undefined') {
    web3 = new Web3(web3.currentProvider);
} else {
    web3 = new Web3(new Web3.providers.HttpProvider("https://rinkeby.infura.io/QPF0qjGpH9OjFuuMrCse"))
}

web3.version.getNetwork(setContract);

if (typeof eth !== 'undefined') {
    eth.setProvider(web3.currentProvider);
} else {
    eth = new Eth(new Eth.HttpProvider('https://rinkeby.infura.io/QPF0qjGpH9OjFuuMrCse'));
    eth.setProvider(web3.currentProvider);
} 

$("#btnStartOver").hide();

//console.log(eth);



