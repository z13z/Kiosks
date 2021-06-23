import {React, useState} from "react"
import RightPanel from "../common/RightPanel/RightPanel"
import SearchPanel from "../common/SearchPanel/SearchPanel"
import UsersGrid from "./UsersGrid";
import UsersWindow from "../../windows/users/UsersWindow"
import axios from "axios";

const UsersPage = () => {
    const [currentState, setCurrentState] = useState({
        userId: "",
        userName: "",
        showWindow: false,
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
        newState.showWindow = true
        updateState()
    }

    const closeUserWindow = () => {
        newState.showWindow = false
        newState.userToShow = null
        updateState()
    }

    const forceGridUpdate = () => {
        newState.forceGridUpdate = !newState.forceGridUpdate
        updateState()
    }

    const successfullyUpdated = () => {
        newState.forceGridUpdate = !newState.forceGridUpdate
        closeUserWindow()
    }

    const editUserAction = (id, username, permissions) => {
        newState.showWindow = true
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
                <button key="createUserButton" onClick={addCreateUserWindow}>create</button>
            </SearchPanel>
            <UsersGrid id={currentState.userId} name={currentState.userName} forceUpdate={currentState.forceGridUpdate}
                       successfullyUpdated={forceGridUpdate} editUserAction={editUserAction}/>
            {currentState.showWindow ? (
                    <UsersWindow onClose={closeUserWindow} successfullyUpdated={successfullyUpdated}
                                 userToShow={newState.userToShow}
                                 axiosMethodToCall={newState.userToShow == null ? axios.put : axios.post}/>)
                : null}
        </RightPanel>

    )
}

export default UsersPage