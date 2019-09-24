import Item1 from '../../images/grassFedBanana.png'
import Item2 from '../../images/friedBanana.jpg'
import Item3 from '../../images/frozenBanana.jpg'
import { CHECK_OUT, GET_ORDER_DETAILS_SUCCESS, ADD_TABLE_NUMBER } from '../actions/action-types/cart-actions'


const initState = {
    items: [
        {id:1,name:'Grass Fed Organic Banana', desc: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Minima, ex.", price:110,img:Item1},
        {id:2,name:'Deep Fried Banana', desc: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Minima, ex.", price:80,img: Item2},
        {id:3,name:'Frozen Banana', desc: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Minima, ex.",price:120,img: Item3}
    ],
    total: 0,
    tableNum: 0,
    orderId: ''
}
const cartReducer= (state = initState,action)=>{

    //INSIDE OrderViewPage COMPONENT
    if(action.type === CHECK_OUT){
        // to be implemented
    }
    else if (action.type === GET_ORDER_DETAILS_SUCCESS){
        return {
            ...state,
            items: action.items,
            total: action.total,
            orderId: action.orderId
        }
        // to be implemented
    }
    else if (action.type === ADD_TABLE_NUMBER) {
        return {
            ...state,
            tableNum: action.tableNum,
        }
    }
        return state

}

export default cartReducer
