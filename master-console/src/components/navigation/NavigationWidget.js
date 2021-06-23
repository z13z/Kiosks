import React from 'react'
import { Link } from 'react-router-dom'

function navigationWidget() {
    const links = [{ address: "/", name: "home" }, { address: "/kiosks", name: "kiosks" }, {
        address: "images",
        name: "images"
    }, { address: "users", name: "users" }]
    
    const permissions = localStorage.getItem('userPermissions').split(',');
    return (
        <div className="LeftPanel">
            <div className="NavigationDiv">
                <ul className="NavigationList">
                    {links.map((link) => {
                        if(permissions.includes(link.name) || link.name === 'home'){
                            return (
                                <Link key={link.name} to={link.address}>
                                    <li>{link.name}</li>
                                </Link>
                            );
                        }else{
                            return null;
                        }
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