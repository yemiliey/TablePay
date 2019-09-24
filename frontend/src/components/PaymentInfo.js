import React from 'react'

const PaymentInfo = () => (
    <div className="row">
        <form className="col s12">
            <div className="row">
                <div className="input-field col s12">
                    <input id="card" type="text" className="validate"/>
                    <label htmlFor="card">Card Number</label>
                </div>
            </div>
            <div className="row">
                <div className="input-field col s6">
                    <input id="first_name" type="text" className="validate" />
                    <label htmlFor="first_name">Expiration Month</label>
                </div>
                <div className="input-field col s6">
                    <input id="last_name" type="text" className="validate" />
                    <label htmlFor="last_name">Expiration Year</label>
                </div>
            </div>
        </form>
    </div>
)

export default PaymentInfo
