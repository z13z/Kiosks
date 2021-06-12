import {React, useState} from "react"
import "./UsersPage.css"
import RightPanel from "../common/RightPanel/RightPanel"
import SearchPanel from "../common/SearchPanel/SearchPanel"
import UsersGrid from "./UsersGrid";

const UsersPage = () => {
    const [currentState, setCurrentState] = useState({
        userId: "",
        userName: ""
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
            </SearchPanel>
            <UsersGrid id={currentState.userId} name={currentState.userName}/>
        </RightPanel>

    )
}

export default UsersPage