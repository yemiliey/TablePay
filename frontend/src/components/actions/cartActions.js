
import { CHECK_OUT, GET_ORDER_DETAILS_START, GET_ORDER_DETAILS_SUCCESS, GET_ORDER_DETAILS_FAILED, ADD_TABLE_NUMBER } from './action-types/cart-actions'
import { getOrder } from '../network/network'
import axios from 'axios'
import { EncryptCard } from '../Encrypt'

const getItems = (tableNum) => {
    try {
      return axios.get(`https://tablepay.uw2.zd.cntr.io/tablepay/5c671cd0719df80001d4656c/order/${tableNum}`)
    } catch (error) {
      console.error(error)
    }
  }
  
const getKeys = () => {
    try {
      return axios.get('https://tablepay.uw2.zd.cntr.io/tablepay/F5YXVF6JV7W86/d8ca4666-6295-ec3f-9e14-d147fd2baa80')
    } catch (error) {
      console.error(error)
    }
  }
  
const checkThatShitOut = (keys, orderId, amount, cardNumber, expMonth, expYear) => {
    console.log("JJJJ " + JSON.stringify(keys.data))
    cardNumber = cardNumber + ''
    var keysJson = keys.data;
    var cardEncrypted = EncryptCard(cardNumber, keysJson.prefix, keysJson.modulus, keysJson.exponent)
    var payload = {
        "amount": amount + "",
        "apikey": "d8ca4666-6295-ec3f-9e14-d147fd2baa80",
        "card_encrypted": cardEncrypted,
        "exp_month": expMonth,
        "exp_year": expYear,
        "first6": cardNumber.substring(0, 6),
        "last4": cardNumber.substring(cardNumber.length - 4, cardNumber.length)
    }
    console.log("MMMMMMMM" + JSON.stringify(payload))
    try {
      return axios.post(`https://tablepay.uw2.zd.cntr.io/pay/F5YXVF6JV7W86/${orderId}`, payload)
    } catch (error) {
      console.error(error)
    }
  }
  
//submit order action
export const checkOut= (orderId, amount, cardNumber, expMonth, expYear, dispatch)=>{
    return getKeys()
        .then(keys => {
            checkThatShitOut(keys, orderId, amount, cardNumber, expMonth, expYear)
                .then(response => {
                    console.log("!!!!!!!!!" + JSON.stringify(response.data))
                    dispatch({
                        type: GET_ORDER_DETAILS_START
                    })
                })
        })
        .catch(error => {
            console.log(error)
        })

}

export const getOrderDetails= (tableNum, dispatch)=>{
    return getItems(tableNum)
    .then(response => {
        console.log("!!!!!!!!!" + JSON.stringify(response.data))
        items.
        dispatch({
            type: GET_ORDER_DETAILS_SUCCESS,
            items: response.data.items,
            total: 123,
            orderId: response.data.order_id
        })
    })
    .catch(error => {
        console.log(error)
    })
}
export const addTableNumber = (tableNum)=> {
    return {
        type: ADD_TABLE_NUMBER,
        tableNum
    }
}
