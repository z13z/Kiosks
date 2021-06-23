import React from 'react'

function rightPanel(props) {
    return (
        <div className="RightPanelDiv">
            {props.children}
        </div>
    )
}

export default rightPanel