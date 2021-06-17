import {React, useState} from "react"
import PopUpWindow from '../common/PopUpWindow'
import {FormGroup, Input, Label} from 'reactstrap';
import {ALL_USER_PERMISSIONS} from '../../../Constants'

const UsersWindow = (props) => {
    const [username, setUsername] = useState(props.username !== undefined ? props.username : "")
    const [password, setPassword] = useState("")
    const [repassword, setRePassword] = useState("")
    const [permissions, setPermissions] = useState([])

    const onSubmitAction = () => {
        console.log(username)
        console.log(password)
        console.log(repassword)
        console.log(permissions)
        props.onClose()
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