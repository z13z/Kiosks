import React from 'react'

function searchPanel(props) {
    return (
        <div className="SearchPanelDiv">

            {props.children.map((field,) => {
                return (
                    <div className="SearchPanelComponent" key={field.key}>
                        {field}
                    </div>
                )
            })}
        </div>
    )
}

export default searchPanel