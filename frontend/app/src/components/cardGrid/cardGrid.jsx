import React, {Children} from 'react';
import './cardGrid.css';

const CardGrid = ({ children }) => {
  
  const result = []
  Children.forEach(children, (child, index) => {
    result.push(child);
  })

  return (
    <>
    <h1 style={{textAlign:"center"}}>Rooms</h1>
    <div className="card-grid">
      {result}
    </div>
    </>
  );
};

export default CardGrid;
