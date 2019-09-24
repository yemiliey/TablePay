import React, { Component } from 'react';
import { connect } from 'react-redux'
import { addTableNumber, getOrderDetails } from './actions/cartActions'
import { Link } from 'react-router-dom'

class HomePage extends Component{
    constructor(props) {
        super(props)
        this.state = {
            tableNumTemp:''
        }
        this.handleChange = this.handleChange.bind(this);
    }
    isValid = (tableNum) => {
        return tableNum > 0;
    }
    handleSubmit = (tableNum)=>{
        // addTableNumber(tableNum)
        // alert(tableNum)
        addTableNumber(tableNum)
        this.props.getOrderDetails(tableNum);
    }
    handleChange(event) {
        this.setState({tableNumTemp: event.target.value});
    }
    render() {
        const { tableNum } = this.props
        return(
        <div className="container" style={{width:"500px"}}>
            <h3>Enter Your Table Number</h3>
            <div className="row">
                <form className="col s12">
                    <div className="row">
                        <div className="input-field col s12">
                            <input
                                id="table_num"
                                type="text"
                                className="validate"
                                value={this.state.tableNumTemp}
                                onChange={this.handleChange}
                            />
                            <label for="table_num"></label>
                        </div>
                    </div>
                </form>
            </div>

            <Link to="/orderview">
                <button className="waves-effect waves-light btn"
                onClick={() => {
                    this.handleSubmit(this.state.tableNumTemp)
                }}>Submit</button>
            </Link>
        </div>
        )
    }
}

const mapStateToProps = (state)=>{
    return{
        // tableNum: state.tableNum
    }
}
const mapDispatchToProps = (dispatch)=>{
    return{
        addTableNumber: (tableNum)=>{dispatch(addTableNumber(tableNum))},
        getOrderDetails: (tableNum)=>{getOrderDetails(tableNum, dispatch)},
    }
}
export default connect(mapStateToProps,mapDispatchToProps)(HomePage)
