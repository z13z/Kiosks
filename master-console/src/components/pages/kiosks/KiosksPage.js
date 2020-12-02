import React from 'react'
import './KiosksPage.css'
import RightPanel from "../common/RightPanel/RightPanel"
import SearchPanel from "../common/SearchPanel/SearchPanel"
import KiosksGrid from "./KiosksGrid";

function kiosksPage() {
    return (
        <RightPanel>
            <SearchPanel>
                <label>
                    id: <label>
                    <input type="number"/>
                </label>
                </label>
                <label>
                    name: <label>
                    <input type="text"/>
                </label>
                </label>
                <button>search</button>
            </SearchPanel>
            <KiosksGrid/>
        </RightPanel>

    )
}

export default kiosksPage