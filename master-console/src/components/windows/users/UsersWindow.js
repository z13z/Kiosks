import {React, useState} from "react"
import PopUpWindow from '../common/PopUpWindow'
import {FormGroup, Input, Label} from 'reactstrap';
import {ALL_USER_PERMISSIONS, JWT_TOKEN_KEY} from '../../../Constants'

const UsersWindow = (props) => {
    const [username, setUsername] = useState(props.userToShow !== null ? props.userToShow.username : "")
    const [password, setPassword] = useState("")
    const [repassword, setRePassword] = useState("")
    const [permissions, setPermissions] = useState(props.userToShow !== null ? props.userToShow.permissions : [])
    const [errorStates, setErrorstates] = useState({username:"", pass:"", repass:""})


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
        if(event.target.value === ''){
            setErrorstates({...errorStates, username:"uNameEmpy"})
        }else{
            setErrorstates({...errorStates, username:""})
        }
    }
    const onPasswordChange = (event) => {
        setPassword(event.target.value)
        if(event.target.value === ''){
            setErrorstates({...errorStates, pass:"passEmpty"})
        }else{
            setErrorstates({...errorStates, pass:""})
        }
    }
    const onRePasswordChange = (event) => {
        setRePassword(event.target.value)
        if(event.target.value !== password){
            setErrorstates({...errorStates, repass:"noMatch"})
        }else{
            setErrorstates({...errorStates, repass:""})
        }

    }
    const onPermissionsChange = (event) => {
        setPermissions(Array.from(event.target.selectedOptions, option => option.value))
    }

    const fieldStyle = {float: 'right'}

    return (
        <PopUpWindow {...props} onSubmit={onSubmitAction}>
            <FormGroup key='usernameGroupKey'>
                <Label for="usernameField" 
                    className={errorStates.username === 'uNameEmpy' ? 'emptyField' : ''}
                    >Username</Label>
                <Input type="text" name="username" id={errorStates.username} value={username} onChange={onUsernameChange}
                       style={fieldStyle} required/>
            </FormGroup>
            <FormGroup key='passwordFieldKey'>
                <Label 
                className={errorStates.pass === 'passEmpty' ? 'emptyField' : ''}
                for="passwordField">Password</Label>
                <Input id={errorStates.pass} type="password" name="password" value={password} onChange={onPasswordChange}
                       style={fieldStyle}/>
            </FormGroup>
            <FormGroup key='rePasswordFieldKey'>
                <Label 
                className={errorStates.repass === 'noMatch' ? 'passNoMatch' : ''}
                for="rePasswordField">Reenter Password</Label>
                <Input  type="password" name="rePassword" id={errorStates.repass} value={repassword}
                       onChange={onRePasswordChange} style={fieldStyle}/>
            </FormGroup>
            <FormGroup key='permissionsFieldKey'>
                <Label for="permissionsField">Permissions</Label>
                <Input type={'select'} name='permissions' id='permissionsField' onChange={onPermissionsChange}
                       style={fieldStyle} multiple={true} defaultValue={ALL_USER_PERMISSIONS}>
                    {ALL_USER_PERMISSIONS.map((permissionName,) => {
                        return <option
                            key={permissionName}
                            defaultValue={permissionName}
                            >{permissionName}</option>
                    })
                    }
                </Input>
            </FormGroup>
        </PopUpWindow>
    )
}

export default UsersWindow;