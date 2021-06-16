import {React, useState} from "react"
import "./UsersPage.css"
import RightPanel from "../common/RightPanel/RightPanel"
import SearchPanel from "../common/SearchPanel/SearchPanel"
import UsersGrid from "./UsersGrid";
import UsersWindow from "../../windows/users/UsersWindow"

const UsersPage = () => {
    const [currentState, setCurrentState] = useState({
        userId: "",
        userName: "",
        showCreateWindow: false
    });

    let newState = {...currentState}

    const updateUserName = (event) => {
        newState.userName = event.target.value
    }

    const updateUserId = (event) => {
        newState.userId = event.target.value
    }

    const updateState = () => {
        setCurrentState(newState)
    }

    const addCreateUserWindow = () => {
        newState.showCreateWindow = true
        updateState()
    }

    return (
        <RightPanel>
            <SearchPanel>
                <label key="userIdField">
                    id: <label>
                    <input type="number" onChange={updateUserId}/>
                </label>
                </label>
                <label key="userNameField">
                    name: <label>
                    <input type="text" onChange={updateUserName}/>
                </label>
                </label>
                <button key="userSearchButton" onClick={updateState}>search</button>
                <button key="createUserButton" onClick={addCreateUserWindow}>createUserButton</button>
            </SearchPanel>
            <UsersGrid id={currentState.userId} name={currentState.userName}/>
            {currentState.showCreateWindow ? (<UsersWindow show='true'/>) : null}
        </RightPanel>

    )
}

export default UsersPage