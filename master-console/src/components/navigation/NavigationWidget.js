import React from 'react'
import { Link } from 'react-router-dom'

function navigationWidget() {
    let links = [{ address: "/", name: "home" }, { address: "/kiosks", name: "kiosks" }, {
        address: "images",
        name: "images"
    }, { address: "users", name: "users" }]
    return (
        <div className="LeftPanel">
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
            <div>
                <div>{localStorage.getItem("currentUser")}</div>
                <a href='/' onClick={() => localStorage.clear()}>logout</a>
            </div>
        </div>

    )
}

export default navigationWidget