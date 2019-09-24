import NodeRSA from 'node-rsa';

export function EncryptCard(cardNumber, prefix, modulus, exponent) {
    var plainText = prefix + cardNumber
    var rsa = new NodeRSA();
    rsa.keyPair.setPublic(modulus.toString(16), exponent.toString(16));
    var res = rsa.encrypt(plainText, 'base64');

    if (res) {
        return res;
    } else {
        return "Not encoded"
    }
}