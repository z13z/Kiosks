import {React, useState} from "react"
import "./UsersPage.css"
import RightPanel from "../common/RightPanel/RightPanel"
import SearchPanel from "../common/SearchPanel/SearchPanel"
import UsersGrid from "./UsersGrid";
import UsersWindow from "../../windows/users/UsersWindow"
import axios from "axios";

const UsersPage = () => {
    const [currentState, setCurrentState] = useState({
        userId: "",
        userName: "",
        showCreateWindow: false,
        forceGridUpdate: false,
        userToShow: null
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

    const closeCreateUserWindow = () => {
        newState.showCreateWindow = false
        newState.userToShow = null
        updateState()
    }

    const forceGridUpdate = () => {
        newState.forceGridUpdate = !newState.forceGridUpdate
        updateState()
    }

    const successfullyUpdated = () => {
        newState.forceGridUpdate = !newState.forceGridUpdate
        closeCreateUserWindow()
    }

    const editUserAction = (id, username, permissions) => {
        newState.showCreateWindow = true
        newState.userToShow = {
            id: id,
            username: username,
            permissions: permissions
        }
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
            <UsersGrid id={currentState.userId} name={currentState.userName} forceUpdate={currentState.forceGridUpdate}
                       successfullyUpdated={forceGridUpdate} editUserAction={editUserAction}/>
            {currentState.showCreateWindow ? (
                    <UsersWindow onClose={closeCreateUserWindow} successfullyUpdated={successfullyUpdated}
                                 userToShow={newState.userToShow}
                                 axiosMethodToCall={newState.userToShow == null ? axios.put : axios.post}/>)
                : null}
        </RightPanel>

    )
}

export default UsersPage