import React from 'react'
import "../legend.css"

const HorizontalLegend = () => {
  let green = "#8ED264";
  let yellow = "#F4EC32";
  let red = "#E5414B";
  return (
    <div className='horizontal-legend-container'>
        <div className='legend-circle' style={{backgroundColor: green}}></div>
        <p className='legend-text'>Available</p>
        <div className='legend-circle' style={{backgroundColor: red}}></div>
        <p className='legend-text'>Occupied</p>
        <div className='legend-circle' style={{backgroundColor: yellow}}></div>
        <p className='legend-text'>Booked</p>

    </div>
  )
}

export default HorizontalLegend