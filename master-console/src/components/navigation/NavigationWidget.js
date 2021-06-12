import React from 'react'
import {Link} from 'react-router-dom'
import './NavigationWidget.css'

function navigationWidget() {
    let links = [{address: "/", name: "home"}, {address: "/kiosks", name: "kiosks"}, {
        address: "images",
        name: "images"
    }, {address: "users", name: "users"}]
    return (
        <div className="NavigationDiv">
            <ul className="NavigationList">
                {links.map((link) => {
                    return (
                        <Link key={link.name} to={link.address}>
                            <li>{link.name}</li>
                        </Link>
                    );
                })}
            </ul>
        </div>
    )
}

export default navigationWidget