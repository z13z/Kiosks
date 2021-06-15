import React from 'react'
import './LoginPage.css'
import axios from "axios"
import {JWT_TOKEN_KEY, USER_PERMISSIONS_KEY, USERNAME_KEY} from "../../../Constants"

const LOGIN_PAGE_PATH = '/login'

const LoginPage = () => {
    let userData = {username: null, password: null};

    const updateUsername = (event) => {
        userData.username = event.target.value
    }
    const updatePassword = (event) => {
        userData.password = event.target.value
    }
    const loginUserAction = (response) => {
        localStorage.setItem(JWT_TOKEN_KEY, response.data.jwt)
        const jwtBody = JSON.parse(atob(response.data.jwt.split(".")[1]))
        localStorage.setItem(USERNAME_KEY, jwtBody.username)
        localStorage.setItem(USER_PERMISSIONS_KEY, jwtBody.permissions.replaceAll('\"', "").slice(1, -1))
        window.location.reload()
    }
    const loginErrorAction = (error) => {
        console.log(error)
    }

    const loginButtonAction = () => {
        axios.post(LOGIN_PAGE_PATH, userData).then(loginUserAction).catch(loginErrorAction)
    }

    return (
        <div className='LoginFormDiv'>
            <div>
                <label key="usernameField">
                    username: <label>
                    <input type="text" onChange={updateUsername}/>
                </label>
                </label>
            </div>
            <div>
                <label key="passwordField">
                    password: <label>
                    <input type="password" onChange={updatePassword}/>
                </label>
                </label>
            </div>
            <div>
                <button key="loginButton" onClick={loginButtonAction}>login</button>
            </div>
        </div>
    )
};

export default LoginPage