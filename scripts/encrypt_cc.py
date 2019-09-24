#######################################
######### BEGIN SCRIPT CONFIG #########
#######################################

MID = "F5YXVF6JV7W86"
API_TOKEN = "d8ca4666-6295-ec3f-9e14-d147fd2baa80"
NUM_ORDERS = 1
ENVIRONMENT = "https://sandbox.dev.clover.com/" # or https://api.clover.com/ or https://eu.clover.com/

#######################################
########## END SCRIPT CONFIG ##########
#######################################

######################################
########## OTHER CONSTANTS ###########
######################################
cardNumber = "6011361000006668"
expMonth = 12
expYear = 2020
CVV = None
######################################
######################################

import requests
import json
from random import randint
import sys
from time import sleep
from Crypto.PublicKey import RSA
from Crypto.Cipher import PKCS1_OAEP
from base64 import b64encode

# Fetch developer pay secrets from GET /v2/merchant/{mId}/pay/key
url = ENVIRONMENT + "v2/merchant/" + MID + "/pay/key"
headers = {"Authorization": "Bearer " + API_TOKEN}
response = requests.get(url, headers = headers)
if response.status_code != 200:
    print("Something went wrong fetching Developer Pay API secrets")
    sys.exit()
print("successfully fetched developer pay api secrets")
response = response.json()

modulus = int(response["modulus"])
exponent = int(response["exponent"])
prefix = str(response["prefix"])

RSAkey = RSA.construct((modulus, exponent))

# create a cipher from the RSA key and use it to encrypt the card number, prepended with the prefix from GET /v2/merchant/{mId}/pay/key
cipher = PKCS1_OAEP.new(RSAkey)
# encode str to byte (https://eli.thegreenplace.net/2012/01/30/the-bytesstr-dichotomy-in-python-3)
encrypted = cipher.encrypt((prefix + cardNumber).encode())

# Base64 encode the resulting encrypted data into a string to use as the cardEncrypted' property.
cardEncrypted = b64encode(encrypted)
print("card encrypted:", cardEncrypted)
