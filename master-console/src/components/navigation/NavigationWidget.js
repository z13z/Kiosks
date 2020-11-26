import React from 'react'
import {Link, NavLink} from 'react-router-dom'
import './NavigationWidget.css'

function NavigationWidget() {
    let links = [{address: "/", name: "home"}, {address: "/kiosks", name: "kiosks"}, {
        address: "images",
        name: "images"
    }]
    return (
        <div className="NavigationDiv">
            <ul className="NavigationList">
                {links.map((link) => {
                    return (
                        <Link to={link.address}>
                            <li>{link.name}</li>
                        </Link>
                    );
                })}
            </ul>
        </div>
    )
}

export default NavigationWidget