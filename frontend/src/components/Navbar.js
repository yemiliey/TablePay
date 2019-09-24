import React from 'react';
import { Link } from 'react-router-dom'
 const Navbar = ()=>{
    return(
            <nav className="nav-wrapper">
                <div className="container">
                    <Link to="/" className="brand-logo">Table Pay</Link>

                    <ul className="right">
                        {/* <li><Link to="/">Order</Link></li> */}
                    </ul>
                </div>
            </nav>


    )
}

export default Navbar;