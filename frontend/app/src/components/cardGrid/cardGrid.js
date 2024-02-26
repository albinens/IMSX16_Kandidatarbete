import React, {Children} from 'react';

const CardGrid = ({ children }) => {
  
  const result = []
  Children.forEach(children, (child, index) => {
    result.push(child);
  })

  return (
    <div className="card-grid">
      {result}
    </div>
  );
};

export default CardGrid;
