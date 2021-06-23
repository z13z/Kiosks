import React from "react"
import NextArrow from '../../../icons/next-arrow.svg'
import PrevArrow from '../../../icons/prev-arrow.svg'
import FirstArrow from '../../../icons/double-prev-arrow.svg'
import LastArrow from '../../../icons/double-next-arrow.svg'

const GridPagination = (props) => {
    let currentPage = props.currentPage
    return (
        <div className="PaginationPanel">
            <span>
                       {currentPage}/{props.pagesCount}
            </span>
            <img src={FirstArrow} style={{"marginLeft": "10px"}} alt="first" className="PaginationIcon"
                 onClick={() => props.changePage(1)}/>
            <img src={PrevArrow} alt="previous" className="PaginationIcon"
                 onClick={() => props.changePage(currentPage - 1)}/>
            <img src={NextArrow} alt="next" className="PaginationIcon"
                 onClick={() => props.changePage(currentPage + 1)}/>
            <img src={LastArrow} alt="last" className="PaginationIcon"
                 onClick={() => props.changePage(props.pagesCount)}/>
        </div>
    )
};

export default GridPagination