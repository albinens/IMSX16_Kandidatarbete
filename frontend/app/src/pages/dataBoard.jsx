import React from "react";
import './styles/dataBoard.css'


function DataBoard() {

  

  return (
    <>
      <div className='page-header' style={windowWidth < 768 ? {marginTop:"3vh"} : {marginTop:"6vh"}}>
        <h1>Available Rooms</h1>
      </div>
    </>
  );
}

export default DataBoard;