import React from 'react'
import "../legend.css"

const VerticalLegend = () => {
  let green = "#8ED264";
  let yellow = "#F4EC32";
  let red = "#E5414B";
  return (
    <div className='vertical-legend-container'>
        <div className='legend-circle' style={{backgroundColor: green}}></div>
        <p>Available</p>
        <div className='legend-circle' style={{backgroundColor: yellow}}></div>
        <p>Booked</p>
        <div className='legend-circle' style={{backgroundColor: red}}></div>
        <p>Occupied</p>

    </div>
  )
}

export default VerticalLegend