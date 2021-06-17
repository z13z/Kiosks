import {React, useState} from "react"
import PopUpWindow from '../common/PopUpWindow'
import {FormGroup, Input, Label} from 'reactstrap';
import {ALL_USER_PERMISSIONS, JWT_TOKEN_KEY} from '../../../Constants'

const UsersWindow = (props) => {
    const [username, setUsername] = useState(props.userToShow !== null ? props.userToShow.username : "")
    const [password, setPassword] = useState("")
    const [repassword, setRePassword] = useState("")
    const [permissions, setPermissions] = useState(props.userToShow !== null ? props.userToShow.permissions : [])

    const onSubmitAction = () => {
        let queryParams = {}
        if (props.userToShow != null) {
            queryParams['id'] = props.userToShow.id
        }
        queryParams['name'] = username
        queryParams['password'] = password
        queryParams['permissions'] = permissions

        props.axiosMethodToCall('/users', queryParams, {headers: {'Authentication': localStorage.getItem(JWT_TOKEN_KEY)}}).then(() => {
            props.successfullyUpdated()
        }).catch(error => {
            if (error.response.status === 401) {
                localStorage.removeItem(JWT_TOKEN_KEY)
                window.location.reload();
            } else if (error.response.status === 403) {
                alert("action is forbidden")
            } else if (error.response.status === 400) {
                alert("user can't be inserted in database. check if username is unique")
            } else {
                throw error;
            }
            props.onClose()
        })
    }

    const onUsernameChange = (event) => {
        setUsername(event.target.value)
    }
    const onPasswordChange = (event) => {
        setPassword(event.target.value)
    }
    const onRePasswordChange = (event) => {
        setRePassword(event.target.value)
    }
    const onPermissionsChange = (event) => {
        setPermissions(Array.from(event.target.selectedOptions, option => option.value))
    }

    return (
        <PopUpWindow {...props} onSubmit={onSubmitAction}>
            <FormGroup key='usernameGroupKey'>
                <Label for="usernameField">Username</Label>
                <Input type="text" name="username" id="usernameField" value={username} onChange={onUsernameChange}
                       required/>
            </FormGroup>
            <FormGroup key='passwordFieldKey'>
                <Label for="passwordField">Password</Label>
                <Input type="password" name="password" id="passwordField" value={password} onChange={onPasswordChange}/>
            </FormGroup>
            <FormGroup key='rePasswordFieldKey'>
                <Label for="rePasswordField">Reenter Password</Label>
                <Input type="password" name="rePassword" id="rePasswordField" value={repassword}
                       onChange={onRePasswordChange}/>
            </FormGroup>
            <FormGroup key='permissionsFieldKey'>
                <Label for="permissionsField">Password</Label>
                <Input type={'select'} name='permissions' id='permissionsField' onChange={onPermissionsChange} multiple>
                    {ALL_USER_PERMISSIONS.map((permissionName,) => {
                        return <option
                            value={permissionName}
                            selected={permissions.includes(permissionName)}>{permissionName}</option>
                    })
                    }
                </Input>
            </FormGroup>
        </PopUpWindow>
    )
}

export default UsersWindow;