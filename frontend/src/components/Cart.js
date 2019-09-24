import React, { Component } from 'react';
import { connect } from 'react-redux'
import { checkOut, getOrderDetails } from './actions/cartActions'
import PaymentInfo from './PaymentInfo'
import { stat } from 'fs';
class Cart extends Component{

    //to remove the item completely
    handleCheckOut = ()=>{
        this.props.checkOut(
            this.props.orderId,
            this.props.total,
            this.cardNumber,
            this.month,
            this.year)
    }
    handleGetOrderDetails = () => {
        this.props.getOrderDetails()
    }
    cardNumberChanged = (cardNumber) => {
        var val = document.getElementById('carddddd').value
        this.cardNumber = val
    }
    monthChanged = (month) => {
        var val = document.getElementById('monthhhh').value
        this.month = val
    }
    yearChanged = (year) => {
        var val = document.getElementById('yearrrrr').value
        this.year = val
    }

    render(){

        let items = this.props.items.length ?
            (
                this.props.items.map(item=>{
                    return(
                        <li className="collection-item avatar">

                            <div className="item-desc">
                                <span className="title">{item.name}</span>
                                <p><b>Price: {item.price}$</b></p>
                            </div>
                        </li>
                    )
                })
            ):
             (
                <p> You have not ordered anything yet</p>
             )
       return(
            <div className="container">
                <div className="cart">
                    <h3>Order Summary</h3>
                    <ul className="collection">
                        {items}
                    </ul>
                </div>
                <div className="collection">
                    <li className="collection-item"><b>Total: {this.props.total} $</b></li>
                </div>
                <h3>Payment Info</h3>
                <div className="row">
                <form className="col s12">0
                    <div className="row">
                        <div className="input-field col s12">
                            <input onChange={this.cardNumberChanged} id="carddddd" type="text" className="validate"/>
                            <label htmlFor="carddddd">Card Number</label>
                        </div>
                    </div>
                    <div className="row">
                        <div className="input-field col s6">
                            <input id="monthhhh" onChange={this.monthChanged} type="text" className="validate" />
                            <label htmlFor="monthhhh">Expiration Month</label>
                        </div>
                        <div className="input-field col s6">
                            <input id="yearrrrr"  onChange={this.yearChanged} type="text" className="validate" />
                            <label htmlFor="yearrrrr">Expiration Year</label>
                        </div>
                    </div>
                </form>
                </div>
                <div className="checkout">
                    <button className="waves-effect waves-light btn" onClick={this.handleCheckOut}>Checkout</button>
                    <button className="waves-effect waves-light btn" style={{marginLeft: '15px'}}>Refresh</button>
                </div>
            </div>
       )
    }
}


const mapStateToProps = (state)=>{
    return{
        items: state.items,
        orderId: state.orderId,
        total: state.total
    }
}
const mapDispatchToProps = (dispatch)=>{
    return {
        checkOut: (orderId, amount, cardNumber, expMonth, expYear)=>{checkOut(orderId, amount, cardNumber, expMonth, expYear, dispatch)},
        getOrderDetails:()=>{getOrderDetails(dispatch)},
    }
}
export default connect(mapStateToProps,mapDispatchToProps)(Cart)
