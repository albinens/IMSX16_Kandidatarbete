import React, { useEffect } from 'react';
import LinePlot from '../components/dataCharts/chart/linePlot';
import Table from '../components/dataCharts/table/table';
import './styles/dataBoard.css';

const DataBoard = () => {
  const dataset = [];
  for (let i = 0; i < 80; i++) {
    dataset.push(Math.random() * 80);
  }

  return (
    <>
      <div className='page-header'>
        <h1>Data Insights</h1>
      </div>
      <div className='two-column-wrapper-dataBoard'>
        {/* LEFT COLUMN */ }
        <div className='left-column-dataBoard'>

        </div>
        {/* RIGHT COLUMN */ }
        <div className='right-column-dataBoard'>
          <LinePlot data={dataset}/>
          <LinePlot data={dataset}/>
        </div>
      </div>
    </>

  );
}

export default DataBoard;